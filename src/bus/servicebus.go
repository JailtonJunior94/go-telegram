package bus

import (
	servicebus "github.com/Azure/azure-service-bus-go"
)

type ServiceBus struct {
	Namespace *servicebus.Namespace
	Queue     *servicebus.Queue
}

func NewServiceBusConnection(connectionString, queueName string) (ServiceBus, error) {
	n, err := servicebus.NewNamespace(servicebus.NamespaceWithConnectionString(connectionString))
	if err != nil {
		return ServiceBus{}, err
	}

	q, err := n.NewQueue(queueName)
	if err != nil {
		return ServiceBus{}, nil
	}

	return ServiceBus{Namespace: n, Queue: q}, nil
}
