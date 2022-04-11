package modelsrv

import (
	"github.com/siraj18/zeroLvl/internal/models"
	"github.com/siraj18/zeroLvl/internal/ports"
)

type service struct {
	postgre ports.PostgreRepository
	cache   ports.CacheRepository
}

func NewModelsService(postgre ports.PostgreRepository, cache ports.CacheRepository) *service {
	return &service{
		postgre: postgre,
		cache:   cache,
	}
}

func (s *service) InitCacheFromDb() error {
	models, err := s.GetModelsFromDb()

	if err != nil {
		return err
	}

	for _, model := range models {
		s.SetModelToCache(model)
	}

	return nil
}

func (s *service) AddModelToDb(model *models.Model) error {
	err := s.postgre.AddModel(model)

	return err
}

func (s *service) GetModelsFromDb() ([]*models.Model, error) {
	models, err := s.postgre.GetModels()

	if err != nil {
		return nil, err
	}

	return models, nil
}

func (s *service) SetModelToCache(model *models.Model) {
	s.cache.Set(model.OrderUID, model)
}

func (s *service) GetModelFromCache(id string) (*models.Model, error) {
	model, err := s.cache.Get(id)

	if err != nil {
		return nil, err
	}

	return model, nil
}
