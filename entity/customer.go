package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model `json:"model"`
	Name       string `json:"name"`
	ZipCode    string `json:"zipCode"`
	City       string `json:"city"`
	Street     string `json:"street"`
	Number     string `json:"number"`
	Phone      string `json:"phone"`
}
