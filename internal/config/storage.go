package config

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrorUser        = errors.New("ErrorUserENV")
	ErrorSeret       = errors.New("ErrorSecretENV")
	ErrorIpStorage   = errors.New("ErrorIpStorageENV")
	ErrorPortStorage = errors.New("ErrorPrtENV")
	ErrorDbName      = errors.New("ErrorDbNameENV")
	ErrorSSlMode     = errors.New("ErrorSSlModeENV")
)

type StorageConfig struct {
	// TODO driver
	User             string
	Secret           string
	IP               string
	Port             string
	DB_Name          string
	SSLmode          string
	UrlConnectionStr *string
}

// TODO переписать, MAP
func NewStorageConfig() (*StorageConfig, error) {
	s := &StorageConfig{
		os.Getenv("DB_USER"),     // User
		os.Getenv("DB_PASSWORD"), // Secret
		os.Getenv("DB_IP"),       // IP
		os.Getenv("DB_PORT"),     // PORT
		os.Getenv("DB_DATABASE"), // DB_Name
		os.Getenv("DB_SSLMODE"),  // SSL MODE

		new(string), // Заглушка
	}

	if s.User == "" {
		return nil, ErrorUser
	}

	if s.Secret == "" {
		return nil, ErrorSeret
	}

	if s.IP == "" {
		return nil, ErrorIpStorage
	}

	if s.Port == "" {
		return nil, ErrorPortStorage
	}

	if s.DB_Name == "" {
		return nil, ErrorDbName
	}

	if s.SSLmode == "" {
		return nil, ErrorSSlMode
	}

	s.UrlConnectionStr = s.GetUrlConnect()

	return s, nil
}

func (s *StorageConfig) GetUrlConnect() *string {
	str := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		s.User,
		s.Secret,
		s.IP,
		s.Port,
		s.DB_Name,
		s.SSLmode)
	return &str
}
