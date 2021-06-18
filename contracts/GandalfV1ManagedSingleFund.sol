// SPDX-License-Identifier: GPL-3.0
pragma solidity =0.8.4;

import './interfaces/IGandalfV1ManagedSingleFund.sol';
import './interfaces/IGandalfV1OracleList.sol';
import './library/OpenZeppelin/IERC20.sol';
import './library/Uniswap/IUniswapV2Router02.sol';

contract GandalfV1ManagedSingleFund is IGandalfV1ManagedSingleFund {
    address public factory;
    address public override manager;
    string public override title;

    event TokenSwap(address tokenIn, uint256 amountIn, address tokenOut, uint256 amountOut);

    IGandalfV1OracleList internal oracleList;
    IUniswapV2Router02 internal router;
    address kovanRouter = 0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D;
    address kovanUSDT = 0x07de306FF27a2B630B1141956844eB1552B956B5;
    address kovanUNI = 0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984;
    address kovanWETH = 0xd0A1E359811322d97991E03f863a0C30C2cF029C;

    IERC20 public inToken;
    uint256 public inTokenAllowance;

    constructor() {
        factory = msg.sender;
        oracleList = IGandalfV1OracleList(0x9A5811ed25F34834e1451869194180feAd119dD2);
        router = IUniswapV2Router02(kovanRouter);
        inToken = IERC20(kovanUSDT);
    }

    // called once by the factory at time of deployment
    function initialize(address _manager, string memory _title) external override {
        require(msg.sender == factory, 'GandalfV1: FORBIDDEN');
        manager = _manager;
        title = _title;
    }

    function getLatestPrice(address _tokenAddress) public view returns (int256) {
        int256 px = oracleList.getLatestPrice(_tokenAddress);
        // emit PriceRead(_tokenAddress, px);
        return px;
    }

    function getOracleTokens() public view returns (address[] memory) {
        return oracleList.getTokenAddresses();
    }

    function checkAllowance() external {
        // inToken = IERC20(kovanUSDT);
        uint256 allowance = inToken.allowance(msg.sender, address(this));
        inTokenAllowance = allowance;
    }
    
    function escrowTokens() external {
        // approve & transfer 10 USDT from user to this contract
        // inToken = IERC20(kovanUSDT);
        // require(inToken.approve(address(this), uint256(10000000)) == true);
        // inToken.transfer(address(this), uint256(1000000));
        inToken.transferFrom(msg.sender, address(this), uint256(1000000));
    }

    function swapTokens() external {
        address tokenIn = 0xd0A1E359811322d97991E03f863a0C30C2cF029C;
        address tokenOut = 0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984;

        // Set the amountIn allowance for the router
        // IERC20 inToken = IERC20(0xd0A1E359811322d97991E03f863a0C30C2cF029C);
        inToken.approve(kovanRouter, uint256(100000000000000000));

        // Place the swap transaction
        uint256 amountIn = 100000000000000000;
        uint256 amountOutMin = 75000000000000000;
        address[] memory path; // in -> out // ETH -> UNI
        path[0] = tokenIn;
        path[1] = tokenOut;
        address to = address(this);
        uint256 deadline = 1623019416;

        router.swapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn, amountOutMin, path, to, deadline);
        emit TokenSwap(tokenIn, amountIn, tokenOut, amountOutMin);
    }
}


IUniswapV2Router02 internal router;
    IGandalfV1OracleList internal oracleList;
    
    address[] public tokenAddresses;
    mapping(address => IERC20) public tokenMap; //Store instantiated tokens to reduce fees
    
    address kovanRouter = 0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D;
    address kovanUSDT = 0x07de306FF27a2B630B1141956844eB1552B956B5;
    address kovanUNI = 0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984;
    address kovanWETH = 0xd0A1E359811322d97991E03f863a0C30C2cF029C;
    
    address account1 = 0x48766749B7a0F3635757C84108050e4EDA72a0E9;
    address account2 = 0xc804A95ebF427a702EF624cd8d61A8a4D61B8CaE;
    
    // IERC20 public inToken;
    // uint256 public inTokenAllowance;
    
    event SwappedTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline);
    
    constructor() {
        oracleList = IGandalfV1OracleList(0x1B47696F1908C3815296bED7B747B7425f615f12);
        router = IUniswapV2Router02(kovanRouter);
        // inToken = IERC20(kovanWETH);
        
        // UNI-USD
        tokenMap[0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984] = IERC20(0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984);
        tokenAddresses.push(0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984);
        // USDT-USD
        tokenMap[0x07de306FF27a2B630B1141956844eB1552B956B5] = IERC20(0x07de306FF27a2B630B1141956844eB1552B956B5);
        tokenAddresses.push(0x07de306FF27a2B630B1141956844eB1552B956B5);
        // WETH-USD
        tokenMap[0xd0A1E359811322d97991E03f863a0C30C2cF029C] = IERC20(0xd0A1E359811322d97991E03f863a0C30C2cF029C);
        tokenAddresses.push(0xd0A1E359811322d97991E03f863a0C30C2cF029C);
    }
    
    function checkAllowance(address tokenAddress) public view returns (uint256) {
        // inToken = IERC20(kovanUSDT);
        return tokenMap[tokenAddress].allowance(msg.sender, address(this));
    }
    
    function checkOracle(address tokenAddress) public view returns (int256) {
        // inToken = IERC20(kovanUSDT);
        return oracleList.getLatestPrice(tokenAddress);
    }
    
    function escrowTokens(address tokenAddress, uint256 amount) external {
        // Transfer from user to this contract
        // IMPORTANT: amount should already have been approved directly from client
        tokenMap[tokenAddress].transferFrom(msg.sender, address(this), amount); //uint256(100000000000000000)
    }

    function swapTokens(address fromTokenAddress, address toTokenAddress, uint256 amount) external {
        // Set the amountIn allowance for the router
        tokenMap[fromTokenAddress].approve(kovanRouter, amount); //uint256(1000000000000000000)
        
        // To calculate token equivalency, token decimals, oracle price, and oracle decimals
        // are needed for the input and output tokens.
        (int256 fromTokenOraclePrice, uint8 fromTokenOracleDecimals) = oracleList.getLatestPrice(fromTokenAddress);
        uint8 fromTokenDecimals = tokenMap[fromTokenAddress].decimals();
        (int256 toTokenOraclePrice, uint8 toTokenOracleDecimals) = oracleList.getLatestPrice(toTokenAddress);
        uint8 toTokenDecimals = tokenMap[toTokenAddress].decimals();
        require(fromTokenOraclePrice > 0, 'Oracle prices show negative values.');
        require(toTokenOraclePrice > 0, 'Oracle prices show negative values.');
        
        // Correct the oracle prices to 18 decimal places
        uint256 correctedFromTokenOraclePrice = uint256(fromTokenOraclePrice);
        if (fromTokenOracleDecimals < uint8(18)) {
            correctedFromTokenOraclePrice = uint256(fromTokenOraclePrice) * uint256(10)**uint256(uint8(18) - fromTokenDecimals);
        }
        uint256 correctedToTokenOraclePrice = uint256(toTokenOraclePrice);
        if (toTokenOracleDecimals < uint8(18)) {
            correctedToTokenOraclePrice = uint256(toTokenOraclePrice) * uint256(10)**uint256(uint8(18) - toTokenDecimals);
        }
        uint256 exactToTokenExpectedAmount = (correctedFromTokenOraclePrice / correctedToTokenOraclePrice) * (amount / uint256(10)**uint256(fromTokenDecimals)) * uint256(10)**uint256(toTokenDecimals);
        
        
        // Place the swap transaction
        uint256 amountIn = amount;
        uint256 amountOutMin = exactToTokenExpectedAmount - (((exactToTokenExpectedAmount * 100) / 15) / 100);
        address[] memory path = new address[](2); // in -> out
        path[0] = fromTokenAddress;
        path[1] = toTokenAddress;
        address to = address(this);
        uint256 deadline = uint256(1623289214);

        router.swapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn, amountOutMin, path, to, deadline);
        emit SwappedTokens(amountIn, amountOutMin, exactToTokenExpectedAmount, path, to, deadline);
    }