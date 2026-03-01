package app

import (
	"fmt"
	"os"
	"store_bot/internal/database"
	"store_bot/internal/models"
	"store_bot/internal/repository"
	handlers "store_bot/internal/services"

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
	moderHandler := handlers.NewModeratorHandler(
		&handlers.ModeratorHandlerConfig{DebugLevel: logrus.ErrorLevel},
		moderRep,
	)
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
		moderHandler, msgHandler, tgBotMsgHandler,
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
