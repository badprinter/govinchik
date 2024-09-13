package telegram

import (
	"fmt"
	"github.com/badprinter/govinchik/internal/storage"
	tele "gopkg.in/telebot.v3"
	"log"
)

func registerHandlers(t *tele.Bot) {
	t.Handle("/start", startHandler)

}

func startHandler(c tele.Context) error {
	teleID := c.Chat().ID
	user, _ := storage.Store.GetUserByID(teleID)

	if user == nil {
		log.Printf("user not found by Telegram ID: %d", teleID)
		user.Name = c.Chat().FirstName
		user.TelegramId = teleID
		user = storage.Store.CreateUser(user)

		if user == nil {
			log.Printf("user not created by Telegram ID: %d", teleID)
			return nil
		}
	}

	log.Printf("user found of created by Telegram ID: %d", teleID)
	msg := fmt.Sprintf("Привет, %s!", user.Name)
	return c.Send(msg)
}
