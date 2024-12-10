// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "forge-std/Script.sol";
import {BaseScript} from "../utils/Base.s.sol";
import {Raffle} from "../../src/Raffle.sol";

contract RaffleScript is BaseScript {
    function run() external chain broadcaster {
        bytes32 CREATE2_SALT = vm.envBytes32("CREATE2_SALT");
        address owner = vm.envAddress("OWNER");
        address signer = vm.envAddress("SIGNER");
        address verifier = vm.envAddress("VERIFIER");
        bytes32 vkey = vm.envBytes32("VKEY");
        address drandOracle = vm.envAddress("DRAND_ORACLE");

        console.log("Owner:", owner);
        console.log("Signer:", signer);
        console.log("Verifier:", verifier);
        console.log("VKey:");
        console.logBytes32(vkey);
        console.log("DrandOracle:", drandOracle);

        Raffle raffle = new Raffle(owner, signer, verifier, vkey, drandOracle);
        console.log("Raffle deployed at:", address(raffle));

        // Write address
        writeAddress("RAFFLE", address(raffle));
        writeAddress("VERIFIER", verifier);
        writeAddress("DRAND_ORACLE", drandOracle);
        writeBytes32("VKEY", vkey);
    }
}
