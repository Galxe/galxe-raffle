package raffle

import (
	"crypto/sha256"
	"encoding/binary"
)

// RevealRaffleWinners calculates the winners, Merkle tree, and root for a given raffle configuration
func RevealRaffleWinners(numParticipants uint32, numWinners uint32, randomness [32]byte) ([]uint32, [][32]byte, [32]byte) {
	// Convert randomness to uint64 seed (using first 8 bytes)
	seed := binary.LittleEndian.Uint64(randomness[:8])

	// Get winners
	winners := selectWinners(numParticipants, numWinners, seed)

	// Convert winners to leaf hashes first
	leaves := make([][32]byte, len(winners))
	for i, winner := range winners {
		winnerBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(winnerBytes, winner)
		leaves[i] = sha256.Sum256(winnerBytes)
	}

	// Build full Merkle tree
	tree := buildMerkleTree(leaves)

	// Get root (last element in the tree)
	var root [32]byte
	if len(tree) > 0 {
		root = tree[len(tree)-1]
	}

	return winners, tree, root
}

func selectWinners(numParticipants, numWinners uint32, randomness uint64) []uint32 {
	n := int(numParticipants)
	m := int(numWinners)

	var selectCount int
	invert := false
	if m > n/2 {
		selectCount = n - m
		invert = true
	} else {
		selectCount = m
	}

	selected := make([]uint32, 0, selectCount)
	usedNumbers := make(map[uint32]bool, selectCount)
	seed := randomness

	for len(selected) < selectCount {
		newNumber := uint32(seed % uint64(n))
		if !usedNumbers[newNumber] {
			usedNumbers[newNumber] = true
			selected = append(selected, newNumber)
		}
		seed = hashSha2(seed, uint64(newNumber))
	}

	if invert {
		result := make([]uint32, 0, n-selectCount)
		for i := uint32(0); i < numParticipants; i++ {
			if !usedNumbers[i] {
				result = append(result, i)
			}
		}
		return result
	}

	return selected
}

func hashSha2(seed, value uint64) uint64 {
	h := sha256.New()
	seedBytes := make([]byte, 8)
	valueBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(seedBytes, seed)
	binary.LittleEndian.PutUint64(valueBytes, value)

	h.Write(seedBytes)
	h.Write(valueBytes)

	result := h.Sum(nil)
	return binary.LittleEndian.Uint64(result[:8])
}

func buildMerkleTree(leaves [][32]byte) [][32]byte {
	if len(leaves) == 0 {
		return [][32]byte{{}} // Return empty root for empty tree
	}

	// Initialize tree with leaves
	tree := make([][32]byte, len(leaves))
	copy(tree, leaves)

	offset := 0
	levelSize := len(leaves)

	for levelSize > 1 {
		nextLevelSize := (levelSize + 1) / 2 // Handles odd number of nodes
		for i := 0; i < levelSize; i += 2 {
			h := sha256.New()
			h.Write(tree[offset+i][:])
			// Only include right child if it exists
			if i+1 < levelSize {
				h.Write(tree[offset+i+1][:])
			} else {
				// If there's no right child, hash with self
				h.Write(tree[offset+i][:])
			}
			var hash [32]byte
			copy(hash[:], h.Sum(nil))
			tree = append(tree, hash)
		}
		offset += levelSize
		levelSize = nextLevelSize
	}

	return tree
}

// GenerateMerkleProof generates a Merkle proof for a given winner’s participantID.
// Requires:
// - participantID: The ID of the participant to prove.
// - winners: The slice of selected winners used to build the Merkle tree.
// - tree: The full Merkle tree as returned by RevealRaffleWinners.
// Returns the proof as a slice of [32]byte nodes.
//
// Optimization: We now directly know numLeaves = len(winners), so we don't need to guess.
func GenerateMerkleProof(participantID uint32, winners []uint32, tree [][32]byte) [][32]byte {
	numLeaves := len(winners)
	if numLeaves == 0 {
		return nil
	}

	// Find the leaf index for participantID
	leafIndex := -1
	for i, w := range winners {
		if w == participantID {
			leafIndex = i
			break
		}
	}
	if leafIndex == -1 {
		// Not a winner
		return nil
	}

	proof := make([][32]byte, 0)
	currentIndex := leafIndex
	offset := 0
	levelSize := numLeaves

	for levelSize > 1 {
		// Calculate sibling index
		var siblingIndex int
		if currentIndex%2 == 0 {
			// Left node, get right sibling
			siblingIndex = offset + currentIndex + 1
			if siblingIndex < offset+levelSize {
				proof = append(proof, tree[siblingIndex])
			} else {
				// No right sibling, use self (special case for odd number of nodes)
				proof = append(proof, tree[offset+currentIndex])
			}
		} else {
			// Right node, get left sibling
			siblingIndex = offset + currentIndex - 1
			proof = append(proof, tree[siblingIndex])
		}

		// Move to next level
		offset += levelSize
		currentIndex /= 2
		levelSize = (levelSize + 1) / 2
	}

	return proof
}

// VerifyMerkleProof verifies that a given leaf value (participant ID) is part of the Merkle tree defined by root.
// Inputs:
// - participantID: The leaf value for which we have a proof.
// - winners: The slice of selected winners used to build the Merkle tree.
// - proof: The slice of sibling [32]byte nodes returned by GenerateMerkleProof.
// - root: The [32]byte merkle root.
//
// Returns true if verification succeeds, false otherwise.
func VerifyMerkleProof(participantID uint32, winners []uint32, proof [][32]byte, root [32]byte) bool {
	// Find the leaf index for participantID
	leafIndex := -1
	for i, w := range winners {
		if w == participantID {
			leafIndex = i
			break
		}
	}
	if leafIndex == -1 {
		// Not a winner
		return false
	}

	// Compute leaf hash
	participantBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(participantBytes, participantID)
	hash := sha256.Sum256(participantBytes)

	currentIndex := leafIndex

	// Iteratively compute parent nodes using the proof
	for _, sibling := range proof {

		h := sha256.New()
		if currentIndex%2 == 0 {
			// Current node is a left child
			h.Write(hash[:])
			h.Write(sibling[:])
		} else {
			// Current node is a right child
			h.Write(sibling[:])
			h.Write(hash[:])
		}
		copy(hash[:], h.Sum(nil))
		currentIndex = currentIndex / 2
	}

	// After processing all proof elements, `hash` should match the root if the leaf is included
	return hash == root
}
