package controller

import (
	"go/api/go-start/infrastructure"
	"go/api/go-start/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBind(&product); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	data := infrastructure.DB.Create(&product)

	if data.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Khong tao duoc product"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": product})
}
func GetProductById(c *gin.Context) {
	var data models.Product
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Id khong hop le"})
		return
	}
	if err := infrastructure.DB.Where("id=?", id).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "San phan khong hop le"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})

}
func GetAllProducts(c *gin.Context) {
	var data []models.Product
	if err := infrastructure.DB.Find(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Not get all products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
