package internal

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func localDb(connString string) *sql.DB {
	dbPath := connString[7:]

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		os.Create(dbPath)
		fmt.Println("db file created")
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(fmt.Sprintf("failed to open db \"%s\", error: %s", connString, err))
	}

	return db
}
