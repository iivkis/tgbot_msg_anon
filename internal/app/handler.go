package app

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"tgbot_msg_anon/internal/actions"
	"tgbot_msg_anon/internal/cfg"
	"tgbot_msg_anon/internal/markup"
	"tgbot_msg_anon/internal/recipients"
	"tgbot_msg_anon/internal/replicas"
	"tgbot_msg_anon/internal/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func commandsHandler(upd *tgbotapi.Update) {
	var userID = upd.Message.From.ID

	recipients.Clear(userID)
	actions.Clear(userID)

	/*КОМАНДЫ ДЛЯ ВСЕХ ПОЛЬЗОВАТЕЛЕЙ*/
	{
		//start
		if ok, _ := regexp.MatchString("^/start$", upd.Message.Text); ok {
			msg := tgbotapi.NewMessage(userID, replicas.Get("my_link", cfg.Env.BotUsername, userID))
			bot.Send(msg)
			return
		}

		//start with ID
		if ok, _ := regexp.MatchString("^/start [0-9]*$", upd.Message.Text); ok {
			recipID, _ := strconv.ParseInt(strings.Split(upd.Message.Text, " ")[1], 10, 64)

			recipients.Set(userID, recipID)
			actions.Set(userID, actions.WAIT_ANON_MSG)

			msg := tgbotapi.NewMessage(userID, replicas.Get("wait_anon_msg", cfg.Env.BotUsername, userID))
			msg.ParseMode = "html"

			bot.Send(msg)
			return
		}

		//get personal link
		if ok, _ := regexp.MatchString("^/get$", upd.Message.Text); ok {
			msg := tgbotapi.NewMessage(userID, replicas.Get("my_link", cfg.Env.BotUsername, userID))
			bot.Send(msg)
			return
		}

		//узнать id
		if ok, _ := regexp.MatchString("^/id$", upd.Message.Text); ok {
			msg := tgbotapi.NewMessage(userID, replicas.Get("my_id", userID))
			msg.ParseMode = "html"
			bot.Send(msg)
			return
		}

		//отменить действия и вернуться в меню
		if ok, _ := regexp.MatchString("^/cancel$", upd.Message.Text); ok {
			msg := tgbotapi.NewMessage(userID, replicas.Get("cancel_action"))
			msg.ParseMode = "html"
			bot.Send(msg)
			return
		}
	}

	//Проверка на админа
	if !isAdmin(upd) {
		msg := tgbotapi.NewMessage(userID, "Если ты хочешь отправить кому-то сообщение, то перейди по его ссылке")
		bot.Send(msg)
		return
	}

	/*КОМАНДЫ ТОЛЬКО ДЛЯ АДМИНОВ*/
	{
		//рассылка
		if ok, _ := regexp.MatchString("^/mailing$", upd.Message.Text); ok {
			msg := tgbotapi.NewMessage(userID, replicas.Get("new_mailing"))
			msg.ParseMode = "html"
			bot.Send(msg)

			actions.Set(userID, actions.NEW_MAILING)
			return
		}

		//статистика
		if ok, _ := regexp.MatchString("^/stats$", upd.Message.Text); ok {
			msg := tgbotapi.NewMessage(userID, replicas.Get("bot_stats", repository.Users.Count(), repository.Messages.Count()))
			msg.ParseMode = "html"
			bot.Send(msg)
			return
		}

		//панель команд
		if ok, _ := regexp.MatchString("^/panel$", upd.Message.Text); ok {
			msg := tgbotapi.NewMessage(userID, replicas.Get("commands_panel"))
			msg.ParseMode = "html"
			bot.Send(msg)
			return
		}
	}

	msg := tgbotapi.NewMessage(userID, "Даже у админов нет таких команд.. будь внимательнее :/")
	bot.Send(msg)

}

func textHandler(upd *tgbotapi.Update) {
	var ID = upd.Message.From.ID

	//Отправка анонимного сообщения
	if actions.If(ID, actions.WAIT_ANON_MSG) {
		recipID := recipients.Get(ID)

		msg := tgbotapi.NewMessage(recipID, replicas.Get("send_anon_msg", upd.Message.MessageID, upd.Message.Text))
		msg.ReplyMarkup = markup.InlineAnonMSG(upd.Message.MessageID)
		msg.ParseMode = "html"

		if _, err := bot.Send(msg); err != nil {
			fmt.Println(err)
			msg = tgbotapi.NewMessage(ID, replicas.Get("send_anon_msg_error_block"))
			msg.ParseMode = "html"
			bot.Send(msg)
		} else {
			if err := repository.Messages.Add(upd.Message.MessageID, ID, recipID); err != nil {
				fmt.Println(err)
			}

			msg = tgbotapi.NewMessage(ID, replicas.Get("success_send_anon_msg", cfg.Env.BotUsername, ID))
			msg.ParseMode = "html"

			bot.Send(msg)
		}

		actions.Clear(ID)
		return
	}

	//рассылка
	if actions.If(ID, actions.NEW_MAILING) {
		ids := repository.Users.GetAllTelegramID()

		var withErr int
		for _, tgID := range ids {
			msg := tgbotapi.NewMessage(tgID, upd.Message.Text)
			msg.ParseMode = "html"
			if _, err := bot.Send(msg); err != nil {
				withErr++
			}
		}

		msg := tgbotapi.NewMessage(ID, replicas.Get("mailing_report", len(ids), withErr, len(ids)-withErr))
		msg.ParseMode = "html"
		bot.Send(msg)

		actions.Clear(ID)
		return
	}

	//ответ на сообщение
	if actions.If(ID, actions.REPLY_TO_MESSAGE) {
		replyMsgID := repository.Users.GetReplyMsgID(ID)
		message := repository.Messages.Get(replyMsgID)

		msg := tgbotapi.NewMessage(message.FromID, replicas.Get("new_reply_message", upd.Message.Text))
		msg.ReplyToMessageID = replyMsgID
		msg.ParseMode = "html"

		bot.Send(msg)

		msg = tgbotapi.NewMessage(ID, replicas.Get("success_send_reply_message"))
		msg.ParseMode = "html"
		bot.Send(msg)

		actions.Clear(ID)
		return
	}

	msg := tgbotapi.NewMessage(upd.Message.From.ID, "Если ты хочешь отправить кому-то сообщение, то перейди по его ссылке")
	bot.Send(msg)
}

func callbackHandler(upd *tgbotapi.Update) {
	var ID = upd.CallbackQuery.From.ID

	//callback ответа на сообщение
	if ok, _ := regexp.MatchString("^reply [0-9]*$", upd.CallbackQuery.Data); ok {
		replyMessageID, _ := strconv.Atoi(strings.Split(upd.CallbackData(), " ")[1])

		repository.Users.SetReplyMsgID(ID, replyMessageID)
		actions.Set(ID, actions.REPLY_TO_MESSAGE)

		msg := tgbotapi.NewMessage(ID, replicas.Get("reply_text"))
		msg.ReplyToMessageID = upd.CallbackQuery.Message.MessageID
		msg.ParseMode = "html"
		bot.Send(msg)

		edit := tgbotapi.NewEditMessageText(upd.CallbackQuery.From.ID, upd.CallbackQuery.Message.MessageID, upd.CallbackQuery.Message.Text)
		bot.Send(edit)

		fmt.Println(upd.CallbackQuery.Message.Text)
		return
	}
}
