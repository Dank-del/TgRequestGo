package core

import "github.com/PaulSonOfLars/gotgbot/v2"

func IsRequestChat(chat *gotgbot.Chat) bool {
	for _, b := range Conf.RequestChats {
		if b == chat.Id {
			return true
		}
	}
	return false
}
