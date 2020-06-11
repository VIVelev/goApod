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
	defer database.Db.Close()

	app := gin.Default()

	app.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods",
			"OPTIONS, GET, POST, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Next()
	})

	authorController := controllers.AuthorController{}
	articleController := controllers.ArticleController{}

	registerControllers(app, map[string]interface{}{
		"/authors":  authorController,
		"/articles": articleController,
	})

	app.POST("/auth", authorController.GetAuthorIDByName)
	app.Run(":5000")
}

func registerControllers(app *gin.Engine, pathControllerMap map[string]interface{}) {
	for path, cont := range pathControllerMap {
		// All Getters
		if allGetter, ok := cont.(controllers.AllGetter); ok {
			app.GET(path, allGetter.GetAll)
		}

		// By ID Getters
		if byIDGetter, ok := cont.(controllers.ByIDGetter); ok {
			app.GET(path+"/:id", byIDGetter.GetByID)
		}

		// Entity Creators
		if creator, ok := cont.(controllers.Creator); ok {
			app.POST(path, creator.Create)
		}

		// By ID Updaters
		if byIDUpdater, ok := cont.(controllers.ByIDUpdater); ok {
			app.PUT(path+"/:id", byIDUpdater.UpdateByID)
		}

		// By ID Deleters
		if byIDDeleter, ok := cont.(controllers.ByIDDeleter); ok {
			app.DELETE(path+"/:id", byIDDeleter.DeleteByID)
		}
	}
}
