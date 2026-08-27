// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script, console} from "forge-std/Script.sol";
import {ClaimVault} from "../src/ClaimVault.sol";

/// @notice Deploys `ClaimVault` to the Creditcoin CC3 Testnet.
/// @dev Creditcoin CC3's JSON-RPC block headers omit `mixHash`, so
///      `forge script`'s local EVM simulation fails with "prevrandao not
///      set" (alloy can't deserialize the block). Use `forge create`
///      instead, which broadcasts directly without simulating against a
///      fetched block:
///      `forge create src/ClaimVault.sol:ClaimVault --rpc-url creditcoin_testnet
///        --private-key $PRIVATE_KEY --broadcast
///        --constructor-args $SOURCE_CONTRACT_ADDRESS $WORKER_ADDRESS $PAYOUT_CAP_WEI`
///      This script is kept for reference/local simulation (`forge script
///      script/DeployClaimVault.s.sol` without `--broadcast` still works
///      against a local fork), but `forge create` is what actually deploys
///      to Creditcoin. Reads the deploying wallet's key from `PRIVATE_KEY`,
///      the trusted DeliveryTrackerMock address from `SOURCE_CONTRACT_ADDRESS`,
///      the worker wallet from `WORKER_ADDRESS`, and the initial payout cap
///      (in wei) from `PAYOUT_CAP_WEI` -- see contracts/.env.example. Never
///      hardcode a key here.
contract DeployClaimVault is Script {
    function run() external returns (ClaimVault vault) {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address sourceContractAddress = vm.envAddress("SOURCE_CONTRACT_ADDRESS");
        address workerAddress = vm.envAddress("WORKER_ADDRESS");
        uint256 payoutCapWei = vm.envUint("PAYOUT_CAP_WEI");

        vm.startBroadcast(deployerPrivateKey);
        vault = new ClaimVault(sourceContractAddress, workerAddress, payoutCapWei);
        vm.stopBroadcast();

        console.log("ClaimVault deployed at:", address(vault));
    }
}
