package order

import (
	"github.com/FernandoCagale/c4-order/internal/errors"
	"github.com/FernandoCagale/c4-order/pkg/entity"
)

type InMemoryRepository struct {
	m map[string]*entity.Customer
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{map[string]*entity.Customer{}}
}

func (repo *InMemoryRepository) FindAll() (orders []*entity.Customer, err error) {
	for _, order := range repo.m {
		orders = append(orders, order)
	}
	return orders, nil
}

func (repo *InMemoryRepository)  FindById(ID string) (order *entity.Customer, err error) {
	for _, order := range repo.m {
		if order.Code == ID {
			return order, nil
		}
	}
	return nil, errors.ErrNotFound
}

func (repo *InMemoryRepository)  DeleteById(ID string) (err error) {
	for _, order := range repo.m {
		if order.Code == ID {
			delete(repo.m, ID)
			return nil
		}
	}
	return errors.ErrNotFound
}

func (repo *InMemoryRepository) Create(e *entity.Customer) (err error) {
	customer := repo.m[e.Code]

	if customer == nil {
		repo.m[e.Code] = e
		return nil
	}

	for _, order := range repo.m[e.Code].Orders {
		if order.Code == e.Orders[0].Code {
			order.Items = e.Orders[0].Items
			return nil
		}
	}
	customer.Orders = append(customer.Orders, e.Orders...)

	return nil
}
