package order

import (
	"github.com/FernandoCagale/c4-order/pkg/entity"
)

type InMemoryRepository struct {
	m map[string]*entity.Customer
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{map[string]*entity.Customer{}}
}

func (repo *InMemoryRepository) Create(e *entity.Customer) (err error) {
	customer := repo.m[e.Code]

	if customer == nil {
		repo.m[e.Code] = e
		return nil
	}
	customer.Orders = append(customer.Orders, e.Orders...)
	return nil
}
