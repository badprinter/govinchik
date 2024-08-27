package config

import (
	"errors"
	"os"
)

type TelegramConfig struct {
	Token string
}

var ErrorTelegramToken = errors.New("ErrorTelegramTokenENV")

func NewTelegramConfig() (*TelegramConfig, error) {

	t := &TelegramConfig{
		os.Getenv("TOKEN"),
	}

	if t.Token == "" {
		return nil, ErrorTelegramToken
	}

	return t, nil
}
