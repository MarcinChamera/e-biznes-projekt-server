package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name       string
	Price      int
	ImageUrl   string `json:"image_url"`
	CategoryId int    `json:"category_id"`
}
