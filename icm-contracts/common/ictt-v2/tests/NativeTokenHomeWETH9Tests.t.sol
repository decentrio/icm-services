// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {NativeTokenHomeTest} from "./NativeTokenHomeTests.t.sol";
import {NativeTokenHomeUpgradeable} from "../TokenHome/NativeTokenHomeUpgradeable.sol";
import {IWrappedNativeToken} from "../../../avalanche/ictt/interfaces/IWrappedNativeToken.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {TransparentUpgradeableProxy} from
    "@openzeppelin/contracts@5.1.0/proxy/transparent/TransparentUpgradeableProxy.sol";

/**
 * @dev WETH9 as deployed on Ethereum mainnet, ported to Solidity 0.8. Unlike {WrappedNativeToken},
 * `withdraw` pays out with `transfer`, which forwards only the 2300 gas stipend to the receiver.
 * Its requires keep WETH9's empty revert data.
 */
// solhint-disable reason-string
contract WETH9 {
    string public name = "Wrapped Ether";
    string public symbol = "WETH";
    uint8 public decimals = 18;

    mapping(address account => uint256 balance) public balanceOf;
    mapping(address owner => mapping(address spender => uint256 amount)) public allowance;

    event Approval(address indexed src, address indexed guy, uint256 wad);
    event Transfer(address indexed src, address indexed dst, uint256 wad);
    event Deposit(address indexed dst, uint256 wad);
    event Withdrawal(address indexed src, uint256 wad);

    receive() external payable {
        deposit();
    }

    function withdraw(
        uint256 wad
    ) external {
        require(balanceOf[msg.sender] >= wad);
        balanceOf[msg.sender] -= wad;
        payable(msg.sender).transfer(wad);
        emit Withdrawal(msg.sender, wad);
    }

    function approve(address guy, uint256 wad) external returns (bool) {
        allowance[msg.sender][guy] = wad;
        emit Approval(msg.sender, guy, wad);
        return true;
    }

    function transfer(address dst, uint256 wad) external returns (bool) {
        return transferFrom(msg.sender, dst, wad);
    }

    function totalSupply() external view returns (uint256) {
        return address(this).balance;
    }

    function deposit() public payable {
        balanceOf[msg.sender] += msg.value;
        emit Deposit(msg.sender, msg.value);
    }

    function transferFrom(address src, address dst, uint256 wad) public returns (bool) {
        require(balanceOf[src] >= wad);
        if (src != msg.sender && allowance[src][msg.sender] != type(uint256).max) {
            require(allowance[src][msg.sender] >= wad);
            allowance[src][msg.sender] -= wad;
        }
        balanceOf[src] -= wad;
        balanceOf[dst] += wad;
        emit Transfer(src, dst, wad);
        return true;
    }
}
// solhint-enable reason-string

/**
 * @dev Runs the NativeTokenHome suite against WETH9, the wrapped token a home on Ethereum uses.
 */
contract NativeTokenHomeWETH9Test is NativeTokenHomeTest {
    function _newWrappedToken() internal override returns (IWrappedNativeToken) {
        return IWrappedNativeToken(address(new WETH9()));
    }

    /**
     * @dev WETH9 reverts without data when an allowance is too low.
     */
    function _insufficientAllowanceError(
        address,
        uint256,
        uint256
    ) internal pure override returns (bytes memory) {
        return "";
    }
}

/**
 * @dev Runs the WETH9 suite with the home behind a {TransparentUpgradeableProxy}, so that WETH9's
 * 2300 gas stipend also has to cover the proxy's delegation into {NativeTokenHomeUpgradeable-receive}.
 */
contract NativeTokenHomeWETH9ProxyTest is NativeTokenHomeWETH9Test {
    address private constant _PROXY_ADMIN_OWNER = address(0xAD);

    function setUp() public override {
        super.setUp();
        NativeTokenHomeUpgradeable implementation =
            new NativeTokenHomeUpgradeable(ICMInitializable.Disallowed);
        app = NativeTokenHomeUpgradeable(
            payable(
                address(
                    new TransparentUpgradeableProxy(
                        address(implementation),
                        _PROXY_ADMIN_OWNER,
                        abi.encodeCall(
                            NativeTokenHomeUpgradeable.initialize,
                            (
                                DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
                                MOCK_TELEPORTER_MESSENGER_ADDRESS,
                                address(wavax)
                            )
                        )
                    )
                )
            )
        );
        tokenHome = app;
        nativeTokenTransferrer = app;
        tokenTransferrer = app;
        _routeDefaultRemoteChains();
    }
}
