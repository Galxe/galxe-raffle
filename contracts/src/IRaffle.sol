// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.24;

/// @title Minimal interface for drand oracle
interface IDrandOracle {
    /// @notice Structure containing randomness data and associated signatures
    /// @param round The drand round number
    /// @param randomness The drand round randomness value
    /// @param signature The drand round signature
    struct Random {
        uint64 round;
        uint64 timestamp;
        bytes32 randomness;
        bytes signature;
    }

    /// @notice Retrieves the complete randomness data for a specific timestamp
    /// @param _timestamp The timestamp to query
    /// @return The Random struct containing the timestamp's data
    function getRandomnessFromTimestamp(uint64 _timestamp) external view returns (Random memory);
}

/// @title Interface for Galxe Raffle Contract
/// @notice Interface for participating in raffle quests
interface IRaffle {
    // Custom Errors
    error InvalidAddress();
    error InvalidQuestID();
    error InvalidRandomness();
    error InvalidSignature();
    error VerifyIdAlreadyUsed(uint256 verifyId);
    error QuestNotExists();
    error QuestNotActive();
    error QuestStillActive();
    error QuestRandomnessAlreadyCommitted();
    error QuestAlreadyRevealed();
    error IncorrectProof();

    // Events
    event SignerUpdated(address signer);
    event VerifierUpdated(address verifier);
    event VkeyUpdated(bytes32 vkey);
    event DrandOracleUpdated(address drandOracle);
    event Participate(uint64 questID, uint256 user, uint64 verifyID);
    event CommitRandomness(uint64 questID, uint64 roundID, bytes32 randomness);
    event Reveal(uint64 questID, uint32 participantCount, uint32 winnerCount, bytes32 randomness, bytes32 merkleRoot);

    // Functions
    function pause() external;
    function unpause() external;
    function setSigner(address _signer) external;

    function participate(uint64 _questID, uint256 _user, uint64 _verifyID, bytes calldata _signature) external;

    function commitRandomness(uint64 _questID, uint64 _timestamp, bytes calldata _signature) external;

    function reveal(uint64 _questID, bytes calldata _publicValues, bytes calldata _proofBytes) external;

    function getQuest(uint256 _questID)
        external
        view
        returns (
            bool _active,
            IDrandOracle.Random memory random,
            uint256 _participantCount,
            uint256 _winnerCount,
            bytes32 _merkleRoot
        );

    function hasParticipated(uint256 _verifyID) external view returns (bool);

    // State Variables
    function signer() external view returns (address);
    function verifier() external view returns (address);
    function vkey() external view returns (bytes32);
    function drandOracle() external view returns (address);
}
