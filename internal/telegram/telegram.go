package telegram

import (
	"github.com/badprinter/govinchik/internal/config"
	"github.com/badprinter/govinchik/internal/mylogger"
	"gopkg.in/telebot.v3"
	"log"
	"time"
)

var (
	mylog *mylogger.MyLogger     = mylogger.NewMyLogger()
	cfg   *config.TelegramConfig = config.NewConfig().Telegram
)

type TelegramBot struct {
	handles
	api *telebot.Bot
}

func NewTelegramBot() *TelegramBot {

	t, err := telebot.NewBot(telebot.Settings{
		Token:  cfg.Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	return &TelegramBot{api: t}
}

func (t *TelegramBot) Start() {
	t.api.Start()
}

func (t *TelegramBot) Stop() {
	t.api.Stop()
}

func (t *TelegramBot) registerHandlers() {
	t.api.Handle("/start", t.start)
}
