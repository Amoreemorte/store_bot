package handlers

import (
	"fmt"
	"store_bot/internal/models"

	"github.com/sirupsen/logrus"
)

type Command string

// TODO: change temporarily commands
var Start_command Command = "/start"
var Hello_command Command = "/hello"
var Create_set_command Command = "/create_set"
var Create_product Command = "/create_product"

type MessageHandler struct {
	cfg    *MessageHandlerConfig
	logger *logrus.Logger
}

func NewMessageHandler(cfg *MessageHandlerConfig) *MessageHandler {
	logger := logrus.New()
	logger.SetLevel(cfg.DebugLevel)

	return &MessageHandler{
		logger: logger,
		cfg:    cfg,
	}
}

func (m *MessageHandler) HandleUpdate(update *models.UpdateContext) (*models.UpdateContext, error) {
	var err error
	update.Msg = &models.MessageConfig{}
	if update.Update.Message != nil {
		if !update.IsModerator {
			update.Msg.ReceiverId = update.Update.Message.SenderId
			update.Msg.Text = "Вы не являетесь администратором бота..."
			return update, nil
		}

		if update.Update.Message.IsCommand() {
			update, err = m.handleCommand(update)
		} else {
			update, err = m.handleMessage(update)
		}
	} else if update.Update.CallbackQuery != nil {
		update, err = m.handleCallback(update)
	}
	if err != nil {
		return nil, fmt.Errorf("%s.HandleUpdate: %w", m.Name(), err)
	}
	return update, nil
}

func (m *MessageHandler) handleMessage(update *models.UpdateContext) (*models.UpdateContext, error) {
	update.Msg.ReceiverId = update.Update.Message.SenderId
	update.Msg.Text = "Пока разрабатываюсь..."
	return update, nil
}

func (m *MessageHandler) handleCommand(update *models.UpdateContext) (*models.UpdateContext, error) {
	update.Msg.ReceiverId = update.Update.Message.SenderId

	switch update.Update.Message.Text {
	case string(Start_command):
		update.Msg.Text = fmt.Sprintf(
			`
				Мои комманды: 
					%s - просмотреть команды, 
					%s - получить приветствие,
					%s - добавить подборку, 
					%s - добавить товар
			`, Start_command, Hello_command, Create_set_command, Create_product,
		)
	case string(Hello_command):
		update.Msg.Text = "Привет, друг!"
	case string(Create_product), string(Create_set_command):
		update.Msg.Text = "Пока в разработке... :("
	default:
		update.Msg.Text = "Не знаю такой команды..."
	}
	return update, nil
}

func (m *MessageHandler) handleCallback(update *models.UpdateContext) (*models.UpdateContext, error) {
	return update, nil
}

func (m *MessageHandler) Name() string {
	return "MessageHandler"
}
