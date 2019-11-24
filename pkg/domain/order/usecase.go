package order

import (
	"github.com/FernandoCagale/c4-order/pkg/entity"
)

type UseCase interface {
	Create(ecommerce *entity.Ecommerce) (err error)
	FindAll() (orders []*entity.Customer, err error)
	FindById(ID string) (order *entity.Customer, err error)
	DeleteById(ID string) (err error)
}
