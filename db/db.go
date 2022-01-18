package db

import (
	"database/sql"
	"fmt"
	"time"

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
	_, err := GetLink(slug)

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

func GetLink(slug string) (string, error) {
	var link string
	db := GetDB()
	row := db.QueryRow("SELECT link FROM links WHERE slug = $1;", slug)

	err := row.Scan(&link)
	return link, err
}

func SlugVisited(slug string) error {
	//increment the total visits
	var visits int
	db := GetDB()
	row := db.QueryRow("SELECT visits FROM link_data WHERE slug = $1;", slug)

	err := row.Scan(&visits)
	if err != nil {
		return err
	}
	visits += 1
	db.Exec("UPDATE link_data SET visits = $1 WHERE slug = $2", visits, slug)

	//increment the day that the slug was visited on
	var dayVisits int
	row = db.QueryRow("SELECT visits FROM link_visits WHERE slug = $1 AND day = CURRENT_DATE;", slug)

	err = row.Scan(&dayVisits)
	if err == sql.ErrNoRows {
		db.Exec("INSERT INTO link_visits (slug) VALUES ($1)", slug) //default visits is 1
	} else if err == nil {
		dayVisits += 1
		db.Exec("UPDATE link_visits SET visits = $1 WHERE slug = $2 AND day = CURRENT_DATE", dayVisits, slug)
	}

	return err
}

func GetLinkData(slug string) (int, time.Time, error) { //todo refactor ugly signature
	db := GetDB()
	row := db.QueryRow("SELECT visits, created_at FROM link_data WHERE slug = $1;", slug)

	var visits int
	var created time.Time
	err := row.Scan(&visits, &created)
	if err != nil {
		return 0, time.Time{}, err
	}
	return visits, created, nil
}
