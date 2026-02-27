package handlers

import (
	"fmt"
	"store_bot/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type TgBotmessageHandler struct {
	logger *logrus.Logger
	bot    *tgbotapi.BotAPI
	cfg    *TgBotmessageHandlerConfig
}

func NewTgBotmessageHandler(cfg *TgBotmessageHandlerConfig, bot *tgbotapi.BotAPI) *TgBotmessageHandler {
	logger := logrus.New()
	logger.SetLevel(cfg.DebugLevel)
	return &TgBotmessageHandler{
		bot:    bot,
		cfg:    cfg,
		logger: logger,
	}
}

func (t *TgBotmessageHandler) HandleUpdate(update *models.UpdateContext) (*models.UpdateContext, error) {
	var msg *tgbotapi.MessageConfig
	if update.Msg != nil {
		msg = models.TgMessageFromMessage(update.Msg)
	}
	_, err := t.bot.Send(msg)
	if err != nil {
		return nil, fmt.Errorf("%s.HandleUpdate: unnable to send msg, %w", t.Name(), err)
	}
	return update, nil
}

func (t *TgBotmessageHandler) Name() string {
	return "TgBotmessageHandler"
}
