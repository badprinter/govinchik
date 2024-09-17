package telegram

import "gopkg.in/telebot.v3"

type handles struct {
}

func (h *handles) start(c telebot.Context) error {
	return c.Send("Hello world")
}
