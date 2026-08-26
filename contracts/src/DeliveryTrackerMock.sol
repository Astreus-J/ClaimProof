// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

/// @title DeliveryTrackerMock
/// @notice Minimal, team-controlled delivery-tracking contract deployed on
///         Ethereum Sepolia. It emits the `DeliveryFailed` event that
///         `ClaimVault` on Creditcoin re-verifies via the Attestcoin
///         Protocol before authorizing a payout. This mock deliberately
///         replaces a real logistics oracle for the hackathon MVP so the
///         trigger event can be simulated on demand without depending on a
///         third-party integration (see docs/architecture.md, "DO NOT DO").
contract DeliveryTrackerMock {
    /// @notice Lifecycle states of a tracked shipment.
    enum ShipmentStatus {
        Pending,
        Delivered,
        Failed
    }

    /// @notice A single tracked shipment.
    struct Shipment {
        address buyer;
        uint256 orderId;
        uint256 slaDeadline;
        ShipmentStatus status;
    }

    /// @notice orderId => shipment record.
    mapping(uint256 => Shipment) public shipments;

    /// @notice Emitted when a new shipment is registered.
    event ShipmentCreated(uint256 indexed orderId, address indexed buyer, uint256 slaDeadline);

    /// @notice Emitted when a shipment's SLA expires without delivery confirmation.
    ///         This is the trigger event ClaimVault verifies on Creditcoin via Attestcoin.
    event DeliveryFailed(uint256 indexed orderId, address indexed buyer, uint256 timestamp);

    /// @notice Emitted when a shipment is confirmed delivered before its SLA deadline.
    event DeliveryConfirmed(uint256 indexed orderId, address indexed buyer, uint256 timestamp);

    /// @notice Thrown when `createShipment` is called with an `orderId` already in use.
    error ShipmentAlreadyExists(uint256 orderId);

    /// @notice Thrown when an operation targets an `orderId` with no shipment record.
    error ShipmentNotFound(uint256 orderId);

    /// @notice Thrown when an operation requires a shipment still in `Pending` status.
    error ShipmentNotPending(uint256 orderId);

    /// @notice Thrown when `reportDeliveryFailure` is called before the SLA deadline.
    error SlaNotExpired(uint256 orderId, uint256 slaDeadline);

    /// @notice Thrown when a zero address is supplied where a buyer address is required.
    error ZeroAddress();

    /// @notice Registers a new shipment with a delivery SLA.
    /// @param orderId Unique identifier for the order, shared with `ClaimVault` on Creditcoin.
    /// @param buyer Address entitled to the delivery-protection payout.
    /// @param slaSeconds Seconds from now within which delivery must be confirmed.
    function createShipment(uint256 orderId, address buyer, uint256 slaSeconds) external {
        if (shipments[orderId].buyer != address(0)) revert ShipmentAlreadyExists(orderId);
        if (buyer == address(0)) revert ZeroAddress();

        uint256 slaDeadline = block.timestamp + slaSeconds;
        shipments[orderId] =
            Shipment({buyer: buyer, orderId: orderId, slaDeadline: slaDeadline, status: ShipmentStatus.Pending});

        emit ShipmentCreated(orderId, buyer, slaDeadline);
    }

    /// @notice Confirms successful delivery before the SLA deadline.
    /// @param orderId The shipment to confirm.
    function confirmDelivery(uint256 orderId) external {
        Shipment storage shipment = shipments[orderId];
        if (shipment.buyer == address(0)) revert ShipmentNotFound(orderId);
        if (shipment.status != ShipmentStatus.Pending) revert ShipmentNotPending(orderId);

        shipment.status = ShipmentStatus.Delivered;
        emit DeliveryConfirmed(orderId, shipment.buyer, block.timestamp);
    }

    /// @notice Marks a shipment as failed once its SLA has expired, emitting
    ///         the event `ClaimVault` verifies via Attestcoin. Callable by
    ///         anyone once the deadline has passed — this mock has no access
    ///         control because it is a deliberately simplified stand-in for a
    ///         real logistics oracle, and is also how the demo simulates a
    ///         delivery failure on cue (see docs/execution.md).
    /// @param orderId The shipment to mark failed.
    function reportDeliveryFailure(uint256 orderId) external {
        Shipment storage shipment = shipments[orderId];
        if (shipment.buyer == address(0)) revert ShipmentNotFound(orderId);
        if (shipment.status != ShipmentStatus.Pending) revert ShipmentNotPending(orderId);
        if (block.timestamp < shipment.slaDeadline) revert SlaNotExpired(orderId, shipment.slaDeadline);

        shipment.status = ShipmentStatus.Failed;
        emit DeliveryFailed(orderId, shipment.buyer, block.timestamp);
    }
}
