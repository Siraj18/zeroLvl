package ports

import "github.com/siraj18/zeroLvl/internal/models"

type PostgreRepository interface {
	GetModels() ([]*models.Model, error)
	AddModel(model *models.Model) error
}

type CacheRepository interface {
	Get(key string) (*models.Model, error)
	Set(key string, value *models.Model)
}
