package telegram

import "github.com/badprinter/govinchik/internal/storage"

var (
	// Основная переменна для работы с базой-данных
	store *storage.StorageManager = storage.NewStorageManager()
)

type TelegramBot struct {
	token *string
	bot   *botAPI
}

func NewTelegram(token *string) *TelegramBot {
	bot := newBotAPI(*token)
	return &TelegramBot{
		token: token,
		bot:   bot,
	}
}

func (t *TelegramBot) Start() {
	t.bot.api.Start()
}
