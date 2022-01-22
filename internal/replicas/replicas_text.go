package replicas

var texts = map[string]string{
	"my_link": "Лови ссылку для получения анонимных сообщений 🙈\nt.me/%s?start=%d",

	"wait_anon_msg": "Напиши сообщение, которое ты хочешь отправить обладателю данной ссылки <i>(сообщение будет доставлено анонимно)</i>.\n\n<b>❤️ Хочешь тоже получать анонимные сообщения?</b> Отправь друзьям свою ссылку: t.me/%s?start=%d",

	"success_send_anon_msg": "<b>Сообщение доставлено 🥳❤️</b>\nТвоя личная ссылка для получения анонимных сообщений: t.me/%s?start=%d",

	"send_anon_msg":             "<b>Новое сообщение ❤️ (#%d)</b>\n\n%s",
	"send_anon_msg_error_block": "<b>Ошибка отправки сообщения</b>\nСкорее всего, пользователь больше не желает получать анонимные сообщения",

	"new_mailing":    "💰 <b>Введи текст рассылки.</b> Для отмены используй команду /cancel",
	"mailing_report": "<b>Рассылка окончена</b>\nВсего пользователей: <b>%d</b>\nОшибок: <b>%d</b>\nУспех: <b>%d</b>",

	"cancel_action": "👀 <b>Действие отменено, Вы возвращены в меню</b>",

	"my_id": "Брат, вот твой id: <b>%d</b>",

	"bot_stats": "<b>🍻 Статистика бота:</b>\nКол-во юзеров: %d\nКол-во анонимных сообщений: %d",

	"reply_text":                 "<b>Отправь текст ответа анониму на его сообщение</b> <i>(для отмены используй команду /cancel)</i>: ",
	"new_reply_message":          "<b>Ответ на сообщение:</b>\n%s",
	"success_send_reply_message": "<b>Ответ на сообщение успешно доставлен 🥳❤️</b>",

	"commands_panel": `
		⚒ <b>Список доступных команд</b>:
		/id - твой id
		/get - получить личную ссылку 
		/mailing - рассылка по стаду
		/stats - статистика 
	`,
}
