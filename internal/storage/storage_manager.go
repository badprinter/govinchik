package storage

import (
	"errors"
	"github.com/badprinter/govinchik/internal/config"
	"github.com/badprinter/govinchik/internal/models"
)

type ManagerStorage struct {
	cfg *config.Config
	db  *StorageAPI

	userManager UserManager
}

func NewStorageManager(cfg *config.Config) (*ManagerStorage, error) {
	// TODO при иницилизации переменой может не хватить системных ресурсах, добовить обработку
	s := &ManagerStorage{
		cfg,
		NewStorageAPI(cfg),

		UserManager{},
	}
	return s, nil
}

func (s *ManagerStorage) GetUserByID(id int64) (*models.User, error) {
	user := s.userManager.getUserByID(id, s.cfg.Storage.UrlConnectionStr)
	if user == nil {
		return nil, errors.New("Не удалось получить user. по id")
	}

	return user, nil
}
func (s *ManagerStorage) GetUserByTelegramID(id int64) (*models.User, error) {
	user := s.userManager.getUserByTelegramId(id, s.cfg.Storage.UrlConnectionStr)
	if user == nil {
		return nil, errors.New("Не удалось получить user через Telegram id")
	}
	return user, nil
}
func (s *ManagerStorage) CreateUser(user *models.User) *models.User {
	return s.userManager.createUser(user, s.cfg.Storage.UrlConnectionStr)
}
func (s *ManagerStorage) PingDB() error {
	err := s.db.Ping()
	if err != nil {
		return err
	}
	return nil
}
