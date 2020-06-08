package controllers

import (
	"net/http"

	"github.com/VIVelev/goApod/models"
	"github.com/gin-gonic/gin"
)

// MainController of the app
type MainController struct{}

// GetAuthors on GET /authors
func (*MainController) GetAuthors(c *gin.Context) {
	authors, err := models.GetAuthors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"authors": authors})
}
