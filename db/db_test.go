package db_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"example.com/shortlink-kata/db"
)

func TestInsert(t *testing.T) {
	testSlug := "star-wars"
	testLink := "tauntaun"
	require.False(t, db.Exists(testSlug), "Expected test data to be absent before test is run")

	err := db.InsertNewLink(testSlug, testLink)
	require.NoError(t, err, "Expected no error in fresh insert")

	require.True(t, db.Exists(testSlug), "Expected test data to be present after the test is run")

	err = cleanup(testSlug)
	require.NoError(t, err, "Expected no error in cleanup")

}

func cleanup(slug string) error {
	db := db.GetDB()

	_, err := db.Exec("DELETE FROM link_visits WHERE slug = $1", slug)
	if err != nil {
		return fmt.Errorf("could not delete row from link_visits: %v", err)
	}

	_, err = db.Exec("DELETE FROM links WHERE slug = $1", slug)
	if err != nil {
		return fmt.Errorf("could not delete row from links: %v", err)
	}
	return nil
}
