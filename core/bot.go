package core

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func BotInit(config *BotConfig) (b *gotgbot.Bot, updater ext.Updater, err error) {
	b, err = gotgbot.NewBot(config.BotToken, &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	if err != nil {
		log.Println(fmt.Sprintf("Failed to create new bot due to %s", err.Error()))
		return nil, updater, err
	}
	updater = ext.NewUpdater(nil)
	err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: config.DropUpdates})
	if err != nil {
		log.Println(fmt.Sprintf("Failed to start polling due to %s", err.Error()))
	}
	log.Println(fmt.Sprintf("%s has started | ID: %d", b.Username, b.Id))
	return b, updater, nil
}
