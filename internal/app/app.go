package app

import (
	"github.com/badprinter/govinchik/internal/config"
	"github.com/badprinter/govinchik/internal/mylogger"
	"github.com/badprinter/govinchik/internal/storage"
	"github.com/badprinter/govinchik/internal/telegram"
	"log"
	"os"
)

var (
	mylog *mylogger.MyLogger = mylogger.NewMyLogger()
	cfg   *config.Config     = config.NewConfig()
)

type App struct {
	telegram *telegram.TelegramBot
}

func New() *App {
	logger := log.New(os.Stdout, "govinchik", log.LstdFlags|log.Lshortfile)
	cfg := config.NewConfig()
	store := storage.NewStorageManager(cfg)
	telegram := telegram.NewTelegram(&cfg.Telegram.Token)
	return &App{
		telegram,
	}
}

func (a *App) Start() {
	a.telegram.Start()
}
