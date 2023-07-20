package main

import (
	"erp/api/domain"
	"erp/api/entity"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func handleRequest(db *gorm.DB) {
	router := gin.Default()
	domain.HandleRequest(db, router)
	router.Run("localhost:9000")
}

func initialMigration(db *gorm.DB) {
	db.AutoMigrate(&entity.Customer{})
}

func main() {
	dsn := "host=0.0.0.0 user=postgres password=postgres dbname=webservice port=5433 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect database")
	}

	initialMigration(db)
	handleRequest(db)
}
