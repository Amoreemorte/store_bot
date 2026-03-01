package handlers

import (
	"errors"
	"fmt"
	"store_bot/internal/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CollectionHandler struct {
	cfg    *CollectionHandlerConfig
	rep    CollectionHandlerRepository
	logger *logrus.Logger
}

var CollectionIsAreadyExisten error = errors.New("CollectionIsAreadyExisten")
var NoCollections error = errors.New("NoCollections")

func NewCollectionHandler(cfg *CollectionHandlerConfig, rep CollectionHandlerRepository) *CollectionHandler {
	logger := logrus.New()
	logger.SetLevel(cfg.DebugLevel)

	return &CollectionHandler{
		rep:    rep,
		logger: logger,
		cfg:    cfg,
	}
}

func (c *CollectionHandler) HandleUpdate(update *models.UpdateContext) (*models.UpdateContext, error) {
	var err error
	if update.Update.Message != nil {
		if update.Update.Message.IsCommand() {
			update, err = c.handleCommand(update)
		} else {
			update, err = c.handleMessage(update)
		}
	}
	return update, err
}

func (c *CollectionHandler) handleCommand(update *models.UpdateContext) (*models.UpdateContext, error) {
	switch update.Update.Message.Text {
	case string(Get_collections):
		collections, err := c.rep.GetCollections()
		if err == gorm.ErrRecordNotFound || len(collections) == 0 {
			update.ValidationError = NoCollections
			return update, nil
		}
		if err != nil {
			return nil, err
		}
		update.Collections = &collections
	}
	return update, nil
}

func (c *CollectionHandler) handleCallback(update *models.UpdateContext) (*models.UpdateContext, error) {
	return update, nil
}

func (c *CollectionHandler) handleMessage(update *models.UpdateContext) (*models.UpdateContext, error) {
	if update.ValidationError != nil || update.ModeratorState == nil {
		return update, nil
	}
	var err error
	switch *update.ModeratorState {
	case models.AddCollection:
		collection, err := c.rep.GetCollection(update.Update.Message.Text)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
		if collection != nil {
			update.ValidationError = CollectionIsAreadyExisten
		}
		_, err = c.rep.CreateCollection(&models.Collection{Name: update.Update.Message.Text})
		if err == nil {
			fmt.Println("collection succsesfully create!")
		}
		update.CollectionName = &update.Update.Message.Text
	}
	return update, err
}

func (c *CollectionHandler) Name() string {
	return "CollectionHandler"
}
