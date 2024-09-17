package app

import (
	"github.com/badprinter/govinchik/internal/config"
	"github.com/badprinter/govinchik/internal/mylogger"
	"github.com/badprinter/govinchik/internal/storage"
	"github.com/badprinter/govinchik/internal/telegram"
)

var (
	mylog *mylogger.MyLogger      = mylogger.NewMyLogger()
	cfg   *config.Config          = config.NewConfig()
	store *storage.StorageManager = storage.NewStorageManager()
)

type App struct {
	telegram *telegram.TelegramBot
}

func New() *App {
	cfg := config.NewConfig()
	telegram := telegram.NewTelegram(&cfg.Telegram.Token)
	return &App{
		telegram,
	}
}

func (a *App) Start() {
	defer store.Close()
	a.telegram.Start()
}
