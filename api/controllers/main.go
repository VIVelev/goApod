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
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Parsing error",
		})
		return
	}

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

// UpdateAuthor on PUT /authors/:id
func (*MainController) UpdateAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Id not int",
		})
		return
	}

	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Parsing error",
		})
		return
	}

	author.ID = id

	isUpdated, err := author.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Database error",
		})
		return
	}
	if !isUpdated {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid id",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Author updated",
	})
}

// DeleteAuthor on DELETE /authors/:id
func (*MainController) DeleteAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Id not int",
		})
		return
	}

	author := models.Author{ID: id}

	isDeleted, err := author.Delete()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Database error",
		})
		return
	}
	if !isDeleted {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid id",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Author deleted",
	})
}
