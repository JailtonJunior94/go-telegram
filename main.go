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
}
