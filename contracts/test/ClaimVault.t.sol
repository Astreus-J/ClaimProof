// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Test} from "forge-std/Test.sol";
import {ClaimVault} from "../src/ClaimVault.sol";

contract ClaimVaultTest is Test {
    ClaimVault internal vault;

    address internal owner = address(this);
    address internal worker = makeAddr("worker");
    address internal buyer = makeAddr("buyer");
    uint256 internal constant PAYOUT_CAP = 1 ether;
    uint256 internal constant ORDER_ID = 1;
    uint256 internal constant PROTECTION_AMOUNT = 0.1 ether;

    function setUp() public {
        vault = new ClaimVault(worker, PAYOUT_CAP);
    }

    function test_Constructor_SetsOwnerWorkerAndCap() public view {
        assertEq(vault.owner(), owner);
        assertEq(vault.worker(), worker);
        assertEq(vault.payoutCap(), PAYOUT_CAP);
    }

    function test_Constructor_RevertsOnZeroAddressWorker() public {
        vm.expectRevert(ClaimVault.ZeroAddress.selector);
        new ClaimVault(address(0), PAYOUT_CAP);
    }

    function test_RegisterOrder_StoresRecordAndEmitsEvent() public {
        vm.expectEmit(true, true, false, true);
        emit ClaimVault.OrderRegistered(ORDER_ID, buyer, PROTECTION_AMOUNT);

        vault.registerOrder(ORDER_ID, buyer, PROTECTION_AMOUNT);

        (address storedBuyer, uint256 storedAmount, bool claimed) = vault.orders(ORDER_ID);
        assertEq(storedBuyer, buyer);
        assertEq(storedAmount, PROTECTION_AMOUNT);
        assertFalse(claimed);
    }

    function test_RegisterOrder_RevertsOnDuplicateOrderId() public {
        vault.registerOrder(ORDER_ID, buyer, PROTECTION_AMOUNT);

        vm.expectRevert(abi.encodeWithSelector(ClaimVault.OrderAlreadyExists.selector, ORDER_ID));
        vault.registerOrder(ORDER_ID, buyer, PROTECTION_AMOUNT);
    }

    function test_RegisterOrder_RevertsOnZeroAddressBuyer() public {
        vm.expectRevert(ClaimVault.ZeroAddress.selector);
        vault.registerOrder(ORDER_ID, address(0), PROTECTION_AMOUNT);
    }

    function test_FundPool_IncreasesBalanceAndEmitsEvent() public {
        address funder = makeAddr("funder");
        vm.deal(funder, 1 ether);

        vm.expectEmit(true, false, false, true);
        emit ClaimVault.PoolFunded(funder, 1 ether);

        vm.prank(funder);
        vault.fundPool{value: 1 ether}();

        assertEq(vault.poolBalance(), 1 ether);
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
}
