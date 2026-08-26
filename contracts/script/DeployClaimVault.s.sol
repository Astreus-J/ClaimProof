// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script, console} from "forge-std/Script.sol";
import {ClaimVault} from "../src/ClaimVault.sol";

/// @notice Deploys `ClaimVault` to the Creditcoin CC3 Testnet.
/// @dev Usage: `forge script script/DeployClaimVault.s.sol --rpc-url creditcoin_testnet --broadcast`
///      Reads the deploying wallet's key from `PRIVATE_KEY`, the worker
///      wallet from `WORKER_ADDRESS`, and the initial payout cap (in wei)
///      from `PAYOUT_CAP_WEI` — see contracts/.env.example. Never hardcode
///      a key here.
contract DeployClaimVault is Script {
    function run() external returns (ClaimVault vault) {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address workerAddress = vm.envAddress("WORKER_ADDRESS");
        uint256 payoutCapWei = vm.envUint("PAYOUT_CAP_WEI");

        vm.startBroadcast(deployerPrivateKey);
        vault = new ClaimVault(workerAddress, payoutCapWei);
        vm.stopBroadcast();

        console.log("ClaimVault deployed at:", address(vault));
    }
}
