package domain

import (
	"erp/api/entity"
	"net/http"

	"github.com/gin-gonic/gin"
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
	result := ctx.db.Find(&costumer)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.IndentedJSON(http.StatusOK, costumer)
}

func handleRequest(db *gorm.DB) {
	ctx := apiContext{
		db,
	}

	router := gin.Default()
	router.GET("/costumer", ctx.getCostumer)
	// router.GET("/costumer/:id", ctx.getAlbumById)
	// router.POST("/costumer", ctx.postAlbum)
	// router.DELETE("/costumer/:id", ctx.deleteAlbum)
	// router.PUT("/costumer/:id", ctx.updateAlbum)

	router.Run("localhost:9000")
}
