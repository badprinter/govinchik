package storage

import "github.com/badprinter/govinchik/internal/config"

type StorageAPI struct {
	cfg *config.StorageConfig
}

func NewStorageAPI(c *config.Config) *StorageAPI {
	return &StorageAPI{
		cfg: c.Storage,
	}
}
