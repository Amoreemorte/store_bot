package handlers

import (
	"store_bot/internal/models"

	"github.com/sirupsen/logrus"
)

type MainHandler struct {
	handlers []UpdateHandler

	logger *logrus.Logger
}

func NewMainHandler(cfg *MainHandlerConfig, handlers ...UpdateHandler) *MainHandler {
	logger := logrus.New()
	logger.SetLevel(cfg.DebugLevel)

	return &MainHandler{
		handlers: handlers,
		logger:   logger,
	}
}

func (m *MainHandler) HandleUpdate(update *models.UpdateContext) (*models.UpdateContext, error) {
	var err error
	for _, handler := range m.handlers {
		update, err = handler.HandleUpdate(update)
		if err != nil {
			m.logger.Error(
				m.Name(),
				":",
				err.Error(),
			)
			break
		}
	}
	return update, err
}

func (m *MainHandler) Name() string {
	return "MainHandler"
}
