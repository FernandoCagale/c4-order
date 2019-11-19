package order

import "github.com/FernandoCagale/c4-order/pkg/entity"

type Repository interface {
	Create(customer *entity.Customer) (err error)
	FindAll() (orders []*entity.Customer, err error)
}
