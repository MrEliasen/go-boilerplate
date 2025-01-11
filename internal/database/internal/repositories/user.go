package repositories

import (
	"database/sql"

	"github.com/placeholder/boiler/internal/database/models"
)

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		conn: db,
	}
}

type UserRepository struct {
	conn *sql.DB
}

func (r *UserRepository) GetByID(int64) (*models.User, error) {
	return nil, nil
}

func (r *UserRepository) Update(*models.User) error {
	return nil
}

func (r *UserRepository) Insert(*models.User) (int64, error) {
	return 0, nil
}
