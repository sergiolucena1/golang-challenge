package models

import (
	"gorm.io/gorm"
)


type Product struct {
	gorm.Model
	PriceInCents int    `json:"price_in_cents"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	DiscountID   int
	Discount     DiscountModel
}
type DiscountModel struct {
	gorm.Model
	Percentage   float64 `json:"percentage"`
	ValueInCents int     `json:"value_in_cents"`
}