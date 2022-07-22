package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	PaymentNumber   int    `json:"paymentnumber"`
	UserId          int    `json:"user_id"`
	StripePaymentId string `json:"stripe_payment_id"`
}
