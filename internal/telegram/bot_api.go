package telegram

import (
	tele "gopkg.in/telebot.v3" // tele - как документации к библиотеке
	"log"
	"time"
)

// botAPI управляет библиотекой "gopkg.in/telebot.v3", другими словами обвяз вокгруг библиотеки для структуры TelegramBot
type botAPI struct {
	api *tele.Bot // api самого бота
}

func newBotAPI(token string) *botAPI {

	defer func() {
		if err := recover(); err != nil {
			log.Fatalln("func telegram.newBotAPI: ", err)
		}
	}()

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	api, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	registerHandlers(api)
	return &botAPI{api}
}
