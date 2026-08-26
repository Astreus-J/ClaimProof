// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Test} from "forge-std/Test.sol";
import {INativeQueryVerifier, NativeQueryVerifierLib} from "../../src/interfaces/INativeQueryVerifier.sol";
import {EvmV1Decoder} from "usc-contracts/decoding/EvmV1Decoder.sol";

/// @notice Shared fixtures for tests that exercise ClaimVault.submitClaim against the
///         Attestcoin precompile. The precompile only exists on Creditcoin's runtime, so
///         local Foundry tests fabricate an EvmV1Decoder-compatible encoded transaction
///         and stub the precompile's two calls — this exercises ClaimVault's real
///         decode/verification logic against a realistic byte format; only the
///         precompile's own cryptographic check is mocked, since it cannot run locally.
abstract contract AttestcoinFixtures is Test {
    bytes32 internal constant DELIVERY_FAILED_EVENT_SIGNATURE = keccak256("DeliveryFailed(uint256,address,uint256)");

    /// @notice Stubs the precompile's verifyAndEmit to accept or reject the proof.
    function _mockVerifyAndEmit(bool result) internal {
        vm.mockCall(
            NativeQueryVerifierLib.PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(INativeQueryVerifier.verifyAndEmit.selector),
            abi.encode(result)
        );
    }

    /// @notice Stubs the precompile's transaction-index derivation (used to build the
    ///         anti-replay key before verification even runs).
    function _mockCalculateTxIndex(uint64 txIndex) internal {
        vm.mockCall(
            NativeQueryVerifierLib.PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(INativeQueryVerifier.calculateTxIndex.selector),
            abi.encode(txIndex)
        );
    }

    /// @notice Builds a single DeliveryFailed log entry exactly as DeliveryTrackerMock emits it.
    function _deliveryFailedLog(address emitter, uint256 orderId, address buyer, uint256 timestamp)
        internal
        pure
        returns (EvmV1Decoder.LogEntryTuple memory)
    {
        bytes32[] memory topics = new bytes32[](3);
        topics[0] = DELIVERY_FAILED_EVENT_SIGNATURE;
        topics[1] = bytes32(orderId);
        topics[2] = bytes32(uint256(uint160(buyer)));
        return EvmV1Decoder.LogEntryTuple({address_: emitter, topics: topics, data: abi.encode(timestamp)});
    }

    /// @notice Builds an EvmV1Decoder-compatible encoded transaction with the given
    ///         receipt status and logs. Only the receipt chunk's content matters for
    ///         ClaimVault — the common-tx and type-specific chunks just need to be
    ///         present; EvmV1Decoder.decodeReceiptFields never decodes their content.
    function _encodeTransaction(uint8 receiptStatus, EvmV1Decoder.LogEntryTuple[] memory logs)
        internal
        pure
        returns (bytes memory)
    {
        bytes[] memory chunks = new bytes[](3);
        chunks[0] = hex"00";
        chunks[1] = hex"00";
        chunks[2] = abi.encode(receiptStatus, uint64(21000), logs, bytes(""));
        return abi.encode(uint8(2), chunks); // type 2 (EIP-1559): receipt lives at chunks[2]
    }
}
