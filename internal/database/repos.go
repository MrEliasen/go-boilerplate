package database

import (
	"github.com/placeholder/boiler/internal/database/internal"
	dbShared "github.com/placeholder/boiler/internal/database/shared"
)

func MigrationsRepository() dbShared.MigrationsRepositoryInterface {
	return internal.MigrationRepo()
}

func UserRepository() dbShared.UserRepositoryInterface {
	return internal.UserRepo()
}
