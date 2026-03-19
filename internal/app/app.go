package app

import (
	"fmt"
	"os"
	"store_bot/internal/database"
	"store_bot/internal/models"
	"store_bot/internal/repository"
	handlers "store_bot/internal/services"
	"strconv"
	"strings"

	"github.com/glebarez/sqlite"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var _ handlers.UpdateHandler = (*handlers.MainHandler)(nil)
var _ handlers.UpdateHandler = (*handlers.ModeratorHandler)(nil)
var _ handlers.UpdateHandler = (*handlers.MessageHandler)(nil)
var _ handlers.UpdateHandler = (*handlers.ModeratorStateGetterHandler)(nil)
var _ handlers.UpdateHandler = (*handlers.ModeratorStateSetterHandler)(nil)

type BotApp struct {
	mainHandler handlers.UpdateHandler
	bot         *tgbotapi.BotAPI
	cfg         *BotAppConfig
	updatesChan chan (*models.UpdateContext)
}

type BotAppConfig struct {
	WorkersNum      int
	UpdatesChanSize int
	Separator       string
}

func NewBotApp(cfg *BotAppConfig) (*BotApp, error) {
	db, err := database.GetGorm(&gorm.Config{}, sqlite.Open("test.sqlite"))
	if err != nil {
		return nil, fmt.Errorf("Unable to create db connections: %w", err)
	}
	db.Exec("PRAGMA foreign_keys = ON;")

	moderRep, err := repository.NewModeratorRepository(&repository.ModeratorRepositoryConfig{MigrateEnable: true}, db)
	if err != nil {
		return nil, fmt.Errorf("Unable to create ModeratorRepository: %w", err)
	}
	mdIds := os.Getenv("MODERATOR_IDS")
	if mdIds != "" {
		ids, err := parseModeratorIds(mdIds, cfg.Separator)
		if err != nil {
			return nil, fmt.Errorf("Unable to parse ModeratorIds: %w", err)
		}
		for _, id := range ids {
			_, err := moderRep.CreateModerator(&models.Moderator{ID: id})
			if err != nil {
				return nil, fmt.Errorf("Failed to create moderator: %w", err)
			}
		}
	}
	moderStateRep, err := repository.NewModeratorStateRepository(&repository.ModeratorStateRepositoryConfig{
		MigrateEnable: true,
	}, db)
	if err != nil {
		return nil, fmt.Errorf("Unable to create ModeratorStateRepository: %w", err)
	}
	collectionRep, err := repository.NewCollectionRepository(&repository.CollectionRepositoryConfig{
		MigrateEnable: true,
	}, db)
	if err != nil {
		return nil, fmt.Errorf("Unable to create CollectionRepository: %w", err)
	}
	cardRep, err := repository.NewCardRepository(&repository.CardRepositoryConfig{
		MigrateEnable: true,
	}, db)
	if err != nil {
		return nil, fmt.Errorf("Unable to create CardRepository: %w", err)
	}

	moderHandler := handlers.NewModeratorHandler(
		&handlers.ModeratorHandlerConfig{DebugLevel: logrus.ErrorLevel},
		moderRep,
	)
	validationHandler := handlers.NewValidationHandler(handlers.GetDefaultValidationHandlerConfig())

	moderStateGetter := handlers.NewModeratorStateGetterhandler(moderStateRep, &handlers.ModeratorStateGetterConfig{
		DebugLevel: logrus.ErrorLevel,
	})
	moderStateSetter := handlers.NewModeratorStateSetterhandler(moderStateRep, &handlers.ModeratorStateSetterConfig{
		DebugLevel: logrus.ErrorLevel,
	})
	collectionHandler := handlers.NewCollectionHandler(&handlers.CollectionHandlerConfig{
		DebugLevel: logrus.ErrorLevel,
	}, collectionRep)
	cardHandler := handlers.NewCardHandler(&handlers.CardHandlerConfig{
		DebugLevel: logrus.ErrorLevel,
	}, cardRep)

	msgHandler := handlers.NewMessageHandler(
		&handlers.MessageHandlerConfig{DebugLevel: logrus.ErrorLevel},
	)
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		return nil, fmt.Errorf("Unable to create bot: %w", err)
	}
	tgBotMsgHandler := handlers.NewTgBotmessageHandler(
		&handlers.TgBotmessageHandlerConfig{DebugLevel: logrus.ErrorLevel},
		bot,
	)

	mainhandler := handlers.NewMainHandler(
		&handlers.MainHandlerConfig{DebugLevel: logrus.ErrorLevel},
		moderHandler,
		moderStateGetter,
		validationHandler,
		moderStateSetter,
		collectionHandler,
		cardHandler,
		msgHandler,
		tgBotMsgHandler,
	)
	return &BotApp{
		mainHandler: mainhandler,
		bot:         bot,
		cfg:         cfg,
		updatesChan: make(chan *models.UpdateContext, cfg.UpdatesChanSize),
	}, nil
}

func (b *BotApp) HandleUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.bot.GetUpdatesChan(u)

	//Start worker pool
	for i := 0; i < b.cfg.WorkersNum; i++ {
		go func() {
			for update := range b.updatesChan {
				b.mainHandler.HandleUpdate(update)
			}
		}()
	}

	for update := range updates {
		b.updatesChan <- &models.UpdateContext{Update: models.UpdateFromTgUpdate(&update)}
	}
}

func parseModeratorIds(strIds string, sep string) ([]int64, error) {
	strIdsArray := strings.Split(strIds, sep)
	ids := make([]int64, 0, len(strIdsArray))
	for _, strId := range strIdsArray {
		id, err := strconv.ParseInt(strId, 10, 64)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}
