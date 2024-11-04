// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "forge-std/Vm.sol";
import "forge-std/console.sol";
import {stdJson} from "forge-std/StdJson.sol";
import {Script} from "forge-std/Script.sol";

/// @notice Script to inherit from to get access to helper functions for deployments.
abstract contract BaseScript is Script {
    using stdJson for string;

    /// @notice Run the command with the `--broadcast` flag to send the transaction to the chain,
    /// otherwise just simulate the transaction execution.
    modifier broadcaster() {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        console.log("Deployer: %s", vm.addr(deployerPrivateKey));
        vm.startBroadcast(deployerPrivateKey);
        _;
        vm.stopBroadcast();
    }

    /// @notice Runs the script on the chain specified in the `CHAIN` env variable.
    /// Must have a `RPC_${CHAIN}` env variable set for the chain (e.g. RPC_MAINNET).
    modifier chain() {
        string memory c = vm.envString("CHAIN");

        // Switch to the chain using the RPC
        vm.createSelectFork(c);

        console.log("Running script on %s", c);

        _;
    }
}
