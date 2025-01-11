package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func libsqlDb(connString string) *sql.DB {
	db, err := sql.Open("libsql", connString)
	if err != nil {
		panic(fmt.Sprintf("failed to open db \"%s\", error: %s", connString, err))
	}

	return db
}
