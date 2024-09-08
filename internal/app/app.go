package app

import (
	"fmt"
	"github.com/badprinter/govinchik/internal/config"
	"github.com/badprinter/govinchik/internal/models"
	"github.com/badprinter/govinchik/internal/storage"
	"github.com/badprinter/govinchik/internal/telegram"
	"log"
	"os"
)

type App struct {
	logger   *log.Logger
	config   *config.Config
	store    *storage.ManagerStorage
	telegram *telegram.TelegramBot
}

func New() *App {
	logger := log.New(os.Stdout, "govinchik", log.LstdFlags|log.Lshortfile)
	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatal(err)
	}
	store, err := storage.NewStorageManager(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	tel, err := telegram.NewTelegramBot(cfg.Telegram.Token)
	if err != nil {
		log.Fatalln(err)
	}
	return &App{
		logger,
		cfg,
		store,
		tel,
	}
}

func (a *App) Start() {
	err := a.store.PingDB()
	if err != nil {
		log.Fatal(err)
	}

	u := &models.User{
		Id:         0,
		Name:       "Egrick",
		TelegramId: 200,
		IsDeleted:  false,
	}
	u = a.store.CreateUser(u)

	user, err := a.store.GetUserByTelegramID(u.TelegramId)
	if err != nil {
		a.logger.Fatal(err)
	}
	fmt.Println(user)
}
