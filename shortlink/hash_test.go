package shortlink

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSameShortlink(t *testing.T) {
	oldFunc := isAlreadyCreated
	isAlreadyCreated = func(string) bool {
		return false
	}
	defer func() {
		isAlreadyCreated = oldFunc
	}()
	testSlug := "fake_slug_to_test"

	hash, err := generateHash(testSlug)
	require.NoError(t, err)
	hash2, err := generateHash(testSlug)
	require.NoError(t, err)

	require.Equal(t, hash, hash2, "Expected Hashes to be the same")
}

func TestShortlinkCollision(t *testing.T) {
	oldFunc := isAlreadyCreated
	timesCalled := 0
	isAlreadyCreated = func(string) bool { //use as a closure
		timesCalled++
		return timesCalled == 2 //returns true on the second call -> allows the first hash to "be created" and forces the second hash to add a character
	}
	defer func() {
		isAlreadyCreated = oldFunc
	}()
	testSlug := "fake_slug_to_test"

	hash, err := generateHash(testSlug)
	require.NoError(t, err)

	hash2, err := generateHash(testSlug)
	require.NoError(t, err)

	require.Equal(t, len(hash)+2, len(hash2), fmt.Sprintf("Expected second hash to have 2 extra characters because of hex encoding\n %v\n %v", hash, hash2))
	require.Contains(t, hash2, hash, "Expected second hash contian the first hash")
}
