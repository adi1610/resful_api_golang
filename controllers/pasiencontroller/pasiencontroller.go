package pasiencontroller

import (
	"encoding/json"
	"net/http"
	"resful-api-golang/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var pasien []models.Pasien
	models.DB.Find(&pasien)
	c.JSON(http.StatusOK, gin.H{"status": 200, "data": pasien})
}

func Show(c *gin.Context) {
	var pasien models.Pasien
	id := c.Param("id")

	if err := models.DB.First(&pasien, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": 404, "message": "Record not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": 200, "data": pasien})
}

func Create(c *gin.Context) {
	var pasien models.Pasien
	if err := c.ShouldBindJSON(&pasien); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 400, "message": err.Error()})
		return
	}

	models.DB.Create(&pasien)
	c.JSON(http.StatusCreated, gin.H{"status": 201, "data": pasien})

}

func Update(c *gin.Context) {
	var pasien models.Pasien
	id := c.Param("id")

	if err := c.ShouldBindJSON(&pasien); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 400, "message": err.Error()})
		return
	}

	if models.DB.Model(&pasien).Where("id = ?", id).Updates(&pasien).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": 404, "message": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 200, "data": pasien})

}

func Delete(c *gin.Context) {
	var pasien models.Pasien
	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 400, "message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&pasien, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": 404, "message": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "Record deleted!"})

}
