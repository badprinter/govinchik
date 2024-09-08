package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/badprinter/govinchik/internal/config"
	"github.com/badprinter/govinchik/internal/storage"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	storage, err := storage.New(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	storage.Connect()

	fmt.Println(cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(cfg.Telegram.Token, opts...)
	if err != nil {
		log.Panicln("Not init token.\n" + err.Error())
	}

	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("RECOVER: %v\n", err)
		}
	}()
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
	con := update.Message.Contact
	if con != nil {
		fmt.Println(con.FirstName)
	}
}
