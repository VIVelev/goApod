package controllers

import (
	"net/http"
	"strconv"

	"github.com/VIVelev/goApod/models"
	"github.com/gin-gonic/gin"
)

// AuthorController - manages authors
type AuthorController struct{}

// GetAll on GET /authors
func (AuthorController) GetAll(c *gin.Context) {
	authors, err := models.GetAllAuthors()
	if err != nil {
		c.JSON(err.Code(), gin.H{
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

	author, dbErr := models.GetAuthorByID(id)
	if err != nil {
		c.JSON(dbErr.Code(), gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, author)
}

// GetAuthorIDByName on POST /auth - used for authentication
func (AuthorController) GetAuthorIDByName(c *gin.Context) {
	var request models.Author
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	author, dbErr := models.GetAuthorByName(request.Name)
	if dbErr != nil {
		c.JSON(dbErr.Code(), gin.H{
			"message": dbErr.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"id": author.ID})
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

	if dbErr := author.Save(); dbErr != nil {
		c.JSON(dbErr.Code(), gin.H{
			"message": dbErr.Error(),
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
	if dbErr := author.Update(); dbErr != nil {
		c.JSON(dbErr.Code(), gin.H{
			"message": dbErr.Error(),
		})

		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
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

	if dbErr := models.DeleteArticleByID(id); dbErr != nil {
		c.JSON(dbErr.Code(), gin.H{
			"message": dbErr.Error(),
		})

		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
