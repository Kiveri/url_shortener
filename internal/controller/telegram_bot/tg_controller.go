package telegram_bot

import (
	"url/internal/domain/model"
	"url/internal/usecase"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type (
	TgController struct {
		bot         tgbotapi.BotAPI
		urlsUseCase urlsUseCase
		baseUrl     string
	}

	urlsUseCase interface {
		CreateUrl(req usecase.CreateUrlReq) (*model.URL, error)
		FindUrl(req usecase.FindUrlReq) (*model.URL, error)
	}
)

func NewController(urlsUseCase urlsUseCase, baseUrl string, bot tgbotapi.BotAPI) *TgController {
	return &TgController{
		bot:         bot,
		urlsUseCase: urlsUseCase,
		baseUrl:     baseUrl,
	}
}
