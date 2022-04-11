package ports

import "github.com/siraj18/zeroLvl/internal/models"

type ModelsService interface {
	AddModelToDb(model *models.Model) error
	GetModelsFromDb() ([]*models.Model, error)
	SetModelToCache(model *models.Model)
	GetModelFromCache(id string) (*models.Model, error)
}
