package handlers

import (
	"errors"
	"fmt"
	"store_bot/internal/models"

	"github.com/sirupsen/logrus"
)

type Command string

var Start_command Command = "/start"
var Hello_command Command = "/hello"
var Create_collection_command Command = "/create_collection"
var Create_product Command = "/create_product"
var Get_collections Command = "/get_collections"

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
	if update.ModeratorState == nil {
		update.Msg.Text = "Не понимаю"
		return update, nil
	}
	switch *update.ModeratorState {
	case models.AddCollection:
		switch update.ValidationError {
		case TooBig:
			update.Msg.Text = "Слишком большое название для коллекции"
		case TooSmall:
			update.Msg.Text = "Слишком маленькое название для коллекции"
		case CollectionIsAreadyExisten:
			update.Msg.Text = "Такая коллекция уже существует..."
		default:
			if update.CollectionName == nil {
				return nil, errors.New("Collection name not set")
			}
			update.Msg.Text = fmt.Sprintf("Вы успешно создали новую коллекцию: <b>%s</b> 🙀", *update.CollectionName)
		}
	}
	return update, nil
}

func (m *MessageHandler) handleCommand(update *models.UpdateContext) (*models.UpdateContext, error) {
	update.Msg.ReceiverId = update.Update.Message.SenderId

	switch update.Update.Message.Text {
	case string(Start_command):
		update.Msg.Text = fmt.Sprintf(
			`
				Мои команды: 
					%s - просмотреть команды, 
					%s - получить приветствие,
					%s - добавить подборку, 
					%s - добавить товар,
					%s - посмотреть коллекции
			`, Start_command, Hello_command, Create_collection_command, Create_product, Get_collections,
		)
	case string(Hello_command):
		update.Msg.Text = "Привет, друг!"
	case string(Create_product):
		update.Msg.Text = "Пока в разработке... :("
	case string(Create_collection_command):
		update.Msg.Text = "Введи название коллекции: "
	case string(Get_collections):
		if update.ValidationError == NoCollections {
			update.Msg.Text = "Ещё не создано ни одной коллекции..."
		} else {
			update.Msg.Text = m.getMessageTextFromCollections(*update.Collections)
		}
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

func (m *MessageHandler) getMessageTextFromCollections(collections []models.Collection) string {
	text := "<b>Текущие колллекции:</b> \n"
	for i, collection := range collections {
		text += fmt.Sprintf("    <b>%d</b>: %s\n", i+1, collection.Name)
	}
	return text
}
