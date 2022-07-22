package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model

	CartNumber int `json:"cart_number"`
	ProductId  int `json:"productid"`
	Quantity   int
}
