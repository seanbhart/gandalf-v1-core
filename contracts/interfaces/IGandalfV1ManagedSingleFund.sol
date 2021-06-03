// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0;

interface IGandalfV1ManagedSingleFund {
    function manager() external view returns (address);
    function title() external view returns (string memory);

    function initialize(address _manager, string memory _title) external;
}
