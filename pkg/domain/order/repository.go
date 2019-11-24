package order

import "github.com/FernandoCagale/c4-order/pkg/entity"

type Repository interface {
	Create(customer *entity.Customer) (err error)
	FindAll() (orders []*entity.Customer, err error)
	FindById(ID string) (order *entity.Customer, err error)
	DeleteById(ID string) (err error)
}
