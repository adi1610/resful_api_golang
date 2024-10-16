package main

import (
	"resful-api-golang/controllers/pasiencontroller"
	"resful-api-golang/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/api/pasien", pasiencontroller.Index)
	r.GET("/api/pasien/:id", pasiencontroller.Show)
	r.POST("/api/pasien", pasiencontroller.Create)
	r.PUT("/api/pasien/:id", pasiencontroller.Update)
	r.DELETE("/api/pasien", pasiencontroller.Delete)

	r.Run(":8081")
}
