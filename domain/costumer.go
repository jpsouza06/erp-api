package domain

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jpsouza06/erp-api/entity"
	"gorm.io/gorm"
)

type Costumer struct {
	costumer *entity.Costumer
}

type apiContext struct {
	db *gorm.DB
}

func (ctx *apiContext) getCostumer(c *gin.Context) {
	var costumer []Costumer
	result := ctx.db.Find(&albums)
	if result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func handleRequest(db *gorm.DB) {
	ctx := apiContext{
		db,
	}

	router := gin.Default()
	router.GET("/albuns", ctx.getAlbuns)
	// router.GET("/albuns/:id", ctx.getAlbumById)
	// router.POST("/albuns", ctx.postAlbum)
	// router.DELETE("/albuns/:id", ctx.deleteAlbum)
	// router.PUT("/albuns/:id", ctx.updateAlbum)

	router.Run("localhost:9000")
}
