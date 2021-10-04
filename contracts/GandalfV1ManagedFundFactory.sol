// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.4;

import './interfaces/IGandalfV1ManagedFundFactory.sol';
import './GandalfV1ManagedSingleFund.sol';

contract GandalfV1ManagedFundFactory is IGandalfV1ManagedFundFactory {
    // address public feeRecipient;
    // address public feeRecipientSetter;

    // mapping(address => address[]) public getFunds;
    mapping(address => mapping(string => address)) public override getFund;
    address[] public override fundList;

    event FundCreated(address indexed manager, string title, address fund, uint256);

    // constructor(address _feeRecipient) {
    //     feeRecipient = _feeRecipient;
    // }

    function fundListLength() external view override returns (uint256) {
        return fundList.length;
    }

    function createFund(address manager, string memory title, address oracleAddress) external override returns (address fund) {
        require(manager != address(0), 'GandalfV1: ZERO_ADDRESS');
        require(getFund[manager][title] == address(0), 'GandalfV1: FUND_EXISTS');

        bytes memory bytecode = type(GandalfV1ManagedSingleFund).creationCode;
        bytes32 salt = keccak256(abi.encodePacked('GandalfV1', manager, title));
        assembly {
            fund := create2(0, add(bytecode, 32), mload(bytecode), salt)
        }
        IGandalfV1ManagedSingleFund(fund).initialize(manager, title, oracleAddress);
        // getFund[manager].push(fund);
        getFund[manager][title] = fund;
        fundList.push(fund);

        emit FundCreated(manager, title, fund, fundList.length);
        return fund;
    }
}
