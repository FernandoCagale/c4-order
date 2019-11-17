package event

import (
	"encoding/json"
	"fmt"
	"github.com/FernandoCagale/c4-order/internal/event"
	"github.com/FernandoCagale/c4-order/pkg/domain/order"
	"github.com/FernandoCagale/c4-order/pkg/entity"
	"log"
)

type OrderEvent struct {
	usecase order.UseCase
	event   event.Event
}

func NewOrder(usecase order.UseCase, event event.Event) *OrderEvent {
	return &OrderEvent{
		usecase: usecase,
		event:   event,
	}
}

func (eventOrder *OrderEvent) ProcessOrder()  {
	messages, err := eventOrder.event.Subscribe("topic.ecommerce", "order")
	if err != nil {
		fmt.Println(err.Error())
	}

	for msg := range messages {
		log.Printf("received message: %s, ORDER: %s", msg.UUID, string(msg.Payload))

		var ecommerce entity.Ecommerce

		if err := json.Unmarshal(msg.Payload, &ecommerce); err != nil {
			fmt.Println(err.Error())
			msg.Nack()
			return
		}

		if err = eventOrder.usecase.Create(&ecommerce); err != nil {
			fmt.Println(err.Error())
			msg.Nack()
			return
		}

		msg.Ack()
	}
}