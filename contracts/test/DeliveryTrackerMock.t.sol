// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {Test} from "forge-std/Test.sol";
import {DeliveryTrackerMock} from "../src/DeliveryTrackerMock.sol";

contract DeliveryTrackerMockTest is Test {
    DeliveryTrackerMock internal tracker;

    address internal buyer = makeAddr("buyer");
    uint256 internal constant ORDER_ID = 1;
    uint256 internal constant SLA_SECONDS = 1 days;

    function setUp() public {
        tracker = new DeliveryTrackerMock();
    }

    function test_CreateShipment_StoresRecordAndEmitsEvent() public {
        vm.expectEmit(true, true, false, true);
        emit DeliveryTrackerMock.ShipmentCreated(ORDER_ID, buyer, block.timestamp + SLA_SECONDS);

        tracker.createShipment(ORDER_ID, buyer, SLA_SECONDS);

        (address storedBuyer, uint256 storedOrderId, uint256 slaDeadline, DeliveryTrackerMock.ShipmentStatus status) =
            tracker.shipments(ORDER_ID);

        assertEq(storedBuyer, buyer);
        assertEq(storedOrderId, ORDER_ID);
        assertEq(slaDeadline, block.timestamp + SLA_SECONDS);
        assertEq(uint8(status), uint8(DeliveryTrackerMock.ShipmentStatus.Pending));
    }

    function test_CreateShipment_RevertsOnDuplicateOrderId() public {
        tracker.createShipment(ORDER_ID, buyer, SLA_SECONDS);

        vm.expectRevert(abi.encodeWithSelector(DeliveryTrackerMock.ShipmentAlreadyExists.selector, ORDER_ID));
        tracker.createShipment(ORDER_ID, buyer, SLA_SECONDS);
    }

    function test_CreateShipment_RevertsOnZeroAddressBuyer() public {
        vm.expectRevert(DeliveryTrackerMock.ZeroAddress.selector);
        tracker.createShipment(ORDER_ID, address(0), SLA_SECONDS);
    }

    function test_ConfirmDelivery_UpdatesStatusBeforeDeadline() public {
        tracker.createShipment(ORDER_ID, buyer, SLA_SECONDS);

        tracker.confirmDelivery(ORDER_ID);

        (,,, DeliveryTrackerMock.ShipmentStatus status) = tracker.shipments(ORDER_ID);
        assertEq(uint8(status), uint8(DeliveryTrackerMock.ShipmentStatus.Delivered));
    }

    function test_ConfirmDelivery_RevertsIfShipmentNotFound() public {
        vm.expectRevert(abi.encodeWithSelector(DeliveryTrackerMock.ShipmentNotFound.selector, ORDER_ID));
        tracker.confirmDelivery(ORDER_ID);
    }

    function test_ConfirmDelivery_RevertsIfAlreadyDelivered() public {
        tracker.createShipment(ORDER_ID, buyer, SLA_SECONDS);
        tracker.confirmDelivery(ORDER_ID);

        vm.expectRevert(abi.encodeWithSelector(DeliveryTrackerMock.ShipmentNotPending.selector, ORDER_ID));
        tracker.confirmDelivery(ORDER_ID);
    }

    function test_ReportDeliveryFailure_RevertsBeforeSlaExpires() public {
        tracker.createShipment(ORDER_ID, buyer, SLA_SECONDS);

        vm.expectRevert(
            abi.encodeWithSelector(DeliveryTrackerMock.SlaNotExpired.selector, ORDER_ID, block.timestamp + SLA_SECONDS)
        );
        tracker.reportDeliveryFailure(ORDER_ID);
    }

    function test_ReportDeliveryFailure_EmitsEventAfterSlaExpires() public {
        tracker.createShipment(ORDER_ID, buyer, SLA_SECONDS);
        vm.warp(block.timestamp + SLA_SECONDS + 1);

        vm.expectEmit(true, true, false, true);
        emit DeliveryTrackerMock.DeliveryFailed(ORDER_ID, buyer, block.timestamp);

        tracker.reportDeliveryFailure(ORDER_ID);

        (,,, DeliveryTrackerMock.ShipmentStatus status) = tracker.shipments(ORDER_ID);
        assertEq(uint8(status), uint8(DeliveryTrackerMock.ShipmentStatus.Failed));
    }

    function test_ReportDeliveryFailure_RevertsIfAlreadyDelivered() public {
        tracker.createShipment(ORDER_ID, buyer, SLA_SECONDS);
        tracker.confirmDelivery(ORDER_ID);
        vm.warp(block.timestamp + SLA_SECONDS + 1);

        vm.expectRevert(abi.encodeWithSelector(DeliveryTrackerMock.ShipmentNotPending.selector, ORDER_ID));
        tracker.reportDeliveryFailure(ORDER_ID);
    }
}
