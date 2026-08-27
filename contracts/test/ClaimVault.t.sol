// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {INativeQueryVerifier} from "../src/interfaces/INativeQueryVerifier.sol";
import {EvmV1Decoder} from "usc-contracts/decoding/EvmV1Decoder.sol";
import {ClaimVault} from "../src/ClaimVault.sol";
import {AttestcoinFixtures} from "./helpers/AttestcoinFixtures.sol";

contract ClaimVaultTest is AttestcoinFixtures {
    ClaimVault internal vault;

    address internal owner = address(this);
    address internal worker = makeAddr("worker");
    address internal buyer = makeAddr("buyer");
    address internal sourceContract = makeAddr("deliveryTrackerMock");
    uint256 internal constant PAYOUT_CAP = 1 ether;
    uint256 internal constant ORDER_ID = 1;
    uint256 internal constant PROTECTION_AMOUNT = 0.1 ether;

    uint64 internal constant CHAIN_KEY = 1;
    uint64 internal constant BLOCK_HEIGHT = 100;
    uint64 internal constant TX_INDEX = 0;

    function setUp() public {
        vault = new ClaimVault(sourceContract, worker, PAYOUT_CAP);
        vm.deal(address(vault), 10 ether);
    }

    // ---------- Sprint 1 skeleton behavior (unchanged) ----------

    function test_Constructor_SetsOwnerWorkerAndCap() public view {
        assertEq(vault.owner(), owner);
        assertEq(vault.worker(), worker);
        assertEq(vault.payoutCap(), PAYOUT_CAP);
        assertEq(vault.sourceContract(), sourceContract);
    }

    function test_Constructor_RevertsOnZeroAddressWorker() public {
        vm.expectRevert(ClaimVault.ZeroAddress.selector);
        new ClaimVault(sourceContract, address(0), PAYOUT_CAP);
    }

    function test_Constructor_RevertsOnZeroAddressSourceContract() public {
        vm.expectRevert(ClaimVault.ZeroAddress.selector);
        new ClaimVault(address(0), worker, PAYOUT_CAP);
    }

    function test_RegisterOrder_StoresRecordAndEmitsEvent() public {
        vm.expectEmit(true, true, false, true);
        emit ClaimVault.OrderRegistered(ORDER_ID, buyer, PROTECTION_AMOUNT);

        vm.prank(worker);
        vault.registerOrder(ORDER_ID, buyer, PROTECTION_AMOUNT);

        (address storedBuyer, uint256 storedAmount, bool claimed) = vault.orders(ORDER_ID);
        assertEq(storedBuyer, buyer);
        assertEq(storedAmount, PROTECTION_AMOUNT);
        assertFalse(claimed);
    }

    function test_RegisterOrder_RevertsOnDuplicateOrderId() public {
        vm.prank(worker);
        vault.registerOrder(ORDER_ID, buyer, PROTECTION_AMOUNT);

        vm.expectRevert(abi.encodeWithSelector(ClaimVault.OrderAlreadyExists.selector, ORDER_ID));
        vm.prank(worker);
        vault.registerOrder(ORDER_ID, buyer, PROTECTION_AMOUNT);
    }

    function test_RegisterOrder_RevertsOnZeroAddressBuyer() public {
        vm.expectRevert(ClaimVault.ZeroAddress.selector);
        vm.prank(worker);
        vault.registerOrder(ORDER_ID, address(0), PROTECTION_AMOUNT);
    }

    function test_RegisterOrder_RevertsIfCallerIsNotWorker() public {
        // Anyone other than the worker registering an order (with themselves as buyer
        // and an arbitrary protectionAmount) is the root of a pool-draining exploit:
        // pair it with a self-triggered, genuinely-attestable DeliveryFailed on
        // DeliveryTrackerMock (which is permissionless by design) and submitClaim
        // would pay out for free. registerOrder must be worker-gated.
        address attacker = makeAddr("attacker");
        vm.prank(attacker);
        vm.expectRevert(ClaimVault.NotWorker.selector);
        vault.registerOrder(ORDER_ID, attacker, 1000 ether);
    }

    function test_FundPool_IncreasesBalanceAndEmitsEvent() public {
        address funder = makeAddr("funder");
        vm.deal(funder, 1 ether);

        vm.expectEmit(true, false, false, true);
        emit ClaimVault.PoolFunded(funder, 1 ether);

        vm.prank(funder);
        vault.fundPool{value: 1 ether}();

        assertEq(vault.poolBalance(), 11 ether);
    }

    function test_SetWorker_OnlyOwnerCanCall() public {
        address newWorker = makeAddr("newWorker");

        vm.prank(buyer);
        vm.expectRevert(ClaimVault.NotOwner.selector);
        vault.setWorker(newWorker);
    }

    function test_SetWorker_UpdatesWorkerAndEmitsEvent() public {
        address newWorker = makeAddr("newWorker");

        vm.expectEmit(true, true, false, false);
        emit ClaimVault.WorkerUpdated(worker, newWorker);

        vault.setWorker(newWorker);

        assertEq(vault.worker(), newWorker);
    }

    function test_SetWorker_RevertsOnZeroAddress() public {
        vm.expectRevert(ClaimVault.ZeroAddress.selector);
        vault.setWorker(address(0));
    }

    function test_SetPayoutCap_OnlyOwnerCanCall() public {
        vm.prank(buyer);
        vm.expectRevert(ClaimVault.NotOwner.selector);
        vault.setPayoutCap(2 ether);
    }

    function test_SetPayoutCap_UpdatesCapAndEmitsEvent() public {
        vm.expectEmit(false, false, false, true);
        emit ClaimVault.PayoutCapUpdated(PAYOUT_CAP, 2 ether);

        vault.setPayoutCap(2 ether);

        assertEq(vault.payoutCap(), 2 ether);
    }

    // ---------- Sprint 2: submitClaim / Attestcoin verification ----------

    function _validProofArgs()
        internal
        view
        returns (
            uint64 chainKey,
            uint64 blockHeight,
            bytes memory encodedTransaction,
            bytes32 merkleRoot,
            INativeQueryVerifier.MerkleProofEntry[] memory siblings,
            bytes32 lowerEndpointDigest,
            bytes32[] memory continuityRoots
        )
    {
        chainKey = CHAIN_KEY;
        blockHeight = BLOCK_HEIGHT;
        merkleRoot = bytes32(uint256(0xabc));
        siblings = new INativeQueryVerifier.MerkleProofEntry[](0);
        lowerEndpointDigest = bytes32(uint256(0xdef));
        continuityRoots = new bytes32[](0);

        EvmV1Decoder.LogEntryTuple[] memory logs = new EvmV1Decoder.LogEntryTuple[](1);
        logs[0] = _deliveryFailedLog(sourceContract, ORDER_ID, buyer, block.timestamp);
        encodedTransaction = _encodeTransaction(1, logs);
    }

    function test_SubmitClaim_VerifiedEventPaysOutBuyer() public {
        vm.prank(worker);
        vault.registerOrder(ORDER_ID, buyer, PROTECTION_AMOUNT);

        (
            uint64 chainKey,
            uint64 blockHeight,
            bytes memory encodedTransaction,
            bytes32 merkleRoot,
            INativeQueryVerifier.MerkleProofEntry[] memory siblings,
            bytes32 lowerEndpointDigest,
            bytes32[] memory continuityRoots
        ) = _validProofArgs();

        _mockCalculateTxIndex(TX_INDEX);
        _mockVerifyAndEmit(true);

        uint256 buyerBalanceBefore = buyer.balance;

        (uint256 orderId, uint256 payoutAmount) = vault.submitClaim(
            chainKey,
            blockHeight,
            encodedTransaction,
            merkleRoot,
            siblings,
            lowerEndpointDigest,
            continuityRoots,
            PROTECTION_AMOUNT
        );

        assertEq(orderId, ORDER_ID);
        assertEq(payoutAmount, PROTECTION_AMOUNT);
        assertEq(buyer.balance, buyerBalanceBefore + PROTECTION_AMOUNT);

        (,, bool claimed) = vault.orders(ORDER_ID);
        assertTrue(claimed);
    }

    function test_SubmitClaim_PayoutCappedBelowProtectionAmount() public {
        vault.setPayoutCap(0.01 ether);
        vm.prank(worker);
        vault.registerOrder(ORDER_ID, buyer, PROTECTION_AMOUNT);

        (
            uint64 chainKey,
            uint64 blockHeight,
            bytes memory encodedTransaction,
            bytes32 merkleRoot,
            INativeQueryVerifier.MerkleProofEntry[] memory siblings,
            bytes32 lowerEndpointDigest,
            bytes32[] memory continuityRoots
        ) = _validProofArgs();

        _mockCalculateTxIndex(TX_INDEX);
        _mockVerifyAndEmit(true);

        (, uint256 payoutAmount) = vault.submitClaim(
            chainKey,
            blockHeight,
            encodedTransaction,
            merkleRoot,
            siblings,
            lowerEndpointDigest,
            continuityRoots,
            PROTECTION_AMOUNT
        );

        assertEq(payoutAmount, 0.01 ether);
    }

    function test_SubmitClaim_PaysExactlySuggestedValueWhenBelowProtectionAndCap() public {
        // The AI claims agent may suggest a partial refund (e.g. a minor SLA breach) —
        // the contract must honor a suggestion strictly below both the registered
        // protectionAmount and the policy cap, not silently pay the maximum instead.
        vm.prank(worker);
        vault.registerOrder(ORDER_ID, buyer, PROTECTION_AMOUNT);

        (
            uint64 chainKey,
            uint64 blockHeight,
            bytes memory encodedTransaction,
            bytes32 merkleRoot,
            INativeQueryVerifier.MerkleProofEntry[] memory siblings,
            bytes32 lowerEndpointDigest,
            bytes32[] memory continuityRoots
        ) = _validProofArgs();

        _mockCalculateTxIndex(TX_INDEX);
        _mockVerifyAndEmit(true);

        uint256 suggestedPayout = PROTECTION_AMOUNT / 4;

        (, uint256 payoutAmount) = vault.submitClaim(
            chainKey,
            blockHeight,
            encodedTransaction,
            merkleRoot,
            siblings,
            lowerEndpointDigest,
            continuityRoots,
            suggestedPayout
        );

        assertEq(payoutAmount, suggestedPayout);
    }

    function test_SubmitClaim_SuggestedValueAbovePolicyCap_IsCappedNotHonored() public {
        // T4 (docs/THREAT_MODEL.md): a compromised or buggy AI claims agent suggesting
        // an excessive value must never result in a payout above the on-chain cap.
        vault.setPayoutCap(0.01 ether);
        vm.prank(worker);
        vault.registerOrder(ORDER_ID, buyer, PROTECTION_AMOUNT);

        (
            uint64 chainKey,
            uint64 blockHeight,
            bytes memory encodedTransaction,
            bytes32 merkleRoot,
            INativeQueryVerifier.MerkleProofEntry[] memory siblings,
            bytes32 lowerEndpointDigest,
            bytes32[] memory continuityRoots
        ) = _validProofArgs();

        _mockCalculateTxIndex(TX_INDEX);
        _mockVerifyAndEmit(true);

        uint256 excessiveSuggestedPayout = 1000 ether;

        (, uint256 payoutAmount) = vault.submitClaim(
            chainKey,
            blockHeight,
            encodedTransaction,
            merkleRoot,
            siblings,
            lowerEndpointDigest,
            continuityRoots,
            excessiveSuggestedPayout
        );

        assertEq(payoutAmount, 0.01 ether);
    }

    function test_SubmitClaim_RevertsOnInvalidProof() public {
        vm.prank(worker);
        vault.registerOrder(ORDER_ID, buyer, PROTECTION_AMOUNT);

        (
            uint64 chainKey,
            uint64 blockHeight,
            bytes memory encodedTransaction,
            bytes32 merkleRoot,
            INativeQueryVerifier.MerkleProofEntry[] memory siblings,
            bytes32 lowerEndpointDigest,
            bytes32[] memory continuityRoots
        ) = _validProofArgs();

        _mockCalculateTxIndex(TX_INDEX);
        _mockVerifyAndEmit(false);

        vm.expectRevert(ClaimVault.InvalidProof.selector);
        vault.submitClaim(
            chainKey,
            blockHeight,
            encodedTransaction,
            merkleRoot,
            siblings,
            lowerEndpointDigest,
            continuityRoots,
            PROTECTION_AMOUNT
        );
    }

    function test_SubmitClaim_RevertsOnNonexistentOrderId() public {
        // Note: order is never registered.
        (
            uint64 chainKey,
            uint64 blockHeight,
            bytes memory encodedTransaction,
            bytes32 merkleRoot,
            INativeQueryVerifier.MerkleProofEntry[] memory siblings,
            bytes32 lowerEndpointDigest,
            bytes32[] memory continuityRoots
        ) = _validProofArgs();

        _mockCalculateTxIndex(TX_INDEX);
        _mockVerifyAndEmit(true);

        vm.expectRevert(abi.encodeWithSelector(ClaimVault.OrderNotFound.selector, ORDER_ID));
        vault.submitClaim(
            chainKey,
            blockHeight,
            encodedTransaction,
            merkleRoot,
            siblings,
            lowerEndpointDigest,
            continuityRoots,
            PROTECTION_AMOUNT
        );
    }

    function test_SubmitClaim_RevertsOnUnsuccessfulSourceTransaction() public {
        vm.prank(worker);
        vault.registerOrder(ORDER_ID, buyer, PROTECTION_AMOUNT);

        EvmV1Decoder.LogEntryTuple[] memory logs = new EvmV1Decoder.LogEntryTuple[](1);
        logs[0] = _deliveryFailedLog(sourceContract, ORDER_ID, buyer, block.timestamp);
        bytes memory encodedTransaction = _encodeTransaction(0, logs); // status != 1

        _mockCalculateTxIndex(TX_INDEX);
        _mockVerifyAndEmit(true);

        vm.expectRevert(ClaimVault.SourceTransactionNotSuccessful.selector);
        vault.submitClaim(
            CHAIN_KEY,
            BLOCK_HEIGHT,
            encodedTransaction,
            bytes32(uint256(0xabc)),
            new INativeQueryVerifier.MerkleProofEntry[](0),
            bytes32(uint256(0xdef)),
            new bytes32[](0),
            PROTECTION_AMOUNT
        );
    }

    function test_SubmitClaim_RevertsOnEventFromUntrustedContract() public {
        vm.prank(worker);
        vault.registerOrder(ORDER_ID, buyer, PROTECTION_AMOUNT);

        address attacker = makeAddr("attackerContract");
        EvmV1Decoder.LogEntryTuple[] memory logs = new EvmV1Decoder.LogEntryTuple[](1);
        logs[0] = _deliveryFailedLog(attacker, ORDER_ID, buyer, block.timestamp);
        bytes memory encodedTransaction = _encodeTransaction(1, logs);

        _mockCalculateTxIndex(TX_INDEX);
        _mockVerifyAndEmit(true);

        vm.expectRevert(ClaimVault.DeliveryFailedEventNotFound.selector);
        vault.submitClaim(
            CHAIN_KEY,
            BLOCK_HEIGHT,
            encodedTransaction,
            bytes32(uint256(0xabc)),
            new INativeQueryVerifier.MerkleProofEntry[](0),
            bytes32(uint256(0xdef)),
            new bytes32[](0),
            PROTECTION_AMOUNT
        );
    }

    function test_SubmitClaim_RevertsOnBuyerMismatch() public {
        vm.prank(worker);
        vault.registerOrder(ORDER_ID, buyer, PROTECTION_AMOUNT);

        address someoneElse = makeAddr("someoneElse");
        EvmV1Decoder.LogEntryTuple[] memory logs = new EvmV1Decoder.LogEntryTuple[](1);
        logs[0] = _deliveryFailedLog(sourceContract, ORDER_ID, someoneElse, block.timestamp);
        bytes memory encodedTransaction = _encodeTransaction(1, logs);

        _mockCalculateTxIndex(TX_INDEX);
        _mockVerifyAndEmit(true);

        vm.expectRevert(abi.encodeWithSelector(ClaimVault.OrderBuyerMismatch.selector, ORDER_ID));
        vault.submitClaim(
            CHAIN_KEY,
            BLOCK_HEIGHT,
            encodedTransaction,
            bytes32(uint256(0xabc)),
            new INativeQueryVerifier.MerkleProofEntry[](0),
            bytes32(uint256(0xdef)),
            new bytes32[](0),
            PROTECTION_AMOUNT
        );
    }

    function test_SubmitClaim_RevertsOnAlreadyClaimedOrder() public {
        vm.prank(worker);
        vault.registerOrder(ORDER_ID, buyer, PROTECTION_AMOUNT);

        (
            uint64 chainKey,
            uint64 blockHeight,
            bytes memory encodedTransaction,
            bytes32 merkleRoot,
            INativeQueryVerifier.MerkleProofEntry[] memory siblings,
            bytes32 lowerEndpointDigest,
            bytes32[] memory continuityRoots
        ) = _validProofArgs();

        _mockCalculateTxIndex(TX_INDEX);
        _mockVerifyAndEmit(true);

        vault.submitClaim(
            chainKey,
            blockHeight,
            encodedTransaction,
            merkleRoot,
            siblings,
            lowerEndpointDigest,
            continuityRoots,
            PROTECTION_AMOUNT
        );

        // A second, differently-keyed proof (different height) for the same order.
        _mockCalculateTxIndex(TX_INDEX + 1);

        vm.expectRevert(abi.encodeWithSelector(ClaimVault.OrderAlreadyClaimed.selector, ORDER_ID));
        vault.submitClaim(
            chainKey,
            blockHeight + 1,
            encodedTransaction,
            merkleRoot,
            siblings,
            lowerEndpointDigest,
            continuityRoots,
            PROTECTION_AMOUNT
        );
    }

    // ---------- Sprint 6: fuzz tests (payout formula & pool balance boundaries) ----------

    /// @notice The payout formula is `min(suggestedPayout, protectionAmount, payoutCap)` —
    ///         asserted here as two independent properties rather than by reimplementing
    ///         the contract's own ternary chain: (1) the payout never exceeds any of the
    ///         three bounds, and (2) it's tight — equal to at least one of them, not some
    ///         arbitrarily smaller value a bug could produce while still satisfying (1).
    function testFuzz_SubmitClaim_PayoutIsExactlyTheMinOfSuggestionProtectionAndCap(
        uint256 protectionAmount,
        uint256 payoutCap,
        uint256 suggestedPayout
    ) public {
        protectionAmount = bound(protectionAmount, 0, 1_000_000 ether);
        payoutCap = bound(payoutCap, 0, 1_000_000 ether);
        // suggestedPayout is left at its full uint256 range on purpose — the AI's
        // suggestion is untrusted input (see docs/THREAT_MODEL.md T4/T10) and the
        // formula must hold even at the extremes (0, type(uint256).max).

        vm.deal(address(vault), 1_000_000 ether); // pool balance isn't what this test probes
        vault.setPayoutCap(payoutCap);
        vm.prank(worker);
        vault.registerOrder(ORDER_ID, buyer, protectionAmount);

        (
            uint64 chainKey,
            uint64 blockHeight,
            bytes memory encodedTransaction,
            bytes32 merkleRoot,
            INativeQueryVerifier.MerkleProofEntry[] memory siblings,
            bytes32 lowerEndpointDigest,
            bytes32[] memory continuityRoots
        ) = _validProofArgs();
        _mockCalculateTxIndex(TX_INDEX);
        _mockVerifyAndEmit(true);

        (, uint256 payoutAmount) = vault.submitClaim(
            chainKey,
            blockHeight,
            encodedTransaction,
            merkleRoot,
            siblings,
            lowerEndpointDigest,
            continuityRoots,
            suggestedPayout
        );

        assertLe(payoutAmount, suggestedPayout, "payout must never exceed the AI's suggestion");
        assertLe(payoutAmount, protectionAmount, "payout must never exceed the registered protection amount");
        assertLe(payoutAmount, payoutCap, "payout must never exceed the policy cap");
        assertTrue(
            payoutAmount == suggestedPayout || payoutAmount == protectionAmount || payoutAmount == payoutCap,
            "payout must be tight: equal to whichever bound is actually binding"
        );
    }

    /// @notice If the pool can't cover the payout, submitClaim must revert (never send a
    ///         partial amount) and leave the order unclaimed, so a later top-up lets the
    ///         same proof be resubmitted successfully — the exact scenario hit live during
    ///         Sprint 5's integration testing tonight (an underfunded pool from prior
    ///         payouts), confirmed here as a permanent regression test.
    function testFuzz_SubmitClaim_RevertsWithoutStateChangeWhenPoolBalanceInsufficient(
        uint256 poolBalance,
        uint256 protectionAmount
    ) public {
        protectionAmount = bound(protectionAmount, 1, 1_000_000 ether); // 0 never underfunds
        poolBalance = bound(poolBalance, 0, protectionAmount - 1); // strictly insufficient

        vm.deal(address(vault), poolBalance);
        vault.setPayoutCap(protectionAmount); // cap isn't the binding constraint here
        vm.prank(worker);
        vault.registerOrder(ORDER_ID, buyer, protectionAmount);

        (
            uint64 chainKey,
            uint64 blockHeight,
            bytes memory encodedTransaction,
            bytes32 merkleRoot,
            INativeQueryVerifier.MerkleProofEntry[] memory siblings,
            bytes32 lowerEndpointDigest,
            bytes32[] memory continuityRoots
        ) = _validProofArgs();
        _mockCalculateTxIndex(TX_INDEX);
        _mockVerifyAndEmit(true);

        vm.expectRevert(ClaimVault.PayoutTransferFailed.selector);
        vault.submitClaim(
            chainKey,
            blockHeight,
            encodedTransaction,
            merkleRoot,
            siblings,
            lowerEndpointDigest,
            continuityRoots,
            protectionAmount
        );

        (,, bool claimed) = vault.orders(ORDER_ID);
        assertFalse(claimed, "a reverted payout must not mark the order claimed - it must stay resubmittable");
        assertEq(vault.poolBalance(), poolBalance, "a reverted payout must not move any funds");
    }

    /// @notice A zero payout cap is a valid (if drastic) policy setting — e.g. an owner
    ///         pausing payouts without disabling claim submission entirely. submitClaim
    ///         must still run to completion (mark the order claimed, emit ClaimPaid) with
    ///         a zero amount, not revert or silently skip the claim.
    function test_SubmitClaim_ZeroPayoutCap_PaysNothingButMarksClaimed() public {
        vault.setPayoutCap(0);
        vm.prank(worker);
        vault.registerOrder(ORDER_ID, buyer, PROTECTION_AMOUNT);

        (
            uint64 chainKey,
            uint64 blockHeight,
            bytes memory encodedTransaction,
            bytes32 merkleRoot,
            INativeQueryVerifier.MerkleProofEntry[] memory siblings,
            bytes32 lowerEndpointDigest,
            bytes32[] memory continuityRoots
        ) = _validProofArgs();
        _mockCalculateTxIndex(TX_INDEX);
        _mockVerifyAndEmit(true);

        uint256 buyerBalanceBefore = buyer.balance;

        (, uint256 payoutAmount) = vault.submitClaim(
            chainKey,
            blockHeight,
            encodedTransaction,
            merkleRoot,
            siblings,
            lowerEndpointDigest,
            continuityRoots,
            PROTECTION_AMOUNT
        );

        assertEq(payoutAmount, 0);
        assertEq(buyer.balance, buyerBalanceBefore);
        (,, bool claimed) = vault.orders(ORDER_ID);
        assertTrue(claimed, "the order is consumed even though nothing was paid - it cannot be resubmitted later");
    }

    /// @notice Symmetric to the zero-cap case: the AI suggesting 0% (e.g. a
    ///         parseable-but-fully-discounted response) must also complete normally with
    ///         a zero payout, not revert.
    function test_SubmitClaim_ZeroSuggestedPayout_PaysNothingButMarksClaimed() public {
        vm.prank(worker);
        vault.registerOrder(ORDER_ID, buyer, PROTECTION_AMOUNT);

        (
            uint64 chainKey,
            uint64 blockHeight,
            bytes memory encodedTransaction,
            bytes32 merkleRoot,
            INativeQueryVerifier.MerkleProofEntry[] memory siblings,
            bytes32 lowerEndpointDigest,
            bytes32[] memory continuityRoots
        ) = _validProofArgs();
        _mockCalculateTxIndex(TX_INDEX);
        _mockVerifyAndEmit(true);

        (, uint256 payoutAmount) = vault.submitClaim(
            chainKey, blockHeight, encodedTransaction, merkleRoot, siblings, lowerEndpointDigest, continuityRoots, 0
        );

        assertEq(payoutAmount, 0);
        (,, bool claimed) = vault.orders(ORDER_ID);
        assertTrue(claimed);
    }
}
