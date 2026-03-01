package handlers

import (
	"store_bot/internal/models"
	"store_bot/internal/repository"
)

type UpdateHandler interface {
	HandleUpdate(*models.UpdateContext) (*models.UpdateContext, error)
	Name() string
}

type ModeratorRepository interface {
	CreateModerator(*models.Moderator) (*models.Moderator, error)
	DeleteModerator(int64, bool) error
	GetModerator(int64) (*models.Moderator, error)
}

var _ ModeratorRepository = (*repository.ModeratorRepository)(nil)
