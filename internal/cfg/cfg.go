package cfg

import (
	"fmt"
	"os"
)

var Env struct {
	BotToken    string
	BotUsername string
}

func init() {
	//init env
	{
		Env.BotToken = getFromEnvDefault("TGBOT_MSG_ANON_TOKEN", "5039024123:AAFno2M_B9tqk6F1rkKcXvtKAECge__zFK0")
		Env.BotUsername = getFromEnvDefault("TGBOT_MSG_ANON_USERNAME", "msg_anon_bot")
	}
}

func getFromEnvDefault(envName string, defaultVal string) string {
	val, ok := os.LookupEnv(envName)
	if !ok {
		fmt.Printf("[cfg:error] env `%s` undefined\n", envName)
		return defaultVal
	}
	return val
}
