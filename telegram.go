package notifier

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sync"
)

type TelegramBot struct {
	sync.RWMutex
	bot *tgbotapi.BotAPI
}

func NewTelegramBot(tgBotToken string) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(tgBotToken)
	if err != nil {
		return nil, err
	}
	return &TelegramBot{
		bot: bot,
	}, nil
}

func (tb *TelegramBot) SendFile(chatId int64, path string) error {
	tb.Lock()
	defer tb.Unlock()
	file := tgbotapi.NewDocument(chatId, tgbotapi.FilePath(path))

	_, err := tb.bot.Send(file)
	if err != nil {
		return err
	}

	return nil
}

func (tb *TelegramBot) SendMessage(chatId int64, msg string) error {
	tb.Lock()
	defer tb.Unlock()
	message := tgbotapi.NewMessage(chatId, msg)

	_, err := tb.bot.Send(message)
	if err != nil {
		return err
	}

	return nil
}
