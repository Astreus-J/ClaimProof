// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {INativeQueryVerifier, NativeQueryVerifierLib} from "./interfaces/INativeQueryVerifier.sol";
import {EvmV1Decoder} from "usc-contracts/decoding/EvmV1Decoder.sol";

/// @title ClaimVault
/// @notice Holds the underwriting pool and per-order protection records on Creditcoin.
///         A claim payout is authorized only after `submitClaim` re-verifies, via the
///         Attestcoin Protocol, that the order's trigger event (a `DeliveryFailed`
///         emission on Ethereum Sepolia) actually happened — see docs/architecture.md
///         and docs/ATTESTCOIN_INTEGRATION.md. The AI claims agent (Sprint 3) may only
///         ever narrow the payout further via `payoutCap`; it has no path to authorize
///         a payout on its own.
contract ClaimVault {
    /// @notice Per-order protection record.
    struct Order {
        address buyer;
        uint256 protectionAmount;
        bool claimed;
    }

    /// @notice keccak256("DeliveryFailed(uint256,address,uint256)") — DeliveryTrackerMock's
    ///         trigger event signature, used to find it among a transaction's decoded logs.
    bytes32 public constant DELIVERY_FAILED_EVENT_SIGNATURE = keccak256("DeliveryFailed(uint256,address,uint256)");

    /// @notice Contract owner, authorized to update policy parameters.
    address public immutable owner;

    /// @notice The trusted `DeliveryTrackerMock` contract address on Ethereum Sepolia.
    /// @dev The precompile proves a log's *content*, not which contract emitted it — any
    ///      Sepolia contract could emit a same-signature event, so the emitter address
    ///      must be checked explicitly against this value (see docs/THREAT_MODEL.md, T1).
    address public immutable sourceContract;

    /// @notice The Attestcoin native query verifier precompile.
    INativeQueryVerifier public immutable VERIFIER;

    /// @notice Address of the off-chain worker authorized to update policy parameters
    ///         on the storefront's behalf. `submitClaim` itself is permissionless —
    ///         see its NatSpec for why that's safe.
    address public worker;

    /// @notice Maximum payout `submitClaim` may release for a single order, regardless
    ///         of the order's registered protection amount.
    uint256 public payoutCap;

    /// @notice orderId => registered protection order.
    mapping(uint256 => Order) public orders;

    /// @notice queryId => whether an Attestcoin proof has already been used for a payout.
    mapping(bytes32 => bool) public processedQueries;

    /// @notice Emitted when a new protection order is registered.
    event OrderRegistered(uint256 indexed orderId, address indexed buyer, uint256 protectionAmount);

    /// @notice Emitted when capital is added to the underwriting pool.
    event PoolFunded(address indexed funder, uint256 amount);

    /// @notice Emitted when the authorized off-chain worker address changes.
    event WorkerUpdated(address indexed previousWorker, address indexed newWorker);

    /// @notice Emitted when the payout cap changes.
    event PayoutCapUpdated(uint256 previousCap, uint256 newCap);

    /// @notice Emitted when a claim is successfully verified and paid out.
    event ClaimPaid(uint256 indexed orderId, address indexed buyer, uint256 amount, bytes32 indexed queryId);

    /// @notice Thrown when a non-owner calls an owner-only function.
    error NotOwner();

    /// @notice Thrown when `registerOrder` is called with an `orderId` already in use.
    error OrderAlreadyExists(uint256 orderId);

    /// @notice Thrown when a zero address is supplied where a non-zero address is required.
    error ZeroAddress();

    /// @notice Thrown when the Attestcoin precompile rejects the inclusion/continuity proof.
    error InvalidProof();

    /// @notice Thrown when a proof's query id has already been used for a payout.
    error QueryAlreadyProcessed(bytes32 queryId);

    /// @notice Thrown when the verified source transaction's receipt status isn't success
    ///         (`0x1`) — the precompile proves inclusion only, never success.
    error SourceTransactionNotSuccessful();

    /// @notice Thrown when no `DeliveryFailed` log from the trusted `sourceContract` is
    ///         found in the verified transaction's receipt.
    error DeliveryFailedEventNotFound();

    /// @notice Thrown when the decoded `orderId` has no matching registered order.
    error OrderNotFound(uint256 orderId);

    /// @notice Thrown when the decoded event's buyer doesn't match the registered order's buyer.
    error OrderBuyerMismatch(uint256 orderId);

    /// @notice Thrown when the order has already been paid out.
    error OrderAlreadyClaimed(uint256 orderId);

    /// @notice Thrown when the payout transfer to the buyer fails.
    error PayoutTransferFailed();

    /// @notice Thrown when a non-worker calls a worker-only function.
    error NotWorker();

    /// @notice Restricts a function to the contract owner.
    modifier onlyOwner() {
        if (msg.sender != owner) revert NotOwner();
        _;
    }

    /// @notice Restricts a function to the authorized off-chain worker.
    modifier onlyWorker() {
        if (msg.sender != worker) revert NotWorker();
        _;
    }

    /// @param sourceContract_ Address of `DeliveryTrackerMock` on Ethereum Sepolia.
    /// @param initialWorker Address of the off-chain worker wallet.
    /// @param initialPayoutCap Maximum payout `submitClaim` may release per order.
    constructor(address sourceContract_, address initialWorker, uint256 initialPayoutCap) {
        if (sourceContract_ == address(0)) revert ZeroAddress();
        if (initialWorker == address(0)) revert ZeroAddress();
        owner = msg.sender;
        sourceContract = sourceContract_;
        worker = initialWorker;
        payoutCap = initialPayoutCap;
        VERIFIER = NativeQueryVerifierLib.getVerifier();
    }

    /// @notice Registers a new protection order, mirroring the shipment created on
    ///         `DeliveryTrackerMock`.
    /// @dev Worker-only: `DeliveryTrackerMock.createShipment`/`reportDeliveryFailure`
    ///      are permissionless by design, so this is the only gate on which
    ///      (orderId, buyer, protectionAmount) triples `submitClaim` will ever honor —
    ///      without it, anyone could self-register an order and self-trigger a
    ///      genuinely attestable failure to drain the pool for free.
    /// @param orderId Unique identifier shared with `DeliveryTrackerMock`.
    /// @param buyer Address entitled to the payout if the claim is verified.
    /// @param protectionAmount Amount payable if the trigger event is verified.
    function registerOrder(uint256 orderId, address buyer, uint256 protectionAmount) external onlyWorker {
        if (orders[orderId].buyer != address(0)) revert OrderAlreadyExists(orderId);
        if (buyer == address(0)) revert ZeroAddress();

        orders[orderId] = Order({buyer: buyer, protectionAmount: protectionAmount, claimed: false});
        emit OrderRegistered(orderId, buyer, protectionAmount);
    }

    /// @notice Adds capital to the underwriting pool.
    function fundPool() external payable {
        emit PoolFunded(msg.sender, msg.value);
    }

    /// @notice Updates the authorized off-chain worker address.
    /// @param newWorker New worker wallet address.
    function setWorker(address newWorker) external onlyOwner {
        if (newWorker == address(0)) revert ZeroAddress();
        emit WorkerUpdated(worker, newWorker);
        worker = newWorker;
    }

    /// @notice Updates the maximum payout `submitClaim` may release.
    /// @param newCap New payout cap, in wei.
    function setPayoutCap(uint256 newCap) external onlyOwner {
        emit PayoutCapUpdated(payoutCap, newCap);
        payoutCap = newCap;
    }

    /// @notice Returns the underwriting pool's current balance.
    /// @return The contract's CTC balance, in wei.
    function poolBalance() external view returns (uint256) {
        return address(this).balance;
    }

    /// @notice Verifies that `DeliveryTrackerMock` emitted `DeliveryFailed` for an order
    ///         on Ethereum Sepolia — re-checked here on-chain via the Attestcoin
    ///         precompile — and pays out that order's registered buyer.
    /// @dev Deliberately permissionless: anyone may relay a valid proof. Safety comes
    ///      from Attestcoin's own re-verification, the emitter-address check, and the
    ///      payout cap — never from restricting who can call this function. The AI
    ///      claims agent's `suggestedPayout` may only ever narrow the payout — it can
    ///      lower it below `protectionAmount` (a partial refund) but can never raise it
    ///      above `protectionAmount` or the on-chain `payoutCap` (see docs/THREAT_MODEL.md, T4).
    /// @param chainKey Attestcoin's identifier for the source chain (Ethereum Sepolia).
    /// @param blockHeight Sepolia block height containing the `DeliveryFailed` transaction.
    /// @param encodedTransaction The attested transaction + receipt, from the Prover API.
    /// @param merkleRoot Merkle root of the transaction's block.
    /// @param siblings Merkle inclusion proof siblings for the transaction.
    /// @param lowerEndpointDigest Continuity proof's lower endpoint digest.
    /// @param continuityRoots Continuity proof's chain of block digests.
    /// @param suggestedPayout The AI claims agent's suggested payout, in wei — advisory
    ///        only, always bounded below by `protectionAmount` and `payoutCap`.
    /// @return orderId The order the verified claim was paid out for.
    /// @return payoutAmount The amount released to the order's buyer.
    function submitClaim(
        uint64 chainKey,
        uint64 blockHeight,
        bytes calldata encodedTransaction,
        bytes32 merkleRoot,
        INativeQueryVerifier.MerkleProofEntry[] calldata siblings,
        bytes32 lowerEndpointDigest,
        bytes32[] calldata continuityRoots,
        uint256 suggestedPayout
    ) external returns (uint256 orderId, uint256 payoutAmount) {
        INativeQueryVerifier.MerkleProof memory merkleProof =
            INativeQueryVerifier.MerkleProof({root: merkleRoot, siblings: siblings});

        // Anti-replay check happens before spending gas on verification.
        bytes32 queryId = _computeQueryId(chainKey, blockHeight, merkleProof);
        if (processedQueries[queryId]) revert QueryAlreadyProcessed(queryId);

        INativeQueryVerifier.ContinuityProof memory continuityProof =
            INativeQueryVerifier.ContinuityProof({lowerEndpointDigest: lowerEndpointDigest, roots: continuityRoots});

        bool verified = VERIFIER.verifyAndEmit(chainKey, blockHeight, encodedTransaction, merkleProof, continuityProof);
        if (!verified) revert InvalidProof();

        processedQueries[queryId] = true;

        // The precompile proves inclusion only, not success — status is checked explicitly.
        EvmV1Decoder.ReceiptFields memory receipt = EvmV1Decoder.decodeReceiptFields(encodedTransaction);
        if (receipt.receiptStatus != 1) revert SourceTransactionNotSuccessful();

        address buyer;
        (orderId, buyer) = _findDeliveryFailedEvent(receipt.receiptLogs);

        Order storage order = orders[orderId];
        if (order.buyer == address(0)) revert OrderNotFound(orderId);
        if (order.buyer != buyer) revert OrderBuyerMismatch(orderId);
        if (order.claimed) revert OrderAlreadyClaimed(orderId);

        uint256 boundedBySuggestion =
            suggestedPayout < order.protectionAmount ? suggestedPayout : order.protectionAmount;
        payoutAmount = boundedBySuggestion > payoutCap ? payoutCap : boundedBySuggestion;
        order.claimed = true;

        emit ClaimPaid(orderId, buyer, payoutAmount, queryId);

        // Low-level call flagged by static analysis (Decurity's arbitrary-low-level-call
        // rule) — not arbitrary in practice: `buyer` is read from a registered Order, only
        // ever set by the trusted worker via registerOrder (onlyWorker), never by this
        // function's caller. Calldata is empty, so this is a plain value transfer, not a
        // delegatecall or attacker-chosen selector. `claimed` is already set above
        // (checks-effects-interactions), so a reentrant call for this same order/proof
        // hits OrderAlreadyClaimed regardless of what `buyer` does with the funds.
        (bool sent,) = buyer.call{value: payoutAmount}("");
        if (!sent) revert PayoutTransferFailed();
    }

    /// @notice Finds the `DeliveryFailed` log emitted by the trusted `sourceContract`
    ///         among a transaction's decoded logs, and decodes its indexed fields.
    /// @dev `EvmV1Decoder` filters by event signature only — it does not know which
    ///      contract emitted a log — so the emitter address is checked here explicitly.
    function _findDeliveryFailedEvent(EvmV1Decoder.LogEntry[] memory logs)
        private
        view
        returns (uint256 orderId, address buyer)
    {
        for (uint256 i = 0; i < logs.length; i++) {
            EvmV1Decoder.LogEntry memory log = logs[i];
            if (
                log.address_ == sourceContract && log.topics.length == 3
                    && log.topics[0] == DELIVERY_FAILED_EVENT_SIGNATURE
            ) {
                return (uint256(log.topics[1]), address(uint160(uint256(log.topics[2]))));
            }
        }
        revert DeliveryFailedEventNotFound();
    }

    /// @notice Derives a collision-resistant anti-replay key from the proof's
    ///         (chainKey, blockHeight, transaction-index-within-block) triple — the
    ///         same transaction attested a second time always yields the same queryId.
    /// @dev The transaction index is required in addition to chainKey/blockHeight
    ///      because multiple transactions in the same block share the same merkleRoot.
    function _computeQueryId(uint64 chainKey, uint64 blockHeight, INativeQueryVerifier.MerkleProof memory merkleProof)
        private
        view
        returns (bytes32 queryId)
    {
        uint64 txIndex = VERIFIER.calculateTxIndex(merkleProof);
        queryId = keccak256(abi.encodePacked(chainKey, blockHeight, txIndex));
    }
}
