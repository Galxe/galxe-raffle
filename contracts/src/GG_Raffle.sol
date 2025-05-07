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
import {IDrandOracle, GGIRaffle} from "./GG_IRaffle.sol";

/// @title Galxe Raffle Contract
/// @author Galxe Team
/// @notice This contract is used to participate in the gg raffle quest.
/// @custom:security-contact security@galxe.com
contract GGRaffle is GGIRaffle, Ownable2Step, Pausable, EIP712 {
    struct RaffleConfig {
        uint256 winnerCount;
        bytes32 merkleRoot;
    }

    struct GGRaffleQuest {
        uint256 questID;
        bool active;
        IDrandOracle.Random random;
        uint256 participantCount;

        // raffleType --> raffleConfig
        mapping(uint256 => RaffleConfig) raffleConfigs;
    }

    /// @dev Signer address
    address public override signer;

    /// @dev Used Verify Ids
    BitMaps.BitMap private usedVerifyIds;

    /// @dev Quests
    mapping(uint256 => GGRaffleQuest) private quests;

    /// @dev The address of the SP1 verifier contract.
    address public override verifier;

    /// @dev The address of the DrandOracle contract.
    address public override drandOracle;

    /// @notice The verification key for the SP1 RISC-V program.
    bytes32 public override vkey;

    /// @notice The constructor for the Raffle contract.
    /// @param _initialOwner The initial owner of the contract.
    /// @param _signer The signer address.
    /// @param _verifier The SP1 verifier address.
    /// @param _vkey The verification key for the SP1 RISC-V program.
    constructor(address _initialOwner, address _signer, address _verifier, bytes32 _vkey, address _drandOracle)
    Ownable(_initialOwner)
    EIP712("Galxe GG Raffle", "1.0.0")
    {
        if (address(_initialOwner) == address(0)) {
            revert GGIRaffle.InvalidAddress();
        }
        if (address(_signer) == address(0)) {
            revert GGIRaffle.InvalidAddress();
        }
        if (address(_verifier) == address(0)) {
            revert GGIRaffle.InvalidAddress();
        }
        signer = _signer;
        verifier = _verifier;
        vkey = _vkey;
        drandOracle = _drandOracle;
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
            revert GGIRaffle.InvalidAddress();
        }
        signer = _signer;
        emit GGIRaffle.SignerUpdated(_signer);
    }

    /// @notice Sets the verifier address.
    /// @param _verifier The new verifier address.
    function setVerifier(address _verifier) public onlyOwner {
        if (_verifier == address(0)) {
            revert GGIRaffle.InvalidAddress();
        }
        verifier = _verifier;
        emit GGIRaffle.VerifierUpdated(_verifier);
    }

    /// @notice Sets the verification key.
    /// @param _vkey The new verification key.
    function setVkey(bytes32 _vkey) public onlyOwner {
        vkey = _vkey;
        emit GGIRaffle.VkeyUpdated(_vkey);
    }

    /// @notice Sets the DrandOracle address.
    /// @param _drandOracle The new DrandOracle address.
    function setDrandOracle(address _drandOracle) public onlyOwner {
        drandOracle = _drandOracle;
        emit GGIRaffle.DrandOracleUpdated(_drandOracle);
    }

    /// @notice Participates in the raffle reward quest.
    /// @param _questID The quest ID.
    /// @param _user The user address.
    /// @param _verifyID The verify ID.
    /// @param _signature The signature.
    function participate(uint256 _questID, uint256 _user, uint256 _verifyID, bytes calldata _signature)
    public
    whenNotPaused
    {
        if (_questID == 0) {
            revert GGIRaffle.InvalidQuestID();
        }

        if (_user == 0) {
            revert GGIRaffle.InvalidAddress();
        }

        if (hasParticipated(_verifyID)) {
            revert GGIRaffle.VerifyIdAlreadyUsed(_verifyID);
        }

        bool isVerified = _verify(_hashParticipate(_questID, _user, _verifyID), _signature);
        if (!isVerified) {
            revert GGIRaffle.InvalidSignature();
        }

        GGRaffleQuest storage quest = quests[_questID];

        if (quest.questID == 0) {
            quest.questID = _questID;
            quest.active = true;
        } else {
            if (quest.active == false) {
                revert GGIRaffle.QuestNotActive();
            }
        }

        BitMaps.set(usedVerifyIds, _verifyID);

        quest.participantCount++;

        uint256 participantID = quest.participantCount;
        emit GGIRaffle.Participate(participantID, _questID, _user, _verifyID);
    }


    /// @notice ParticipatesBatch in the raffle reward quest.
    /// @param _questID The quest ID.
    /// @param _user The user address.
    /// @param _verifyIDs The verify ID array.
    /// @param _signature The signature.
    function participateBatch(uint256 _questID, uint256 _user, uint256[] calldata _verifyIDs, bytes calldata _signature)
    public
    whenNotPaused
    {
        if (_questID == 0) {
            revert GGIRaffle.InvalidQuestID();
        }

        if (_user == 0) {
            revert GGIRaffle.InvalidAddress();
        }

        if (_verifyIDs.length == 0) {
            revert GGIRaffle.InvalidQuestIDArray();
        }

        for (uint256 i = 0; i < _verifyIDs.length; i++) {
            if (hasParticipated(_verifyIDs[i])) {
                revert GGIRaffle.VerifyIdAlreadyUsed(_verifyIDs[i]);
            }
            BitMaps.set(usedVerifyIds, _verifyIDs[i]);
        }

        bool isVerified = _verify(_hashParticipateBatch(_questID, _user, _verifyIDs), _signature);
        if (!isVerified) {
            revert GGIRaffle.InvalidSignature();
        }

        GGRaffleQuest storage quest = quests[_questID];

        if (quest.questID == 0) {
            quest.questID = _questID;
            quest.active = true;
        } else {
            if (quest.active == false) {
                revert GGIRaffle.QuestNotActive();
            }
        }

        // less update for quest.participantCount to reduce gasfee
        uint256 _participantCount = quest.participantCount;
        quest.participantCount += _verifyIDs.length;

        for (uint256 i = 0; i < _verifyIDs.length; i++) {
            uint256 participantID = _participantCount + i + 1;

            emit GGIRaffle.Participate(participantID, _questID, _user, _verifyIDs[i]);
        }
    }


    /// @notice Commits the randomness for the quest.
    /// @param _questID The quest ID.
    /// @param _signature The signature.
    function commitRandomness(uint256 _questID, uint256 _timestamp, bytes calldata _signature) public {
        GGRaffleQuest storage quest = quests[_questID];

        if (quest.questID == 0) {
            revert GGIRaffle.QuestNotExists();
        }

        if (!quest.active) {
            revert GGIRaffle.QuestNotActive();
        }

        if (quest.random.randomness != bytes32(0)) {
            revert GGIRaffle.QuestRandomnessAlreadyCommitted();
        }

        bool isVerified = _verify(_hashCommitRandomness(_questID, _timestamp), _signature);
        if (!isVerified) {
            revert GGIRaffle.InvalidSignature();
        }

        IDrandOracle.Random memory random = IDrandOracle(drandOracle).getRandomnessFromTimestamp(_timestamp);

        if (random.round == 0 || random.randomness == bytes32(0) || random.signature.length == 0) {
            revert GGIRaffle.InvalidRandomness();
        }

        quest.random = random;
        quest.active = false;

        emit GGIRaffle.CommitRandomness(_questID, random.round, random.randomness);
    }

    /// @notice Verifies the proof and reveals the raffle result.
    /// @param _questID The quest ID.
    /// @param _raffleType The raffle type (1 present grand, 2 present small).
    /// @param _publicValues The public value for proof verification encoded as bytes.
    /// @param _proofBytes The proof of the program execution the SP1 zkVM encoded as bytes.
    function reveal(uint256 _questID, uint256 _raffleType, bytes calldata _publicValues, bytes calldata _proofBytes) public whenNotPaused {
        GGRaffleQuest storage quest = quests[_questID];

        if (quest.questID == 0) {
            revert GGIRaffle.QuestNotExists();
        }

        if (quest.active) {
            revert GGIRaffle.QuestStillActive();
        }

        if (quest.raffleConfigs[_raffleType].merkleRoot != bytes32(0)) {
            revert GGIRaffle.QuestAlreadyRevealed();
        }

        (uint32 participantCount, uint32 winnerCount, bytes32 randomness, bytes32 merkleRoot) =
                            abi.decode(_publicValues, (uint32, uint32, bytes32, bytes32));

        if (participantCount != quest.participantCount || randomness != quest.random.randomness) {
            revert GGIRaffle.IncorrectProof();
        }

        ISP1Verifier(verifier).verifyProof(vkey, _publicValues, _proofBytes);

        quest.raffleConfigs[_raffleType].merkleRoot = merkleRoot;
        quest.raffleConfigs[_raffleType].winnerCount = winnerCount;

        emit GGIRaffle.Reveal(_questID, _raffleType, participantCount, winnerCount, randomness, merkleRoot);
    }

    function getQuest(uint256 _questID, uint256 _raffleType)
    external
    view
    returns (
        bool _active,
        IDrandOracle.Random memory random,
        uint256 _participantCount,
        uint256 _winnerCount,
        bytes32 _merkleRoot
    )
    {
        GGRaffleQuest storage quest = quests[_questID];
        return (quest.active, quest.random, quest.participantCount, quest.raffleConfigs[_raffleType].winnerCount, quest.raffleConfigs[_raffleType].merkleRoot);
    }

    function hasParticipated(uint256 _verifyID) public view returns (bool) {
        return BitMaps.get(usedVerifyIds, _verifyID);
    }

    function _verify(bytes32 _hash, bytes calldata _signature) private view returns (bool) {
        return ECDSA.recover(_hash, _signature) == signer;
    }

    function _hashParticipate(uint256 _questID, uint256 _user, uint256 _verifyID) private view returns (bytes32) {
        return _hashTypedDataV4(
            keccak256(
                abi.encode(
                    keccak256(
                        "Participate(uint256 questID,uint256 user,uint256 verifyID)"
                    ),
                    _questID,
                    _user,
                    _verifyID
                )
            )
        );
    }

    function _hashParticipateBatch(uint256 _questID, uint256 _user, uint256[] calldata _verifyIDs) private view returns (bytes32) {
        return _hashTypedDataV4(
            keccak256(
                abi.encode(
                    keccak256(
                        "ParticipateBatch(uint256 questID,uint256 user,uint256[] verifyIDs)"
                    ),
                    _questID,
                    _user,
                    keccak256(abi.encodePacked(_verifyIDs))
                )
            )
        );
    }

    function _hashCommitRandomness(uint256 _questID, uint256 _timestamp) private view returns (bytes32) {
        return _hashTypedDataV4(
            keccak256(
                abi.encode(
                    keccak256("CommitRandomness(uint256 questID,uint256 timestamp)"
                    ),
                    _questID,
                    _timestamp
                )
            )
        );
    }
}
