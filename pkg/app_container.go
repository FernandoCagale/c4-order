package pkg

import (
	eventBoot "github.com/FernandoCagale/c4-order/api/event"
	"github.com/FernandoCagale/c4-order/api/handlers"
	"github.com/FernandoCagale/c4-order/api/routers"
	"github.com/FernandoCagale/c4-order/internal/event"
	"github.com/FernandoCagale/c4-order/pkg/domain/order"
	"github.com/google/wire"
)

var Container = wire.NewSet(order.Set, handlers.Set, routers.Set, event.Set, eventBoot.Set)
