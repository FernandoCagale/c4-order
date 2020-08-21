package event

import (
	"github.com/FernandoCagale/c4-order/internal/broker/consumer"
)

type OrderEvent struct {
	consumer consumer.Consumer
}

func NewOrder(consumer consumer.Consumer) *OrderEvent {
	return &OrderEvent{
		consumer: consumer,
	}
}

func (event *OrderEvent) MakeEvents() {
	go event.consumer.Consumer()
}
