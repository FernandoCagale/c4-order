package entity

import (
	"github.com/go-ozzo/ozzo-validation"
	"time"
)

type Ecommerce struct {
	ID       string            `json:"id"`
	Customer EcommerceCustomer `json:"customer"`
	Order    EcommerceOrder    `json:"order"`
}

func (e Ecommerce) Validate() error {
	return validation.ValidateStruct(&e,
	)
}

func (e Ecommerce) ToCustomer() Customer {
	items := []*Item{}

	for _, orderItem := range e.Order.Items {
		item := &Item{
			Code:        orderItem.Code,
			Description: orderItem.Description,
			Quantity:    orderItem.Quantity,
			Value:       orderItem.Value,
		}
		items = append(items, item)
	}

	return Customer{
		Code: e.Customer.Code,
		Name: e.Customer.Name,
		Orders: []*Order{
			{
				Code:  e.Order.Code,
				Data:  e.Order.Date,
				Items: items,
			},
		},
	}
}

type EcommerceCustomer struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type EcommerceOrder struct {
	Code  string           `json:"code"`
	Date  time.Time        `json:"date"`
	Items []*EcommerceItem `json:"items"`
}

type EcommerceItem struct {
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Value       float64 `json:"value"`
}
