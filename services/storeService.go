package services

import (
	"context"

	"github.com/enesyalinkaya/go-kv-store/models"
)

type StoreServiceConfig struct {
	StoreModel models.StoreModel
}

type storeService struct {
	StoreModel models.StoreModel
}

// StoreService interface represents a store service functionality
type StoreService interface {
	Set(ctx context.Context, key string, value string) string
	Get(ctx context.Context, key string) string
	Flush(ctx context.Context)
}

// it is Create new StoreService instance
func NewStoreService(c *StoreServiceConfig) StoreService {
	return &storeService{
		StoreModel: c.StoreModel,
	}
}

// Implementing Get function. It gets value from StoreModel and return
func (s *storeService) Get(ctx context.Context, key string) string {
	return s.StoreModel.Get(key)
}

// Implementing Set function. It sets value
func (s *storeService) Set(ctx context.Context, key string, value string) string {
	return s.StoreModel.Set(key, value)
}

// Implementing Flush function. It flushes in-memory key-value store
func (s *storeService) Flush(ctx context.Context) {
	s.StoreModel.Flush()
}
