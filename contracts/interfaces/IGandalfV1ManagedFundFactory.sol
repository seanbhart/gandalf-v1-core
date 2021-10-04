// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0;

interface IGandalfV1ManagedFundFactory {
    event FundCreated(address indexed manager, address fund, uint);

    function getFund(address manager, string memory title) external view returns (address fund);
    function fundList(uint) external view returns (address fund);
    function fundListLength() external view returns (uint);

    function createFund(address manager, string memory title, address oracleAddress) external returns (address fund);
}
