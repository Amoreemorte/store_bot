package handlers

import (
	"store_bot/internal/models"
)

type UpdateHandler interface {
	HandleUpdate(*models.UpdateContext) (*models.UpdateContext, error)

	Name() string
}

var _ UpdateHandler = (*MainHandler)(nil)
var _ UpdateHandler = (*ModeratorHandler)(nil)
var _ UpdateHandler = (*MessageHandler)(nil)

type ModeratorRepository interface {
	CreateModerator(*models.Moderator) (*models.Moderator, error)
	DeleteModerator(int64, bool) error
	GetModerator(int64) (*models.Moderator, error)
}
