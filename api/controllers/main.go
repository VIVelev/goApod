package controllers

import (
	"net/http"
	"strconv"

	"github.com/VIVelev/goApod/models"
	"github.com/gin-gonic/gin"
)

// MainController of the app
type MainController struct{}

// GetAuthors on GET /authors
func (*MainController) GetAuthors(c *gin.Context) {
	authors, err := models.GetAuthors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database error",
		})
		return
	}
	c.JSON(http.StatusOK, authors)
}

// GetAuthor on GET /author/:id
func (*MainController) GetAuthor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Id not uint",
		})
		return
	}

	author, err := models.GetAuthor(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database error",
		})
		return
	}
	if author == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Invalid id",
		})
		return
	}
	c.JSON(http.StatusOK, author)
}
