package environments

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ServiceBusConnection   = ""
	NotificationsQueueName = ""
	TelegramBaseUrl        = ""
	BotKey                 = ""
	ChatID                 = 0
)

func NewConfig() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ServiceBusConnection = os.Getenv("SERVICEBUS_CONNECTION")
	NotificationsQueueName = os.Getenv("NOTIFICATIONS_QUEUE_NAME")
	TelegramBaseUrl = os.Getenv("TELEGRAM_BASE_URL")
	BotKey = os.Getenv("BOT_KEY")
	ChatID, err = strconv.Atoi(os.Getenv("CHAT_ID"))
	if err != nil {
		ChatID = 0
	}
}
