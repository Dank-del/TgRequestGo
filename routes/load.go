package routes

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func Load(d *ext.Dispatcher) {
	d.AddHandler(handlers.NewCommand("request", requestHandler))
}
