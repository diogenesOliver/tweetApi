package main

import (
	"github.com/diogenesOliver/goAPI/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	routes.AppRoutes(app)
	err := app.Run(":8080")
	if err != nil {
		return
	}
}
