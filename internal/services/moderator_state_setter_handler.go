package handlers

import (
	"store_bot/internal/models"

	"github.com/sirupsen/logrus"
)

type ModeratorStateSetterHandler struct {
	rep    ModeratorStateSetterRepository
	cfg    *ModeratorStateSetterConfig
	logger *logrus.Logger
}

func NewModeratorStateSetterhandler(rep ModeratorStateSetterRepository, cfg *ModeratorStateSetterConfig) *ModeratorStateSetterHandler {
	logger := logrus.New()
	logger.SetLevel(cfg.DebugLevel)

	return &ModeratorStateSetterHandler{
		logger: logger,
		cfg:    cfg,
		rep:    rep,
	}
}

func (m *ModeratorStateSetterHandler) HandleUpdate(update *models.UpdateContext) (*models.UpdateContext, error) {
	if !update.IsModerator {
		return update, nil
	}
	return update, nil
}

func (m *ModeratorStateSetterHandler) Name() string {
	return "ModeratorStateSetterHandler"
}
