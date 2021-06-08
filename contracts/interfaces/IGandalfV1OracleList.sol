// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0;

interface IGandalfV1OracleList {
    function getTokenAddresses() external view returns (address[] memory);
    function getLatestPrice(address _tokenAddress) external view returns (int256);
}
