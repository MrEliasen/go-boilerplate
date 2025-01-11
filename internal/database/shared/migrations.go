package shared

import "github.com/placeholder/boiler/internal/database/models"

type MigrationsRepositoryInterface interface {
	GetByName(string) (*models.Migration, error)
	ListAll() ([]*models.Migration, error)
	Insert(string) (int64, error)
	DeleteByName(string) (int64, error)
}
