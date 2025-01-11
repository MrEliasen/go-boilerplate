package shared

import "github.com/placeholder/boiler/internal/database/models"

type UserRepositoryInterface interface {
	GetByID(int64) (*models.User, error)
	Update(*models.User) error
	Insert(*models.User) (int64, error)
}
