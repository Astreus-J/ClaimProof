// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Script, console} from "forge-std/Script.sol";
import {DeliveryTrackerMock} from "../src/DeliveryTrackerMock.sol";

/// @notice Deploys `DeliveryTrackerMock` to Ethereum Sepolia.
/// @dev Usage: `forge script script/DeployDeliveryTrackerMock.s.sol --rpc-url sepolia --broadcast --verify`
///      Reads the deploying wallet's key from the `PRIVATE_KEY` environment
///      variable (see contracts/.env.example) — never hardcode a key here.
contract DeployDeliveryTrackerMock is Script {
    function run() external returns (DeliveryTrackerMock tracker) {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");

        vm.startBroadcast(deployerPrivateKey);
        tracker = new DeliveryTrackerMock();
        vm.stopBroadcast();

        console.log("DeliveryTrackerMock deployed at:", address(tracker));
    }
}
