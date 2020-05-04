package routers

import (
	"github.com/FernandoCagale/c4-order/api/handlers"
	"github.com/gorilla/mux"
)

type SystemRoutes struct {
	healthHandler *handlers.HealthHandler
	orderHandler  *handlers.OrderHandler
}

func (routes *SystemRoutes) MakeHandlers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", routes.healthHandler.Health).Methods("GET")
	r.HandleFunc("/orders", routes.orderHandler.Create).Methods("POST")
	r.HandleFunc("/orders", routes.orderHandler.FindAll).Methods("GET")
	r.HandleFunc("/orders/{id}", routes.orderHandler.FindById).Methods("GET")
	r.HandleFunc("/orders/{id}", routes.orderHandler.DeleteById).Methods("DELETE")

	return r
}

func NewSystem(healthHandler *handlers.HealthHandler, orderHandler *handlers.OrderHandler) *SystemRoutes {
	return &SystemRoutes{
		healthHandler: healthHandler,
		orderHandler:  orderHandler,
	}
}
