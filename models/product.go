package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	PriceInCents int    `json:"price_in_cents"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Discount     DiscountModel
}
type DiscountModel struct {
	gorm.Model
	//DiscountId          int     `json:"discount_id"  gorm:"foreignKey"`
	Percentage   float64 `json:"percentage"`
	ValueInCents int     `json:"value_in_cents"`
}
