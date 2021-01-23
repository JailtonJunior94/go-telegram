package main

import (
	"context"
	"fmt"
	"github/jailtonjunior94/go-telegram/src/bus"
	"github/jailtonjunior94/go-telegram/src/environments"
	"github/jailtonjunior94/go-telegram/src/handlers"
	"log"
	"time"

	servicebus "github.com/Azure/azure-service-bus-go"
)

func main() {
	environments.NewConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	fmt.Print(cancel)

	bus, err := bus.NewServiceBusConnection(environments.ServiceBusConnection, environments.NotificationsQueueName)
	if err != nil {
		log.Fatal(err)
	}

	var handle servicebus.HandlerFunc = handlers.Handle
	if err := bus.Queue.Receive(ctx, handle); err != nil {
		log.Fatal(err)
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	// defer cancel()

	// connectionString := "Endpoint=sb://notifications-telegram.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=Aa8i+TuYJtZouin8jm92HGfDtnKukIILb6y3iXIdrl0="

	// nameSpace, err := servicebus.NewNamespace(servicebus.NamespaceWithConnectionString(connectionString))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// queue, err := nameSpace.NewQueue("notifications")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // defer queue.Close()

	// body := sendMessageReqBody{
	// 	ChatID: 1,
	// 	Text:   "asa",
	// }

	// data, _ := json.Marshal(body)
	// message := servicebus.Message{
	// 	Data: data,
	// }

	// s, err := queue.ScheduleAt(ctx, time.Now().Add(10*time.Minute), &message)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Print(s)

	// if err := queue.Send(ctx, &message); err != nil {
	// 	log.Fatal(err)
	// }

	// if err := queue.CancelScheduled(ctx, s[0]); err != nil {
	// 	log.Fatal(err)
	// }
}
