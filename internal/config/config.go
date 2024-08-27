package config

type Config struct {
	Storage  *StorageConfig
	Telegram *TelegramConfig
}

func NewConfig() (*Config, error) {

	// Storage
	storage, err := NewStorageConfig()
	if err != nil {
		return nil, err
	}

	// Telegram bot api
	telegram, err := NewTelegramConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		storage,
		telegram,
	}, nil
}

func (c *Config) String() string {
	return "Telegram tokent: " + c.Telegram.Token +
		"\nIP_Database: " + c.Storage.IP
}
