package routers

import (
	"github.com/FernandoCagale/c4-order/api/event"
	"github.com/FernandoCagale/c4-order/api/handlers"
	"github.com/gorilla/mux"
	"time"
)

type SystemRoutes struct {
	healthHandler *handlers.HealthHandler
	orderHandler  *handlers.OrderHandler
	orderEvent    *event.OrderEvent
}

func (routes *SystemRoutes) MakeEvents() {
	time.Sleep(5 * time.Second)

	routes.orderEvent.ProcessOrder()
}

func (routes *SystemRoutes) MakeHandlers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", routes.healthHandler.Health).Methods("GET")
	r.HandleFunc("/order", routes.orderHandler.Create).Methods("POST")

	return r
}

func NewSystem(healthHandler *handlers.HealthHandler, orderHandler *handlers.OrderHandler, orderEvent *event.OrderEvent) *SystemRoutes {
	return &SystemRoutes{
		healthHandler: healthHandler,
		orderHandler:  orderHandler,
		orderEvent:    orderEvent,
	}
}
