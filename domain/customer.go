package domain

import (
	"erp/api/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApiContext struct {
	db *gorm.DB
}

func (ctx *ApiContext) getCustomer(c *gin.Context) {
	var customer []entity.Customer
	result := ctx.db.Find(&customer)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.IndentedJSON(http.StatusOK, customer)
}

func (ctx *ApiContext) getCustomerById(c *gin.Context) {
	var customer []entity.Customer
	id := c.Params.ByName("id")
	result := ctx.db.Find(&customer, id)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.IndentedJSON(http.StatusOK, customer)
}

func (ctx *ApiContext) postCustomer(c *gin.Context) {
	var customer entity.Customer

	if err := c.BindJSON(&customer); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	result := ctx.db.Create(&customer)

	if result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.IndentedJSON(http.StatusCreated, customer)
}

func (ctx *ApiContext) deleteCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	result := ctx.db.Delete(&entity.Customer{}, id)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.Status(http.StatusOK)
}

func (ctx *ApiContext) updateCustomer(c *gin.Context) {
	var customer entity.Customer
	var newCustomer entity.Customer

	id := c.Params.ByName("id")

	if err := c.BindJSON(&newCustomer); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.db.First(&customer, id)

	newCustomer.ID = customer.ID
	newCustomer.CreatedAt = customer.CreatedAt
	newCustomer.DeletedAt = customer.DeletedAt
	newCustomer.UpdatedAt = customer.UpdatedAt

	result := ctx.db.Save(&newCustomer)

	if result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
	}

	c.IndentedJSON(http.StatusOK, customer)
}

func HandleRequest(db *gorm.DB, router *gin.Engine) {
	ctx := ApiContext{
		db,
	}

	router.GET("/customer", ctx.getCustomer)
	router.GET("/customer/:id", ctx.getCustomerById)
	router.POST("/customer", ctx.postCustomer)
	router.DELETE("/customer/:id", ctx.deleteCustomer)
	router.PUT("/customer/:id", ctx.updateCustomer)

}
