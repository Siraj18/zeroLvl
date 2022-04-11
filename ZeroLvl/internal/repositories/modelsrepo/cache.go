package modelsrepo

import (
	"errors"

	"github.com/siraj18/zeroLvl/internal/models"
	"github.com/siraj18/zeroLvl/pkg/inmemory"
)

var ErrorNotFound = errors.New("item not found")

type CacheRepository struct {
	cache *inmemory.Cache
}

func NewCacheRepository(cache *inmemory.Cache) *CacheRepository {
	return &CacheRepository{
		cache: cache,
	}
}

func (rep *CacheRepository) Get(key string) (*models.Model, error) {

	item := rep.cache.Get(key)

	if item == nil {
		return nil, ErrorNotFound
	}

	model, ok := item.(*models.Model)
	if !ok {
		return nil, errors.New("error when get model")
	}

	return model, nil
}

func (rep *CacheRepository) Set(key string, value *models.Model) {
	rep.cache.Set(value.OrderUID, value)
}
