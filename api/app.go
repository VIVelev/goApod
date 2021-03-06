package main

import (
	"net/http"

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

	if err := database.Prepare(); err != nil {
		panic(err)
	}

	app := gin.Default()

	app.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods",
			"OPTIONS, GET, POST, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Next()
		}
	})

	authorController := controllers.AuthorController{}
	articleController := controllers.ArticleController{}
	likeController := controllers.LikeController{}
	eventController := controllers.EventController{}
	locationController := controllers.LocationController{}

	registerEntityControllers(app, map[string]interface{}{
		"/authors":   authorController,
		"/articles":  articleController,
		"/likes":     likeController,
		"/events":    eventController,
		"/locations": locationController,
	})

	app.POST("/auth", authorController.GetAuthorIDByName)
	app.DELETE("/likes/:authorId/:articleId", likeController.DeleteByAuthorAndArticle)
	app.GET("/apod", articleController.GetAPOD)
	app.GET("/top-articles", articleController.GetTopArticles)
	app.Run(":5000")
}

func registerEntityControllers(app *gin.Engine, pathControllerMap map[string]interface{}) {
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
