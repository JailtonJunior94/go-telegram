package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github/jailtonjunior94/go-telegram/src/dtos"
	"github/jailtonjunior94/go-telegram/src/services"
	"log"

	servicebus "github.com/Azure/azure-service-bus-go"
)

func Handle(ctx context.Context, msg *servicebus.Message) error {
	telegram := services.NewTelegramService()

	messageReceive := &dtos.MessageReceive{}
	err := json.Unmarshal(msg.Data, messageReceive)
	if err != nil {
		log.Fatal(err)
	}

	dateFormat := messageReceive.Date.Format("2 Jan 2006 15:04:05")
	message := fmt.Sprintf(`Erro na/n APLICAÇÃO: %s:
									  MÉTODO: %s
									  DATA/HORA: %s`,
		messageReceive.Application,
		messageReceive.Method,
		dateFormat)

	if err = telegram.SendMessage(message); err != nil {
		log.Fatal(err)
	}

	return msg.Complete(ctx)
}
