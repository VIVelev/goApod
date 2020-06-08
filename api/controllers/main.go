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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Id not int",
		})
		return
	}

	author, err := models.GetAuthor(id)
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

// CreateAuthor on POST /authors
func (*MainController) CreateAuthor(c *gin.Context) {
	var request struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Parsing error",
		})
		return
	}

	author := models.Author{Name: request.Name}
	if err := author.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database error",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Inserted new author",
	})
}

// UpdateAuthor on PUT /authors
func (*MainController) UpdateAuthor(c *gin.Context) {
	var request struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Parsing error",
		})
		return
	}

	author := models.Author{ID: request.ID, Name: request.Name}
	if err := author.Update(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Database error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Author updated",
	})
}
