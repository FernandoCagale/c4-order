package order

import (
	"github.com/FernandoCagale/c4-order/internal/event"
	"github.com/FernandoCagale/c4-order/pkg/entity"
	"github.com/FernandoCagale/c4-order/internal/errors"
)

const TOPIC = "topic.order"

type OrderUseCase struct {
	repo  Repository
	event event.Event
}

func NewUseCase(repo Repository, event event.Event) *OrderUseCase {
	return &OrderUseCase{
		repo:  repo,
		event: event,
	}
}

func (usecase *OrderUseCase) Create(e *entity.Ecommerce) error {
	err := e.Validate()
	if err != nil {
		return errors.ErrInvalidPayload
	}

	customer := e.ToCustomer()

	if err = usecase.repo.Create(&customer); err != nil {
		return err
	}

	return nil
}