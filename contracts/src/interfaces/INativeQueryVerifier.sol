// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

/// @title INativeQueryVerifier
/// @notice Interface for the Attestcoin Protocol's Block Prover precompile, a native
///         runtime precompile on Creditcoin that verifies a source-chain transaction's
///         inclusion in a finalized, attested block via a Merkle proof plus a continuity
///         proof chaining block digests. This interface definition is the standard,
///         self-contained boilerplate every Attestcoin dApp declares itself — Gluwa
///         does not publish it as an installable package, only as reference source in
///         `gluwa/attestcoin-protocol-examples`. See docs/ATTESTCOIN_INTEGRATION.md.
interface INativeQueryVerifier {
    /// @notice One sibling hash on the path from a transaction leaf to the Merkle root.
    struct MerkleProofEntry {
        bytes32 hash;
        bool isLeft;
    }

    /// @notice The Merkle inclusion proof for one transaction within its block.
    struct MerkleProof {
        bytes32 root;
        MerkleProofEntry[] siblings;
    }

    /// @notice Chains block digests so a single attested checkpoint cryptographically
    ///         anchors a whole range of blocks back to the transaction's block.
    struct ContinuityProof {
        bytes32 lowerEndpointDigest;
        bytes32[] roots;
    }

    /// @notice Verifies inclusion of `encodedTransaction` at the given source chain and
    ///         height. Proves inclusion only — callers must separately check the decoded
    ///         transaction's receipt status field for success.
    /// @param chainKey Attestcoin's identifier for the source chain.
    /// @param height Block height on the source chain containing the transaction.
    /// @param encodedTransaction The ABI-encoded transaction + receipt data to verify.
    /// @param merkleProof Merkle inclusion proof for the transaction within its block.
    /// @param continuityProof Continuity proof anchoring the block to an attested checkpoint.
    /// @return success True if the proof verifies against the protocol's attested state.
    function verifyAndEmit(
        uint64 chainKey,
        uint64 height,
        bytes calldata encodedTransaction,
        MerkleProof calldata merkleProof,
        ContinuityProof calldata continuityProof
    ) external returns (bool success);

    /// @notice Derives the transaction's index within its block from its Merkle proof,
    ///         used to build a collision-resistant anti-replay key.
    function calculateTxIndex(MerkleProof calldata merkleProof) external view returns (uint64);
}

/// @notice Helper for addressing the Native Query Verifier precompile.
library NativeQueryVerifierLib {
    /// @dev Fixed runtime precompile address on Creditcoin (decimal 4050).
    address constant PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000FD2;

    function getVerifier() internal pure returns (INativeQueryVerifier) {
        return INativeQueryVerifier(PRECOMPILE_ADDRESS);
    }
}
