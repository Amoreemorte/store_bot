package handlers

import (
	"errors"
	"store_bot/internal/models"
)

type ValidationHandler struct {
	cfg *ValidationHandlerConfig
}

func NewValidationHandler(cfg *ValidationHandlerConfig) *ValidationHandler {
	return &ValidationHandler{
		cfg: cfg,
	}
}

var TooBig error = errors.New("TooLong")
var TooSmall error = errors.New("TooShort")
var UnknownType error = errors.New("UnknownType")

func (v *ValidationHandler) HandleUpdate(update *models.UpdateContext) (*models.UpdateContext, error) {
	if !update.IsModerator {
		return update, nil
	}

	return update, nil
}

func (v *ValidationHandler) validateField(field any, max int, min int) error {
	var err error
	switch value := field.(type) {
	case int:
		if value > max {
			err = TooBig
		} else if value < min {
			err = TooSmall
		}
	case string:
		if len(value) > max {
			err = TooBig
		} else if len(value) < min {
			err = TooSmall
		}
	default:
		err = UnknownType
	}
	return err
}
