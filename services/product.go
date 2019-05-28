package product

import (
	"jbh-product/db"
	"jbh-product/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	var products []models.Product
	db := db.GetDB()
	db.Find(&products)
	c.JSON(200, products)
}

func Create(c *gin.Context) {
	var product models.Product
	var db = db.GetDB()
	if err := c.BindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}
	db.Create(&product)

	if err := db.Create(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &product)
}
