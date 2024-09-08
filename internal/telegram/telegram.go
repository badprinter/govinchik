package telegram

type TelegramBot struct {
	token *string
}

func NewTelegramBot(tokent string) (*TelegramBot, error) {
	return &TelegramBot{token: &tokent}, nil
}
