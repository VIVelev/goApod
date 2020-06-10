package controllers

import (
	"net/http"
	"strconv"

	"github.com/VIVelev/goApod/models"
	"github.com/gin-gonic/gin"
)

// AuthorController of the app
type AuthorController struct{}

// GetAll on GET /authors
func (AuthorController) GetAll(c *gin.Context) {
	authors, err := models.GetAllAuthors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, authors)
}

// GetByID on GET /authors/:id
func (AuthorController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})

		return
	}

	author, err := models.GetAuthorByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, author)
}

// Create on POST /authors
func (AuthorController) Create(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := author.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, author)
}

// UpdateByID on PUT /authors/:id
func (AuthorController) UpdateByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})

		return
	}

	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	author.ID = id
	if err := author.Update(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Author updated",
	})
}

// DeleteByID on DELETE /authors/:id
func (AuthorController) DeleteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := models.DeleteAuthorByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
