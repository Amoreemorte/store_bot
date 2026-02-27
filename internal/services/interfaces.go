package handlers

import "store_bot/internal/models"

type UpdateHandler interface {
	HandleUpdate(*models.UpdateContext) (*models.UpdateContext, error)

	Name() string
}

var _ UpdateHandler = (*MainHandler)(nil)
