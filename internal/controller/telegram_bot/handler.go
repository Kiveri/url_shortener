package telegram_bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type createShortUrlRequestTg struct {
	LongUrl string `json:"long_url"`
}

type createShortUrlResponseTg struct {
	ShortUrl string `json:"short_url"`
}

func (t *TgController) HandleUpdate() {
	var req createShortUrlRequestTg
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := t.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			switch update.Message.Command() {
			case "start":
				t.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я бот для сокращения ссылок"))
			case "shorten":
				originalUrl := update.Message.CommandArguments()
				url, err := t.urlsUseCase.CreateUrl(originalUrl)
			}

		}
	}
}
