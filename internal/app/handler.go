package app

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"tgbot_msg_anon/internal/actions"
	"tgbot_msg_anon/internal/recipients"
	"tgbot_msg_anon/internal/replicas"
	"tgbot_msg_anon/internal/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func commandsHandler(upd *tgbotapi.Update) {
	var ID = upd.Message.From.ID

	recipients.Clear(ID)
	actions.Clear(ID)

	/*КОМАНДЫ ДЛЯ ВСЕХ ПОЛЬЗОВАТЕЛЕЙ*/
	{
		//start
		if ok, _ := regexp.MatchString("^/start$", upd.Message.Text); ok {
			msg := tgbotapi.NewMessage(ID, replicas.Get("my_link", ID))
			bot.Send(msg)
			return
		}

		//start with ID
		if ok, _ := regexp.MatchString("^/start [0-9]*$", upd.Message.Text); ok {
			recipID, _ := strconv.ParseInt(strings.Split(upd.Message.Text, " ")[1], 10, 64)

			recipients.Set(ID, recipID)
			actions.Set(ID, actions.WAIT_ANON_MSG)

			msg := tgbotapi.NewMessage(ID, replicas.Get("wait_anon_msg", ID))
			msg.ParseMode = "html"

			bot.Send(msg)
			return
		}

		//get personal link
		if ok, _ := regexp.MatchString("^/get$", upd.Message.Text); ok {
			msg := tgbotapi.NewMessage(ID, replicas.Get("my_link", ID))
			bot.Send(msg)
			return
		}

		//узнать id
		if ok, _ := regexp.MatchString("^/id$", upd.Message.Text); ok {
			msg := tgbotapi.NewMessage(ID, replicas.Get("my_id", ID))
			msg.ParseMode = "html"
			bot.Send(msg)
			return
		}
	}

	//Проверка на админа
	if !isAdmin(upd) {
		msg := tgbotapi.NewMessage(ID, "Если ты хочешь отправить кому-то сообщение, то перейди по его ссылке")
		bot.Send(msg)
		return
	}

	/*КОМАНДЫ ТОЛЬКО ДЛЯ АДМИНОВ*/
	{
		//отменить действия и вернуться в меню
		if ok, _ := regexp.MatchString("^/cancel$", upd.Message.Text); ok {
			msg := tgbotapi.NewMessage(ID, replicas.Get("cancel_action"))
			bot.Send(msg)
			return
		}

		//рассылка
		if ok, _ := regexp.MatchString("^/mailing$", upd.Message.Text); ok {
			msg := tgbotapi.NewMessage(ID, replicas.Get("new_mailing"))
			bot.Send(msg)

			actions.Set(ID, actions.NEW_MAILING)
			return
		}

		//кол-во пользователей
		if ok, _ := regexp.MatchString("^/amount$", upd.Message.Text); ok {
			msg := tgbotapi.NewMessage(ID, replicas.Get("amount_users", len(repository.Users.GetAllTelegramID())))
			msg.ParseMode = "html"
			bot.Send(msg)
			return
		}

		//панель команд
		if ok, _ := regexp.MatchString("^/panel$", upd.Message.Text); ok {
			msg := tgbotapi.NewMessage(ID, replicas.Get("commands_panel"))
			bot.Send(msg)
			return
		}
	}

	msg := tgbotapi.NewMessage(ID, "Даже у админов нет таких команд.. будь внимательнее :/")
	bot.Send(msg)

}

func textHandler(upd *tgbotapi.Update) {
	var ID = upd.Message.From.ID

	//Отправка анонимного сообщения
	if actions.If(ID, actions.WAIT_ANON_MSG) {
		recipID := recipients.Get(ID)

		msg := tgbotapi.NewMessage(recipID, replicas.Get("send_anon_msg", upd.Message.MessageID, upd.Message.Text))
		msg.ParseMode = "html"

		if _, err := bot.Send(msg); err != nil {
			fmt.Println(err)
			msg = tgbotapi.NewMessage(ID, replicas.Get("send_anon_msg_error_block"))
			msg.ParseMode = "html"
			bot.Send(msg)
		} else {
			msg = tgbotapi.NewMessage(ID, replicas.Get("success_send_anon_msg", ID))
			msg.ParseMode = "html"
			bot.Send(msg)
		}

		actions.Clear(ID)
		return
	}

	//рассылка
	if actions.If(ID, actions.NEW_MAILING) {
		ids := repository.Users.GetAllTelegramID()

		var withErr uint
		for _, tgID := range ids {
			msg := tgbotapi.NewMessage(tgID, upd.Message.Text)
			msg.ParseMode = "html"
			if _, err := bot.Send(msg); err != nil {
				withErr++
			}
		}

		msg := tgbotapi.NewMessage(ID, replicas.Get("mailing_report", len(ids), withErr))
		msg.ParseMode = "html"
		bot.Send(msg)

		actions.Clear(ID)
		return
	}

	msg := tgbotapi.NewMessage(upd.Message.From.ID, "Если ты хочешь отправить кому-то сообщение, то перейди по его ссылке")
	bot.Send(msg)
}
