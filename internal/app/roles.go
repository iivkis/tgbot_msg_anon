package app

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var admins = []int64{
	444650304, //@iivkis
	610869869, //@dangerrrrrrr
}

func isAdmin(upd *tgbotapi.Update) bool {
	for _, id := range admins {
		if id == upd.Message.From.ID {
			return true
		}
	}
	return false
}
