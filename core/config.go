package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type BotConfig struct {
	BotToken     string  `json:"bot_token"`
	AdminChat    int64   `json:"admin_chat"`
	RequestChats []int64 `json:"request_chats"`
	DropUpdates  bool    `json:"drop_updates"`
}

func NewBotConfig() *BotConfig {
	return &BotConfig{}
}

func (ac *BotConfig) ReadFile(fileName string) {
	b, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(b, &ac)
	if err != nil {
		log.Fatal(err)
	}
}
