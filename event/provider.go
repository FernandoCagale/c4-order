package event

import (
	"github.com/FernandoCagale/c4-order/internal/broker/consumer"
	"github.com/FernandoCagale/c4-order/internal/broker/producer"
	"github.com/FernandoCagale/c4-order/pkg/domain/order"
	"github.com/google/wire"
)

var Set = wire.NewSet(NewOrder, order.Set, consumer.Set, producer.Set)
