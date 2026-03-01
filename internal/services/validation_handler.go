package handlers

import (
	"errors"
	"fmt"
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
	var err error
	if update.Update.Message != nil {
		if update.Update.Message.IsCommand() {
			update, err = v.handleCommand(update)
		} else {
			update, err = v.handleCommand(update)
		}
	} else if update.Update.CallbackQuery != nil {
		update, err = v.handleCallback(update)
	}
	if err != nil {
		return nil, fmt.Errorf("%s.HandleUpdate: %w", v.Name(), err)
	}
	return update, nil
}

func (v *ValidationHandler) Name() string {
	return "ValidationHandler"
}

func (v *ValidationHandler) handleCommand(update *models.UpdateContext) (*models.UpdateContext, error) {
	return update, nil
}

func (v *ValidationHandler) handleCallback(update *models.UpdateContext) (*models.UpdateContext, error) {
	return update, nil
}
func (v *ValidationHandler) handleMessage(update *models.UpdateContext) (*models.UpdateContext, error) {
	switch *update.ModeratorState {
	case models.AddCollection:
		err := v.validateField(
			update.Update.Message.Text,
			v.cfg.MaxCardNameLength,
			v.cfg.MinCardNameLength,
		)
		if err == UnknownType {
			return nil, err
		}
		update.ValidationError = err
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
