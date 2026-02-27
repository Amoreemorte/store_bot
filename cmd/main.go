package main

import (
	"fmt"
	"os"
	"store_bot/internal/app"
)

func main() {
	app, err := app.NewBotApp(&app.BotAppConfig{
		WorkersNum:      20,
		UpdatesChanSize: 40,
	})
	if err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		os.Exit(1)
	}
	app.HandleUpdates()
}
