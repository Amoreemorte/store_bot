package handlers

import (
	"fmt"
	"store_bot/internal/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ModeratorHandler struct {
	cfg    *ModeratorHandlerConfig
	rep    ModeratorRepository
	logger *logrus.Logger
}

func NewModeratorHandler(cfg *ModeratorHandlerConfig, rep ModeratorRepository) *ModeratorHandler {
	logger := logrus.New()
	logger.SetLevel(cfg.DebugLevel)

	return &ModeratorHandler{
		logger: logger,
		cfg:    cfg,
		rep:    rep,
	}
}

func (m *ModeratorHandler) HandleUpdate(update *models.UpdateContext) (*models.UpdateContext, error) {
	var moderatorId int64
	if update.Update.Message != nil {
		moderatorId = update.Update.Message.SenderId
	} else if update.Update.CallbackQuery != nil {
		moderatorId = update.Update.CallbackQuery.SenderId
	}

	_, err := m.rep.GetModerator(moderatorId)
	if err == nil {
		update.IsModerator = true
		return update, nil
	}
	if err == gorm.ErrRecordNotFound {
		update.IsModerator = false
		return update, nil
	}
	return nil, fmt.Errorf("%s.HandleUpdate: %w", m.Name(), err)
}

func (m *ModeratorHandler) Name() string {
	return "ModeratorHandler"
}
