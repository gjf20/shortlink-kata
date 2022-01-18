package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetDB() *sql.DB { //todo connection pooling

	connStr := "user=postgres dbname=test_link sslmode=disable" //would interpolate the password with environment variables - no password set for test database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Printf("Failed to keep connection alive, %v \n", err)
	}

	return db
}

func Exists(slug string) bool {
	db := GetDB()
	row := db.QueryRow("SELECT link FROM links WHERE slug = $1;", slug)

	var link string
	err := row.Scan(&link)

	if err == sql.ErrNoRows {
		return false
	}
	if err != nil {
		fmt.Printf("encountered error querying database: %v\n", err)
	}

	return true
}

func InsertNewLink(slug, link string) error {
	db := GetDB()
	result, err := db.Exec("INSERT INTO links (slug, link) VALUES ($1, $2)", slug, link)
	if err != nil {
		return fmt.Errorf("could not insert row: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get affected rows of insert: %v", err)
	}
	if rowsAffected != 1 {
		return fmt.Errorf("unexpectedly inserted %v rows", rowsAffected)
	}
	return nil
}
