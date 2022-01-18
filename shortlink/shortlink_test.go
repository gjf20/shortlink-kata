package shortlink

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShortlink(t *testing.T) {
	expectedHash := "test_hash_slug"
	testLink := "test_link"
	oldFunc := getHash
	getHash = func(string) (string, error) {
		return expectedHash, nil
	}
	oldInsertFunc := insert
	wasCalled := false
	insert = func(string, string) error {
		wasCalled = true
		return nil
	}
	defer func() {
		getHash = oldFunc
		insert = oldInsertFunc
	}()

	resp, err := createNewShortlink(&NewShortRequest{Link: testLink})
	require.NoError(t, err)
	require.True(t, wasCalled)
	require.Equal(t, expectedHash, resp.Slug)
	require.Equal(t, testLink, resp.Link)
}

func TestCustomShortlink(t *testing.T) {
	customSlug := "test_hash_slug"
	testLink := "test_link"
	oldFunc := getHash
	hashWasCalled := false
	getHash = func(string) (string, error) {
		hashWasCalled = true
		return "fake slug", nil
	}
	oldInsertFunc := insert
	insertWasCalled := false
	insert = func(string, string) error {
		insertWasCalled = true
		return nil
	}
	defer func() {
		getHash = oldFunc
		insert = oldInsertFunc
	}()

	resp, err := createNewShortlink(&NewShortRequest{Link: testLink, CustomSlug: customSlug})
	require.NoError(t, err)
	require.True(t, insertWasCalled)
	require.False(t, hashWasCalled)
	require.Equal(t, customSlug, resp.Slug)
	require.Equal(t, testLink, resp.Link)
}
