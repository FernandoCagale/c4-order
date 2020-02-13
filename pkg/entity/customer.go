package entity

import (
	"time"
)

type Customer struct {
	Code   string        `json:"code" bson:"_id"`
	Name   string        `json:"name"`
	Orders []*Order      `json:"orders"`
}

type Order struct {
	Code  string    `json:"code"`
	Data  time.Time `json:"data"`
	Items []*Item   `json:"items"`
}

type Item struct {
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Value       float64 `json:"value"`
}
