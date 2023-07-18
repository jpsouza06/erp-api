package main

import (
	"erp/api/entity"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type apiContext struct {
	db *gorm.DB
}

// func (ctx *apiContext) getAlbuns(c *gin.Context) {
// 	var albums []Album
// 	result := ctx.db.Find(&albums)
// 	if result.Error != nil {
// 		c.AbortWithError(http.StatusInternalServerError, result.Error)
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// func (ctx *apiContext) getAlbumById(c *gin.Context) {
// 	var album Album
// 	id := c.Params.ByName("id")
// 	result := ctx.db.First(&album, id)

// 	if result.Error != nil {
// 		c.AbortWithError(http.StatusNotFound, result.Error)
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, album)
// }

// func (ctx *apiContext) postAlbum(c *gin.Context) {
// 	var newAlbum Album

// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	result := ctx.db.Create(&newAlbum)

// 	if result.Error != nil {
// 		c.AbortWithError(http.StatusInternalServerError, result.Error)
// 		return
// 	}
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }

// func (ctx *apiContext) deleteAlbum(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	result := ctx.db.Delete(&Album{}, id)

// 	if result.Error != nil {
// 		c.AbortWithError(http.StatusNotFound, result.Error)
// 		return
// 	}
// 	c.Status(http.StatusOK)
// }

// func (ctx *apiContext) updateAlbum(c *gin.Context) {
// 	var album Album
// 	var newAlbum Album

// 	id := c.Params.ByName("id")

// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	ctx.db.First(&album, id)

// 	album.Title = newAlbum.Title
// 	album.Artist = newAlbum.Artist
// 	album.Price = newAlbum.Price

// 	result := ctx.db.Save(&album)

// 	if result.Error != nil {
// 		c.AbortWithError(http.StatusInternalServerError, result.Error)
// 	}

// 	c.IndentedJSON(http.StatusOK, album)
// }

func initialMigration(db *gorm.DB) {
	db.AutoMigrate(&entity.Costumer{})
}

func main() {
	dsn := "host=0.0.0.0 user=postgres password=postgres dbname=webservice port=5433 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect database")
	}

	initialMigration(db)
	domain.handleRequest
}
