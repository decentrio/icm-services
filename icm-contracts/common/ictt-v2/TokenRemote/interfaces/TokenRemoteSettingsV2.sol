// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice Settings for constructing an ICTT v2 {ITokenRemote} contract.
 * @param blockchainID The blockchain ID of the chain the remote is deployed on.
 * @param teleporterMessenger The messenger used to send to and receive from the token home's chain.
 * @param teleporterManager Address that manages which messengers this contract uses for each chain.
 * @param tokenHomeBlockchainID The blockchain ID of the associated TokenHome instance.
 * @param tokenHomeAddress The address of the associated token home contract.
 * @param tokenHomeDecimals The number of decimal places used by the token home's token.
 */
struct TokenRemoteSettingsV2 {
    bytes32 blockchainID;
    address teleporterMessenger;
    address teleporterManager;
    bytes32 tokenHomeBlockchainID;
    address tokenHomeAddress;
    uint8 tokenHomeDecimals;
}
