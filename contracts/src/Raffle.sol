// SPDX-License-Identifier: Apache-2.0

/*
 Copyright 2024 Galxe.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
 */
pragma solidity ^0.8.24;

import {Pausable} from "@openzeppelin/contracts/utils/Pausable.sol";
import {Ownable2Step} from "@openzeppelin/contracts/access/Ownable2Step.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {EIP712} from "@openzeppelin/contracts/utils/cryptography/EIP712.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {BitMaps} from "@openzeppelin/contracts/utils/structs/BitMaps.sol";
import {ISP1Verifier} from "@sp1-contracts/ISP1Verifier.sol";
import {IRaffle} from "./IRaffle.sol";

/// @title Galxe Raffle Contract
/// @author Galxe Team
/// @notice This contract is used to participate in the raffle quest.
/// @custom:security-contact security@galxe.com
contract Raffle is IRaffle, Ownable2Step, Pausable, EIP712 {
    struct RaffleQuest {
        uint64 questID;
        bool active;
        IRaffle.Random random;
        mapping(uint256 => uint256) participantIds;
        uint256 participantCount;
        uint256 winnerCount;
        bytes32 merkleRoot;
    }

    /// @dev Signer address
    address public override signer;

    /// @dev Verify Ids
    BitMaps.BitMap private verifyIds;

    /// @dev Quests
    mapping(uint256 => RaffleQuest) private quests;

    /// @dev The address of the SP1 verifier contract.
    address public override verifier;

    /// @notice The verification key for the SP1 RISC-V program.
    bytes32 public override vkey;

    /// @notice The constructor for the Raffle contract.
    /// @param _initialOwner The initial owner of the contract.
    /// @param _signer The signer address.
    /// @param _verifier The SP1 verifier address.
    /// @param _vkey The verification key for the SP1 RISC-V program.
    constructor(address _initialOwner, address _signer, address _verifier, bytes32 _vkey)
        Ownable(_initialOwner)
        EIP712("Galxe Raffle", "1.0.0")
    {
        if (address(_initialOwner) == address(0)) {
            revert IRaffle.InvalidAddress();
        }
        if (address(_signer) == address(0)) {
            revert IRaffle.InvalidAddress();
        }
        if (address(_verifier) == address(0)) {
            revert IRaffle.InvalidAddress();
        }
        signer = _signer;
        verifier = _verifier;
        vkey = _vkey;
    }

    /// @notice Pauses the contract.
    function pause() public onlyOwner {
        _pause();
    }

    /// @notice Unpauses the contract.
    function unpause() public onlyOwner {
        _unpause();
    }

    /// @notice Sets the signer address.
    /// @param _signer The signer address.
    function setSigner(address _signer) public onlyOwner {
        if (_signer == address(0)) {
            revert IRaffle.InvalidAddress();
        }
        signer = _signer;
        emit IRaffle.SignerUpdated(_signer);
    }

    /// @notice Sets the verifier address.
    /// @param _verifier The new verifier address.
    function setVerifier(address _verifier) public onlyOwner {
        if (_verifier == address(0)) {
            revert IRaffle.InvalidAddress();
        }
        verifier = _verifier;
        emit IRaffle.VerifierUpdated(_verifier);
    }

    /// @notice Sets the verification key.
    /// @param _vkey The new verification key.
    function setVkey(bytes32 _vkey) public onlyOwner {
        vkey = _vkey;
        emit IRaffle.VkeyUpdated(_vkey);
    }

    /// @notice Participates in the raffle reward quest.
    /// @param _questID The quest ID.
    /// @param _user The user address.
    /// @param _verifyID The verify ID.
    /// @param _signature The signature.
    function participate(uint64 _questID, uint256 _user, uint64 _verifyID, bytes calldata _signature)
        public
        whenNotPaused
    {
        if (_questID == 0) {
            revert IRaffle.InvalidQuestID();
        }

        if (_user == 0) {
            revert IRaffle.InvalidAddress();
        }

        if (hasParticipated(_verifyID)) {
            revert IRaffle.VerifyIdAlreadyUsed(_verifyID);
        }

        bool isVerified = _verify(_hashParticipate(_questID, _user, _verifyID), _signature);
        if (!isVerified) {
            revert IRaffle.InvalidSignature();
        }

        RaffleQuest storage quest = quests[_questID];

        if (quest.questID == 0) {
            quest.questID = _questID;
            quest.active = true;
        } else {
            if (quest.active == false) {
                revert IRaffle.QuestNotActive();
            }
        }

        BitMaps.set(verifyIds, _verifyID);

        quest.participantCount++;

        uint256 participantID = quest.participantCount;
        quest.participantIds[uint256(uint160(_user))] = participantID;

        emit IRaffle.Participate(participantID, _questID, _user, _verifyID);
    }

    /// @notice Commits the randomness for the quest using Gravity L1 native randomness.
    /// @dev Reads `block.prevrandao` (unbiasable on Gravity via consensus DKG + WVUF) and gates
    ///      the call with an off-chain signature from `signer`. Anyone may relay the signed call,
    ///      so `signer` does not need to hold gas.
    ///
    ///      SECURITY (accepted trade-off): the signature authorizes the quest, not a specific
    ///      block. Because `block.prevrandao` is fixed per block, a relayer can wrap this call in
    ///      their own contract and revert the whole transaction on a losing draw, then retry in a
    ///      later block for a fresh value (test-and-abort). This is intentionally accepted here in
    ///      exchange for keeping `signer` gas-free; the operator is expected to commit promptly.
    /// @param _questID The quest ID.
    /// @param _signature The signer's authorization over the quest ID.
    function commitRandomness(uint64 _questID, bytes calldata _signature) public {
        RaffleQuest storage quest = quests[_questID];

        if (quest.questID == 0) {
            revert IRaffle.QuestNotExists();
        }

        if (!quest.active) {
            revert IRaffle.QuestNotActive();
        }

        if (quest.random.randomness != bytes32(0)) {
            revert IRaffle.QuestRandomnessAlreadyCommitted();
        }

        bool isVerified = _verify(_hashCommitRandomness(_questID), _signature);
        if (!isVerified) {
            revert IRaffle.InvalidSignature();
        }

        bytes32 randomness = bytes32(block.prevrandao);
        if (randomness == bytes32(0)) {
            // Randomness is disabled or not yet available on the chain.
            revert IRaffle.InvalidRandomness();
        }

        quest.random = IRaffle.Random({randomness: randomness, blockNumber: uint64(block.number)});
        quest.active = false;

        emit IRaffle.CommitRandomness(_questID, uint64(block.number), randomness);
    }

    /// @notice Verifies the proof and reveals the raffle result.
    /// @param _questID The quest ID.
    /// @param _publicValues The public value for proof verification encoded as bytes.
    /// @param _proofBytes The proof of the program execution the SP1 zkVM encoded as bytes.
    function reveal(uint64 _questID, bytes calldata _publicValues, bytes calldata _proofBytes) public whenNotPaused {
        RaffleQuest storage quest = quests[_questID];

        if (quest.questID == 0) {
            revert IRaffle.QuestNotExists();
        }

        if (quest.active) {
            revert IRaffle.QuestStillActive();
        }

        if (quest.merkleRoot != bytes32(0)) {
            revert IRaffle.QuestAlreadyRevealed();
        }

        (uint32 participantCount, uint32 winnerCount, bytes32 randomness, bytes32 merkleRoot) =
            abi.decode(_publicValues, (uint32, uint32, bytes32, bytes32));

        if (participantCount != quest.participantCount || randomness != quest.random.randomness) {
            revert IRaffle.IncorrectProof();
        }

        ISP1Verifier(verifier).verifyProof(vkey, _publicValues, _proofBytes);

        quest.merkleRoot = merkleRoot;
        quest.winnerCount = winnerCount;

        emit IRaffle.Reveal(_questID, participantCount, winnerCount, randomness, merkleRoot);
    }

    function getQuest(uint256 _questID)
        external
        view
        returns (
            bool _active,
            IRaffle.Random memory random,
            uint256 _participantCount,
            uint256 _winnerCount,
            bytes32 _merkleRoot
        )
    {
        RaffleQuest storage quest = quests[_questID];
        return (quest.active, quest.random, quest.participantCount, quest.winnerCount, quest.merkleRoot);
    }

    function hasParticipated(uint256 _verifyID) public view returns (bool) {
        return BitMaps.get(verifyIds, _verifyID);
    }

    function _verify(bytes32 _hash, bytes calldata _signature) private view returns (bool) {
        return ECDSA.recover(_hash, _signature) == signer;
    }

    function _hashParticipate(uint64 _questID, uint256 _user, uint64 _verifyID) private view returns (bytes32) {
        return _hashTypedDataV4(
            keccak256(
                abi.encode(
                    keccak256("Participate(uint64 questID,uint256 user,uint64 verifyID)"), _questID, _user, _verifyID
                )
            )
        );
    }

    function _hashCommitRandomness(uint64 _questID) private view returns (bytes32) {
        return _hashTypedDataV4(keccak256(abi.encode(keccak256("CommitRandomness(uint64 questID)"), _questID)));
    }
}
