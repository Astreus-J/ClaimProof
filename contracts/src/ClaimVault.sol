// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

/// @title ClaimVault
/// @notice Holds the underwriting pool and per-order protection records on
///         Creditcoin. A claim payout is authorized only after re-verifying,
///         via the Attestcoin Protocol, that the order's trigger event (a
///         `DeliveryFailed` emission on Ethereum Sepolia) actually happened
///         — see docs/architecture.md and docs/ATTESTCOIN_INTEGRATION.md.
///         That verification is implemented in `submitClaim`, added in
///         Sprint 2; this skeleton establishes order registration, the
///         underwriting pool, and access control only.
contract ClaimVault {
    /// @notice Per-order protection record.
    struct Order {
        address buyer;
        uint256 protectionAmount;
        bool claimed;
    }

    /// @notice Contract owner, authorized to update policy parameters.
    address public immutable owner;

    /// @notice Address of the off-chain worker authorized to submit claims
    ///         once `submitClaim` is added in Sprint 2.
    address public worker;

    /// @notice Maximum payout `submitClaim` may authorize for a single order,
    ///         regardless of what the AI claims agent suggests.
    uint256 public payoutCap;

    /// @notice orderId => registered protection order.
    mapping(uint256 => Order) public orders;

    /// @notice Emitted when a new protection order is registered.
    event OrderRegistered(uint256 indexed orderId, address indexed buyer, uint256 protectionAmount);

    /// @notice Emitted when capital is added to the underwriting pool.
    event PoolFunded(address indexed funder, uint256 amount);

    /// @notice Emitted when the authorized off-chain worker address changes.
    event WorkerUpdated(address indexed previousWorker, address indexed newWorker);

    /// @notice Emitted when the payout cap changes.
    event PayoutCapUpdated(uint256 previousCap, uint256 newCap);

    /// @notice Thrown when a non-owner calls an owner-only function.
    error NotOwner();

    /// @notice Thrown when `registerOrder` is called with an `orderId` already in use.
    error OrderAlreadyExists(uint256 orderId);

    /// @notice Thrown when a zero address is supplied where a non-zero address is required.
    error ZeroAddress();

    /// @notice Restricts a function to the contract owner.
    modifier onlyOwner() {
        if (msg.sender != owner) revert NotOwner();
        _;
    }

    /// @param initialWorker Address of the off-chain worker wallet.
    /// @param initialPayoutCap Maximum payout `submitClaim` may authorize per order.
    constructor(address initialWorker, uint256 initialPayoutCap) {
        if (initialWorker == address(0)) revert ZeroAddress();
        owner = msg.sender;
        worker = initialWorker;
        payoutCap = initialPayoutCap;
    }

    /// @notice Registers a new protection order, mirroring the shipment
    ///         created on `DeliveryTrackerMock`.
    /// @param orderId Unique identifier shared with `DeliveryTrackerMock`.
    /// @param buyer Address entitled to the payout if the claim is verified.
    /// @param protectionAmount Amount payable if the trigger event is verified.
    function registerOrder(uint256 orderId, address buyer, uint256 protectionAmount) external {
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

    /// @notice Updates the maximum payout `submitClaim` may authorize.
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
}
