package main

import (
	"github.com/gin-gonic/gin"

	"github.com/my-org/controllers"
	"github.com/my-org/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase() // new

	r.GET("/books", controllers.FindBooks) // new

	r.Run()
}
