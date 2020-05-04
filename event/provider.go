package event

import (
	"github.com/FernandoCagale/c4-order/pkg/domain/order"
	eventImp "github.com/FernandoCagale/c4-order/internal/event"
	"github.com/google/wire"
)

var Set = wire.NewSet(NewOrder, order.Set, eventImp.Set)
