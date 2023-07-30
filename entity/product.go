package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model  `json:"model"`
	Description string  `json:"description"`
	Reference   string  `json:"reference"`
	Unit        string  `json:"unit"`
	Price       float64 `json:"price"`
	Cost        float64 `json:"cost"`
}
