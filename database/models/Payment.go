package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	AmountToPay int `json:"amounttopay"`
	CartId      int `json:"cart_id"`
	Street      string
	HouseNumber string `json:"house_number"`
	PostalCode  string `json:"postal_code"`
	City        string
}
