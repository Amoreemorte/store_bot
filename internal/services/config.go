package handlers

import "github.com/sirupsen/logrus"

type MainHandlerConfig struct {
	DebugLevel logrus.Level
}

type ModeratorHandlerConfig struct {
	DebugLevel logrus.Level
}

type MessageHandlerConfig struct {
	DebugLevel logrus.Level
}

type TgBotmessageHandlerConfig struct {
	DebugLevel logrus.Level
}

type ModeratorStateGetterConfig struct {
	DebugLevel logrus.Level
}

type ModeratorStateSetterConfig struct {
	DebugLevel logrus.Level
}
