package order

import (
	"github.com/FernandoCagale/c4-order/pkg/entity"
)

type UseCase interface {
	Create(ecommerce *entity.Ecommerce) (err error)
}
