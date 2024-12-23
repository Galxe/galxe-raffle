// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "forge-std/Script.sol";
import {BaseScript} from "./utils/Base.s.sol";
import {Raffle} from "../src/Raffle.sol";

contract ParticipateScript is BaseScript {
    function run() external chain broadcaster {
        Raffle raffle = Raffle(0x27f1EDBb0B983DA21121fb59298134ab31Cecf73);
        raffle.participate(
            11382,
            1052489016591343323470769874470399378961347062869,
            20853153,
            hex"a276747f578e6e1df2b90257f5e64d76b67b25254c477458ff920d9809c56d9f49c8e05d733780242f2f86164f335df5b308098fe38aec9f632762482b37ac831b"
        );
    }
}
