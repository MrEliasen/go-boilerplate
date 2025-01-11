package internal

import (
	"github.com/placeholder/boiler/internal/database/internal/repositories"
	dbShared "github.com/placeholder/boiler/internal/database/shared"
)

var (
	migrationRepo dbShared.MigrationsRepositoryInterface
	userRepo      dbShared.UserRepositoryInterface
)

func UserRepo() dbShared.UserRepositoryInterface {
	if userRepo == nil {
		userRepo = repositories.NewUserRepository(conn)
	}

	return userRepo
}

func MigrationRepo() dbShared.MigrationsRepositoryInterface {
	if migrationRepo == nil {
		migrationRepo = repositories.NewMigrationRepository(conn)
	}

	return migrationRepo
}
