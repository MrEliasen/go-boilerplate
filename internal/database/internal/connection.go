package internal

import (
	"database/sql"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var conn *sql.DB

func Connect(connString string) *sql.DB {
	if conn != nil {
		return conn
	}

	if strings.Contains(connString, "file://") {
		conn = localDb(connString)
	} else {
		conn = libsqlDb(connString)
	}

	migrate(conn)
	return conn
}
