package routes

import (
	"TgRequestGo/core"
	"fmt"
	"html"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func requestHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser
	chat := ctx.EffectiveChat
	if !core.IsRequestChat(chat) {
		return ext.EndGroups
	}
	args := ctx.Args()
	if len(args) == 1 {
		return ext.EndGroups
	}
	txt := fmt.Sprintf("<b>%s</b> requested\n<code>%s</code>", html.EscapeString(user.FirstName), html.EscapeString(strings.Join(args[1:], " ")))
	_, err := b.SendMessage(core.Conf.AdminChat, txt, &gotgbot.SendMessageOpts{ParseMode: "html"})
	return err
}
