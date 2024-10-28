// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.24;

/// @title Interface for Galxe Raffle Contract
/// @notice Interface for participating in raffle quests
interface IRaffle {
    // Custom Errors
    error InvalidAddress();
    error InvalidQuestID();
    error InvalidBeacon();
    error InvalidSignature();
    error VerifyIdAlreadyUsed(uint256 verifyId);
    error QuestNotExists();
    error QuestNotActive();
    error QuestStillActive();
    error QuestRandomnessAlreadyCommitted();
    error QuestAlreadyRevealed();
    error IncorrectProof();

    // Events
    event Participate(uint64 questID, uint256 user, uint64 verifyID);
    event CommitRandomness(uint64 questID, uint64 roundID, bytes32 randomness);
    event Reveal(uint64 questID, uint32 participantCount, uint32 winnerCount, bytes32 randomness, bytes32 merkleRoot);

    // Structs
    struct DrandBeacon {
        uint64 round;
        bytes32 randomness;
        bytes signature;
        bytes previousSignature;
    }

    // Functions
    function pause() external;
    function unpause() external;
    function setSigner(address _signer) external;

    function participate(uint64 _questID, uint256 _user, uint64 _verifyID, bytes calldata _signature) external;

    function commitRandomness(uint64 _questID, DrandBeacon calldata _beacon, bytes calldata _signature) external;

    function reveal(uint64 _questID, bytes calldata _publicValues, bytes calldata _proofBytes) external;

    function getQuest(uint256 _questID)
        external
        view
        returns (bool _active, DrandBeacon memory _beacon, bytes32 _merkleRoot);

    function hasParticipated(uint256 _verifyID) external view returns (bool);

    // State Variables
    function signer() external view returns (address);
    function verifier() external view returns (address);
    function vkey() external view returns (bytes32);
}
