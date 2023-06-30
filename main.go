package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type album struct {
	gorm.Model `json:"model"`
	Title      string  `json:"title"`
	Artist     string  `json:"artist"`
	Price      float64 `json:"price"`
}

type apiContext struct {
	db *gorm.DB
}

func (self *apiContext) getAlbums(c *gin.Context) {
	var albums []album
	result := self.db.Find(&albums)
	if result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func (self *apiContext) postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	result := self.db.Create(&newAlbum)

	if result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}
	c.IndentedJSON(http.StatusOK, newAlbum)
}

// func Find(id string) *album {
// 	for _, n := range albums {
// 		if id == n.ID {
// 			return &n
// 		}
// 	}
// 	return nil
// }

func handleRequest(db *gorm.DB) {
	ctx := apiContext{
		db,
	}
	router := gin.Default()
	router.GET("/albuns", ctx.getAlbums)
	// router.GET("/albums/:id", getAlbumsById)
	router.POST("/albuns", ctx.postAlbums)

	router.Run("localhost:8080")
}

func initialMigration(db *gorm.DB) {
	db.AutoMigrate(&album{})
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
