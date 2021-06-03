// SPDX-License-Identifier: GPL-3.0
pragma solidity =0.8.4;

import './interfaces/IGandalfV1ManagedSingleFund.sol';

contract GandalfV1ManagedSingleFund is IGandalfV1ManagedSingleFund {
    address public factory;
    address public override manager;
    string public override title;

    constructor() {
        factory = msg.sender;
    }

    // called once by the factory at time of deployment
    function initialize(address _manager, string memory _title) external override {
        require(msg.sender == factory, 'GandalfV1: FORBIDDEN');
        manager = _manager;
        title = _title;
    }
}
