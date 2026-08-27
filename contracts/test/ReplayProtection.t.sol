// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {INativeQueryVerifier} from "../src/interfaces/INativeQueryVerifier.sol";
import {EvmV1Decoder} from "usc-contracts/decoding/EvmV1Decoder.sol";
import {ClaimVault} from "../src/ClaimVault.sol";
import {AttestcoinFixtures} from "./helpers/AttestcoinFixtures.sol";

/// @notice Dedicated coverage for T2 (docs/THREAT_MODEL.md): resubmitting an
///         already-processed Attestcoin proof must never authorize a second payout.
contract ReplayProtectionTest is AttestcoinFixtures {
    ClaimVault internal vault;

    address internal worker = makeAddr("worker");
    address internal buyer = makeAddr("buyer");
    address internal sourceContract = makeAddr("deliveryTrackerMock");
    uint256 internal constant PAYOUT_CAP = 1 ether;
    uint256 internal constant ORDER_ID = 1;
    uint256 internal constant PROTECTION_AMOUNT = 0.1 ether;

    uint64 internal constant CHAIN_KEY = 1;
    uint64 internal constant BLOCK_HEIGHT = 100;
    uint64 internal constant TX_INDEX = 0;
    bytes32 internal constant MERKLE_ROOT = bytes32(uint256(0xabc));
    bytes32 internal constant LOWER_ENDPOINT_DIGEST = bytes32(uint256(0xdef));

    function setUp() public {
        vault = new ClaimVault(sourceContract, worker, PAYOUT_CAP);
        vm.deal(address(vault), 10 ether);
        vm.prank(worker);
        vault.registerOrder(ORDER_ID, buyer, PROTECTION_AMOUNT);
    }

    function _submit(bytes memory encodedTransaction) internal returns (uint256 orderId, uint256 payoutAmount) {
        return vault.submitClaim(
            CHAIN_KEY,
            BLOCK_HEIGHT,
            encodedTransaction,
            MERKLE_ROOT,
            new INativeQueryVerifier.MerkleProofEntry[](0),
            LOWER_ENDPOINT_DIGEST,
            new bytes32[](0),
            PROTECTION_AMOUNT
        );
    }

    function test_ResubmittingSameProof_RevertsWithoutSecondPayout() public {
        EvmV1Decoder.LogEntryTuple[] memory logs = new EvmV1Decoder.LogEntryTuple[](1);
        logs[0] = _deliveryFailedLog(sourceContract, ORDER_ID, buyer, block.timestamp);
        bytes memory encodedTransaction = _encodeTransaction(1, logs);

        _mockCalculateTxIndex(TX_INDEX);
        _mockVerifyAndEmit(true);

        _submit(encodedTransaction);
        uint256 buyerBalanceAfterFirstClaim = buyer.balance;

        bytes32 queryId = keccak256(abi.encodePacked(CHAIN_KEY, BLOCK_HEIGHT, TX_INDEX));
        assertTrue(vault.processedQueries(queryId));

        vm.expectRevert(abi.encodeWithSelector(ClaimVault.QueryAlreadyProcessed.selector, queryId));
        _submit(encodedTransaction);

        assertEq(buyer.balance, buyerBalanceAfterFirstClaim, "replay must not move funds a second time");
    }

    function test_ResubmittingSameProof_RevertsBeforeReVerifyingWithPrecompile() public {
        EvmV1Decoder.LogEntryTuple[] memory logs = new EvmV1Decoder.LogEntryTuple[](1);
        logs[0] = _deliveryFailedLog(sourceContract, ORDER_ID, buyer, block.timestamp);
        bytes memory encodedTransaction = _encodeTransaction(1, logs);

        _mockCalculateTxIndex(TX_INDEX);
        _mockVerifyAndEmit(true);
        _submit(encodedTransaction);

        // Even if the precompile were to (incorrectly) start rejecting the proof on a
        // second look, replay protection must reject it first, before verification runs.
        _mockVerifyAndEmit(false);

        bytes32 queryId = keccak256(abi.encodePacked(CHAIN_KEY, BLOCK_HEIGHT, TX_INDEX));
        vm.expectRevert(abi.encodeWithSelector(ClaimVault.QueryAlreadyProcessed.selector, queryId));
        _submit(encodedTransaction);
    }
}
