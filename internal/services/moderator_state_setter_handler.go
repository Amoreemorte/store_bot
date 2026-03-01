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
	var err error
	if update.Update.Message != nil {
		if update.Update.Message.IsCommand() {
			update, err = m.handleCommand(update)
		} else {
			update, err = m.handleCommand(update)
		}
	} else if update.Update.CallbackQuery != nil {
		update, err = m.handleCallback(update)
	}
	return update, err
}

func (m *ModeratorStateSetterHandler) Name() string {
	return "ModeratorStateSetterHandler"
}

func (m *ModeratorStateSetterHandler) handleCommand(update *models.UpdateContext) (*models.UpdateContext, error) {
	var err error
	switch update.Update.Message.Text {
	case string(Create_collection_command):
		if update.ModeratorState == nil {
			_, err = m.rep.CreateModeratorState(
				&models.ModeratorState{
					ModeratorId: update.Update.Message.SenderId,
					State:       models.AddCollection,
				},
			)
		} else {
			err = m.rep.UpdateModeratorState(
				update.Update.Message.SenderId,
				models.AddCollection,
			)
		}
	default:
		if update.ModeratorState == nil {
			_, err = m.rep.CreateModeratorState(
				&models.ModeratorState{
					ModeratorId: update.Update.Message.SenderId,
					State:       models.NoAction,
				},
			)
		} else {
			err = m.rep.UpdateModeratorState(
				update.Update.Message.SenderId,
				models.NoAction,
			)
		}
	}
	return update, err
}

func (m *ModeratorStateSetterHandler) handleCallback(update *models.UpdateContext) (*models.UpdateContext, error) {
	return update, nil
}
func (m *ModeratorStateSetterHandler) handleMessage(update *models.UpdateContext) (*models.UpdateContext, error) {
	var err error
	if update.ModeratorState == nil {
		_, err = m.rep.CreateModeratorState(
			&models.ModeratorState{
				ModeratorId: update.Update.Message.SenderId,
				State:       models.NoAction,
			},
		)
	} else {
		err = m.rep.UpdateModeratorState(
			update.Update.Message.SenderId,
			models.NoAction,
		)
	}
	return update, err
}
