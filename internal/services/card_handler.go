package handlers

import (
	"store_bot/internal/models"

	"github.com/sirupsen/logrus"
)

type CardHandler struct {
	cfg    *CardHandlerConfig
	rep    CardHandlerRepository
	logger *logrus.Logger
}

func NewCardHandler(cfg *CardHandlerConfig, rep CardHandlerRepository) *CardHandler {
	logger := logrus.New()
	logger.SetLevel(cfg.DebugLevel)

	return &CardHandler{
		rep:    rep,
		logger: logger,
		cfg:    cfg,
	}
}

func (c *CardHandler) HandleUpdate(update *models.UpdateContext) (*models.UpdateContext, error) {
	var err error
	if update.Update.Message != nil {
		if update.Update.Message.IsCommand() {
			update, err = c.handleCommand(update)
		} else {
			update, err = c.handleMessage(update)
		}
	}
	return update, err
}

func (c *CardHandler) handleCommand(update *models.UpdateContext) (*models.UpdateContext, error) {
	return update, nil
}

func (c *CardHandler) handleCallback(update *models.UpdateContext) (*models.UpdateContext, error) {
	return update, nil
}

func (c *CardHandler) handleMessage(update *models.UpdateContext) (*models.UpdateContext, error) {
	if update.ValidationError != nil || update.ModeratorState == nil {
		return update, nil
	}
	return update, nil
}

func (c *CardHandler) Name() string {
	return "CardHandler"
}
