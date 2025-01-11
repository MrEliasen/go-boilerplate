package internal

import (
	"database/sql"
	"embed"
	"strings"

	"github.com/placeholder/boiler/pkg/logger"
)

//go:embed migrations/*sql
var migrationFiles embed.FS

func migrate(conn *sql.DB) {
	logger.Logger().Info("Running migrations")

	files, err := migrationFiles.ReadDir("migrations")
	if err != nil {
		logger.Logger().Panic(err)
	}

	mgRepo := MigrationRepo()

	for _, m := range files {
		if m.IsDir() {
			continue
		}

		name := strings.TrimSuffix(m.Name(), ".sql")

		found, err := mgRepo.GetByName(name)
		if err != nil {
			if !strings.Contains(err.Error(), "no such table: migrations") {
				logger.Logger().Panic(err)
			}
		}

		if found != nil {
			continue
		}

		content, err := migrationFiles.ReadFile("migrations/" + m.Name())
		if err != nil {
			logger.Logger().Panic(err)
		}

		res, err := conn.Exec(string(content))
		if err != nil {
			logger.Logger().Panic(err)
		}

		_, err = res.RowsAffected()
		if err != nil {
			logger.Logger().Panic(err)
		}

		logger.Logger().Debugf("%s\n", name)
		mgRepo.Insert(name)
	}
}
