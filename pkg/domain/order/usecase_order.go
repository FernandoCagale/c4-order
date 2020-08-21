package order

import (
	"github.com/FernandoCagale/c4-order/internal/broker/producer"
	"github.com/FernandoCagale/c4-order/internal/errors"
	"github.com/FernandoCagale/c4-order/pkg/entity"
)

const TOPIC = "order.registered"

type OrderUseCase struct {
	repo     Repository
	producer producer.Producer
}

func NewUseCase(repo Repository, producer producer.Producer) *OrderUseCase {
	return &OrderUseCase{
		repo:     repo,
		producer: producer,
	}
}

func (usecase *OrderUseCase) FindAll() (orders []*entity.Customer, err error) {
	return usecase.repo.FindAll()
}

func (usecase *OrderUseCase) FindById(ID string) (order *entity.Customer, err error) {
	return usecase.repo.FindById(ID)
}

func (usecase *OrderUseCase) DeleteById(ID string) (err error) {
	return usecase.repo.DeleteById(ID)
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

	if err := usecase.producer.Producer(TOPIC, customer); err != nil {
		return err
	}

	return nil
}
