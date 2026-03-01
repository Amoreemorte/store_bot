package handlers

import (
	"fmt"
	"store_bot/internal/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ModeratorStateGetterHandler struct {
	rep    ModeratorStateGetterRepository
	cfg    *ModeratorStateGetterConfig
	logger *logrus.Logger
}

func NewModeratorStateGetterhandler(rep ModeratorStateGetterRepository, cfg *ModeratorStateGetterConfig) *ModeratorStateGetterHandler {
	logger := logrus.New()
	logger.SetLevel(cfg.DebugLevel)

	return &ModeratorStateGetterHandler{
		logger: logger,
		cfg:    cfg,
		rep:    rep,
	}
}

func (m *ModeratorStateGetterHandler) HandleUpdate(update *models.UpdateContext) (*models.UpdateContext, error) {
	if !update.IsModerator {
		return update, nil
	}
	var moderatorId int64
	if update.Update.Message != nil {
		moderatorId = update.Update.Message.SenderId
	} else if update.Update.CallbackQuery != nil {
		moderatorId = update.Update.CallbackQuery.SenderId
	}
	state, err := m.rep.GetModeratorState(moderatorId)
	if err == nil {
		update.ModeratorState = &state.State
		return update, nil
	}
	if err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("%s.HandleUpdate: %w", m.Name(), err)
	}
	return update, nil
}

func (m *ModeratorStateGetterHandler) Name() string {
	return "ModeratorStateGetterHandler"
}
