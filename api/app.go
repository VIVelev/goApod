package main

import (
	"github.com/VIVelev/goApod/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	app := gin.Default()
	app.GET("/", controllers.Test)
	app.Run(":5000")
}
