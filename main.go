package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model `json:"model"`
	Title      string  `json:"title"`
	Artist     string  `json:"artist"`
	Price      float64 `json:"price"`
}

type apiContext struct {
	db *gorm.DB
}

func (ctx *apiContext) getAlbuns(c *gin.Context) {
	var albums []Album
	result := ctx.db.Find(&albums)
	if result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func (ctx *apiContext) getAlbumById(c *gin.Context) {
	var album Album
	id := c.Params.ByName("id")
	result := ctx.db.First(&album, id)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func (ctx *apiContext) postAlbum(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	result := ctx.db.Create(&newAlbum)

	if result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func (ctx *apiContext) deleteAlbum(c *gin.Context) {
	id := c.Params.ByName("id")
	result := ctx.db.Delete(&Album{}, id)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.Status(http.StatusOK)
}

func (ctx *apiContext) updateAlbum(c *gin.Context) {
	var album Album
	var newAlbum Album

	id := c.Params.ByName("id")

	if err := c.BindJSON(&newAlbum); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.db.First(&album, id)

	album.Title = newAlbum.Title
	album.Artist = newAlbum.Artist
	album.Price = newAlbum.Price

	result := ctx.db.Save(&album)

	if result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
	}

	c.IndentedJSON(http.StatusOK, album)
}

func handleRequest(db *gorm.DB) {
	ctx := apiContext{
		db,
	}
	router := gin.Default()
	router.GET("/albuns", ctx.getAlbuns)
	router.GET("/albuns/:id", ctx.getAlbumById)
	router.POST("/albuns", ctx.postAlbum)
	router.DELETE("/albuns/:id", ctx.deleteAlbum)
	router.PUT("/albuns/:id", ctx.updateAlbum)

	router.Run("localhost:9000")
}

func initialMigration(db *gorm.DB) {
	db.AutoMigrate(&Album{})
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
