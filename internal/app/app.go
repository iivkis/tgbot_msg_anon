package app

import (
	"fmt"
	"tgbot_msg_anon/internal/cfg"
	"tgbot_msg_anon/internal/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func Launch() {
	//init bot
	botAPI, err := tgbotapi.NewBotAPI(cfg.Env.BotToken)
	if err != nil {
		panic(err)
	}
	// botAPI.Debug = true

	//set global bot API
	bot = botAPI

	//settings config
	uc := tgbotapi.NewUpdate(0)
	uc.Timeout = 30

	updates := bot.GetUpdatesChan(uc)

	//listen channel and handle messages
	fmt.Println("[bot:notify] Батька ждёт сообщений")
	for upd := range updates {
		//если пришло сообщение
		if upd.Message != nil {
			//регаем юзера в базе
			usr := repository.UserModel{TelegramID: upd.Message.From.ID}
			repository.Users.Create(&usr)

			fmt.Printf("[bot:msg] From: http://t.me/%.20s | Text:\"%s\"\n", upd.Message.Chat.UserName, upd.Message.Text)
			if upd.Message.IsCommand() {
				go commandsHandler(&upd)
			} else if upd.Message.Text != "" {
				go textHandler(&upd)
			}
			continue
		}

		//если пришло уведомление
		if upd.CallbackQuery != nil {
			fmt.Printf("[bot:callback] From: http://t.me/%.20s | Data:\"%s\"\n", upd.CallbackQuery.From.UserName, upd.CallbackData())
			go callbackHandler(&upd)
		}
	}
}
