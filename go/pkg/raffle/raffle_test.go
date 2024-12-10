package raffle

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"slices"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Fixture paths for different raffle configurations
const raffleFixture_1000_choose_10 = "../../../zk/fixtures/groth16-1000-10-82e8b6bbf24681c9d3c928f988aa6eef88f41f164e5df290e3dca3b8f6ce3f07-fixture.json"
const raffleFixture_1000_choose_900 = "../../../zk/fixtures/groth16-1000-900-82e8b6bbf24681c9d3c928f988aa6eef88f41f164e5df290e3dca3b8f6ce3f07-fixture.json"

// Suite defines the test suite structure containing fixture data and parsed randomness and expected root
type Suite struct {
	suite.Suite
	fixture struct {
		NumParticipants uint32   `json:"numParticipants"`
		NumWinners      uint32   `json:"numWinners"`
		Randomness      string   `json:"randomness"`
		MerkleRoot      string   `json:"merkleRoot"`
		Winners         []uint32 `json:"winners"`
	}
	randomness [32]byte
	merkleRoot [32]byte
}

// TestRaffleSuite runs all the tests in the Suite
func TestRaffleSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

// loadFixture reads and parses a JSON fixture file, extracting:
// - Raffle parameters (number of participants and winners)
// - Random seed used for winner selection
// - Expected Merkle root for verification
func (s *Suite) loadFixture(path string) {
	// Load fixture data
	fixtureData, err := os.ReadFile(path)
	s.Require().NoError(err, "failed to read fixture file")

	err = json.Unmarshal(fixtureData, &s.fixture)
	s.Require().NoError(err, "failed to unmarshal fixture data")

	// Load randomness from fixture
	randomnessStr := s.fixture.Randomness
	if len(randomnessStr) >= 2 && randomnessStr[:2] == "0x" {
		randomnessStr = randomnessStr[2:]
	}
	randomnessBytes, err := hex.DecodeString(randomnessStr)
	s.Require().NoError(err, "failed to decode randomness hex")
	copy(s.randomness[:], randomnessBytes)

	// Load expected root from fixture
	merkleRootStr := s.fixture.MerkleRoot
	if len(merkleRootStr) >= 2 && merkleRootStr[:2] == "0x" {
		merkleRootStr = merkleRootStr[2:]
	}
	merkleRootBytes, err := hex.DecodeString(merkleRootStr)
	s.Require().NoError(err, "failed to decode merkle root hex")
	copy(s.merkleRoot[:], merkleRootBytes)
}

// TestRevealRaffleWinners_1000_choose_10 verifies winner selection and Merkle tree
// generation for a raffle with 1000 participants and 10 winners
func (s *Suite) TestRevealRaffleWinners_1000_choose_10() {
	s.loadFixture(raffleFixture_1000_choose_10)
	winners, tree, root := RevealRaffleWinners(s.fixture.NumParticipants, s.fixture.NumWinners, s.randomness)
	s.Equal(int(s.fixture.NumWinners), len(winners), "incorrect number of winners")
	s.Equal(s.merkleRoot, root, "Merkle root mismatch")
	s.Equal(s.merkleRoot, tree[len(tree)-1], "Merkle tree root mismatch")
}

// TestRevealRaffleWinners_1000_choose_900 verifies winner selection and Merkle tree
// generation for a raffle with 1000 participants and 900 winners
func (s *Suite) TestRevealRaffleWinners_1000_choose_900() {
	s.loadFixture(raffleFixture_1000_choose_900)
	winners, tree, root := RevealRaffleWinners(s.fixture.NumParticipants, s.fixture.NumWinners, s.randomness)
	s.Equal(int(s.fixture.NumWinners), len(winners), "incorrect number of winners")
	s.Equal(s.merkleRoot, root, "Merkle root mismatch")
	s.Equal(s.merkleRoot, tree[len(tree)-1], "Merkle tree root mismatch")
}

func (s *Suite) TestGenerateMerkleProof_1000_choose_10() {
	s.loadFixture(raffleFixture_1000_choose_10)
	proofSize := int(math.Ceil(math.Log2(float64(s.fixture.NumWinners))))

	winners, tree, root := RevealRaffleWinners(s.fixture.NumParticipants, s.fixture.NumWinners, s.randomness)

	// verify all winner proofs
	for _, participantID := range winners {
		merkleProof := GenerateMerkleProof(participantID, winners, tree)
		hexProof := make([]string, len(merkleProof))
		for i, proof := range merkleProof {
			hexProof[i] = fmt.Sprintf("0x%x", proof)
		}
		s.Equal(proofSize, len(merkleProof), "Merkle proof size mismatch")

		proofHex := make([]string, len(merkleProof))
		for i, proof := range merkleProof {
			proofHex[i] = fmt.Sprintf("0x%x", proof)
		}

		verified := VerifyMerkleProof(participantID, winners, merkleProof, root)
		s.True(verified, "Merkle proof verification failed")
	}

	// verify non-winner proofs
	for participantID := uint32(0); participantID < s.fixture.NumParticipants; participantID++ {
		if !slices.Contains(winners, participantID) {
			merkleProof := GenerateMerkleProof(participantID, winners, tree)
			s.Equal(0, len(merkleProof), "Merkle proof size should be 0 for non-winner")
			verified := VerifyMerkleProof(participantID, winners, merkleProof, root)
			s.False(verified, "Merkle proof verification failed for non-winner")
		}
	}
}

func (s *Suite) TestGenerateMerkleProof_1000_choose_900() {
	s.loadFixture(raffleFixture_1000_choose_900)
	proofSize := int(math.Ceil(math.Log2(float64(s.fixture.NumWinners))))

	winners, tree, root := RevealRaffleWinners(s.fixture.NumParticipants, s.fixture.NumWinners, s.randomness)

	// verify all winner proofs
	for _, participantID := range winners {
		merkleProof := GenerateMerkleProof(participantID, winners, tree)
		hexProof := make([]string, len(merkleProof))
		for i, proof := range merkleProof {
			hexProof[i] = fmt.Sprintf("0x%x", proof)
		}
		s.Equal(proofSize, len(merkleProof), "Merkle proof size mismatch")

		proofHex := make([]string, len(merkleProof))
		for i, proof := range merkleProof {
			proofHex[i] = fmt.Sprintf("0x%x", proof)
		}

		verified := VerifyMerkleProof(participantID, winners, merkleProof, root)
		s.True(verified, "Merkle proof verification failed")
	}

	// verify non-winner proofs
	for participantID := uint32(0); participantID < s.fixture.NumParticipants; participantID++ {
		if !slices.Contains(winners, participantID) {
			merkleProof := GenerateMerkleProof(participantID, winners, tree)
			s.Equal(0, len(merkleProof), "Merkle proof size should be 0 for non-winner")
			verified := VerifyMerkleProof(participantID, winners, merkleProof, root)
			s.False(verified, "Merkle proof verification failed for non-winner")
		}
	}
}
