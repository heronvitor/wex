package repository

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func createDB(t *testing.T) (db *sql.DB) {
	dbURL := "postgresql://wex:pass@localhost/wex_test?sslmode=disable" //os.Getenv("DB_URL")
	if !regexp.MustCompile(".*test").MatchString(dbURL) {
		t.Fatal("test should be executed using test database")
	}

	m, err := migrate.New("file://../../../db/migrations", dbURL)
	if err != nil {
		t.Fatal(err.Error())
	}

	if err = m.Down(); err != nil {
		fmt.Println(err.Error())
	}

	if err = m.Up(); err != nil {
		t.Fatal(err.Error())
	}
	m.Close()

	if db, err = sql.Open("postgres", dbURL); err != nil {
		t.Fatal(err.Error())
	}
	return
}
