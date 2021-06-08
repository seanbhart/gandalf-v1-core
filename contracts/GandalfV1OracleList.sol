// SPDX-License-Identifier: GPL-3.0
pragma solidity =0.8.4;

import './interfaces/IGandalfV1OracleList.sol';
import './library/OpenZeppelin/Ownable.sol';
import './library/Uniswap/AggregatorV3Interface.sol';

contract GandalfV1OracleList is IGandalfV1OracleList, Ownable {
    address[] public tokenAddresses;
    mapping(address => AggregatorV3Interface) public oracleMap;

    // // address ADAUSD = 0xAE48c91dF1fE419994FFDa27da09D5aC69c30f55;
    // address BTCUSD = 0x6135b13325bfC4B00278B4abC5e20bbce2D6580e;
    // address BNBUSD = 0x8993ED705cdf5e84D0a3B754b5Ee0e1783fcdF16;
    // // address DOTUSD = 0x1C07AFb8E2B827c5A4739C6d59Ae3A5035f28734;
    // address ETHUSD = 0x9326BFA02ADD2366b30bacB125260Af641031331;
    // // address FTMETH = 0x2DE7E4a9488488e0058B95854CC2f7955B35dC9b;
    // UNI = 0x9b6Ff80Ff8348852d5281de45E66B7ED36E7B8a9
    // address UNIUSD = 0xDA5904BdBfB4EF12a3955aEcA103F51dc87c7C39;
    // address XRPUSD = 0x3eA2b7e3ed9EA9120c3d6699240d1ff2184AC8b3;

    event OracleAdded(address _tokenAddress, address _oracleAddress);
    event OracleRemoved(address _tokenAddress, address _oracleAddress);

    constructor () {
        // Add default pairs
        // UNI-USD
        oracleMap[0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984] = AggregatorV3Interface(0xDA5904BdBfB4EF12a3955aEcA103F51dc87c7C39);
        tokenAddresses.push(0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984);
        // USDT-USD
        oracleMap[0x07de306FF27a2B630B1141956844eB1552B956B5] = AggregatorV3Interface(0x2ca5A90D34cA333661083F89D831f757A9A50148);
        tokenAddresses.push(0x07de306FF27a2B630B1141956844eB1552B956B5);
        // WETH-USD
        oracleMap[0xd0A1E359811322d97991E03f863a0C30C2cF029C] = AggregatorV3Interface(0x9326BFA02ADD2366b30bacB125260Af641031331);
        tokenAddresses.push(0xd0A1E359811322d97991E03f863a0C30C2cF029C);
    }

    function getLatestPrice(address _tokenAddress) override public view returns (int256) {
        (
            , //uint80 roundID, 
            int256 price,
            , //uint256 startedAt,
            , //uint256 timestamp,
            //uint80 answeredInRound
        ) = oracleMap[_tokenAddress].latestRoundData();
        return price;
    }

    function getTokenAddresses() override external view returns (address[] memory) {
        return tokenAddresses;
    }

    function addOracle(address _tokenAddress, address _oracleAddress) external onlyOwner {
        require(oracleMap[_tokenAddress].decimals() > 0, "addOracle: Token and Oracle pair already exists.");

        // Add the pair
        tokenAddresses.push(_tokenAddress);
        oracleMap[_tokenAddress] = AggregatorV3Interface(_oracleAddress);

        emit OracleAdded(_tokenAddress, _oracleAddress);
    }

    function removeOracle(address _tokenAddress) external onlyOwner {
        AggregatorV3Interface oracleInstance = oracleMap[_tokenAddress];
        require(oracleInstance.decimals() > 0, "removeOracle: Token and Oracle pair does not exist.");

        // Remove the pair
        for (uint8 i = 0; i < tokenAddresses.length; i++) {
            tokenAddresses[i] = tokenAddresses[tokenAddresses.length - 1];
            delete tokenAddresses[tokenAddresses.length - 1];
        }
        delete oracleMap[_tokenAddress];

        emit OracleRemoved(_tokenAddress, address(oracleInstance));
    }
}
