package main

import (
	"github.com/VIVelev/goApod/controllers"
	"github.com/VIVelev/goApod/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	if err := database.Connect(); err != nil {
		panic(err)
	}
	defer database.Connection.Close()
	testController := controllers.TestController{}
	app := gin.Default()
	app.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods",
			"OPTIONS, GET, POST, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Next()
	})
	app.GET("/", testController.Test)
	app.Run(":5000")
}
