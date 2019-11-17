package handlers

import (
	"encoding/json"
	"github.com/FernandoCagale/c4-order/api/render"
	"github.com/FernandoCagale/c4-order/internal/errors"
	"github.com/FernandoCagale/c4-order/pkg/domain/order"
	"github.com/FernandoCagale/c4-order/pkg/entity"
	"net/http"
)

type OrderHandler struct {
	usecase order.UseCase
}

func NewOrder(usecase order.UseCase) *OrderHandler {
	return &OrderHandler{
		usecase: usecase,
	}
}

func (handler *OrderHandler) Orders(w http.ResponseWriter, r *http.Request) {
	render.Response(w, map[string]bool{"ok": true}, http.StatusOK)
}

func (handler *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var ecommerce *entity.Ecommerce

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ecommerce); err != nil {
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := handler.usecase.Create(ecommerce); err != nil {
		switch err {
		case errors.ErrInvalidPayload:
			render.ResponseError(w, err, http.StatusBadRequest)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, ecommerce, http.StatusCreated)
}
