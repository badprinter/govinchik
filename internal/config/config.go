package config

import "log"

// Глобальная переменна которая содежрит в себе все нужные параметры.
var Cfg *Config = NewConfig()

type Config struct {
	Storage  *StorageConfig
	Telegram *TelegramConfig
}

func NewConfig() *Config {

	// Storage
	storage, err := NewStorageConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Telegram bot api
	telegram, err := NewTelegramConfig()
	if err != nil {
		log.Fatalln(err)
	}

	return &Config{
		storage,
		telegram,
	}
}
