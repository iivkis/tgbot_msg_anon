package replicas

var texts = map[string]string{
	"my_link": "Лови ссылку для получения анонимных сообщений 🙈\nt.me/msg_anon_bot?start=%d",

	"wait_anon_msg": "Напиши сообщение, которое ты хочешь отправить обладателю данной ссылки <i>(сообщение будет доставлено анонимно)</i>.\n\n<b>❤️ Хочешь тоже получать анонимные сообщения? Отправь друзьям свою ссылку:</b> t.me/msg_anon_bot?start=%d",

	"success_send_anon_msg": "<b>Сообщение доставлено 🥳❤️</b>\nТвоя личная ссылка для получения анонимных сообщений: \nt.me/msg_anon_bot?start=%d",

	"send_anon_msg":             "<b>Новое сообщение ❤️ (#%d)</b>\n%s",
	"send_anon_msg_error_block": "<b>Ошибка отправки сообщения</b>\nСкорее всего, пользователь больше не желает получать анонимные сообщения",

	"new_mailing":    "Введи текст рассылки. Для отмены используй команду /cancel",
	"mailing_report": "<b>Рассылка окончена</b>\nВсего пользователей: <b>%d</b>\nКол-во ошибок при отправке: <b>%d</b>",

	"cancel_action": "Действие отменено, ты телепортнут в меню",

	"my_id": "Брат, вот твой id: <b>%d</b>",

	"amount_users": "Кол-во юзеров: <b>%d</b>",

	"commands_panel": `
		/panel - панель команд
		/get - получить личную ссылку 
		/id - твой id
		/mailing - рассылка по стаду
		/amount - кол-во голов в стаде
	`,
}