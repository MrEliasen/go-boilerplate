package database

import (
	"database/sql"
	"os"

	"github.com/placeholder/boiler/internal/database/internal"
)

func Connection() *sql.DB {
	connString := os.Getenv("SQLITE_DB_URL")
	return internal.Connect(connString)
}
