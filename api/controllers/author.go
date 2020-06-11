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
			"message": "Id not int",
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

// GetAuthorByName on POST /auth
func (AuthorController) GetAuthorByName(c *gin.Context) {
	var request models.Author
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Parsing error",
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
			"message": "Parsing error",
		})
		return
	}

	if dbErr := author.Save(); dbErr != nil {
		c.JSON(dbErr.Code(), gin.H{
			"message": dbErr.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Inserted new author",
	})
}

// UpdateByID on PUT /authors/:id
func (AuthorController) UpdateByID(c *gin.Context) {
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

	dbErr := author.Update()
	if dbErr != nil {
		c.JSON(dbErr.Code(), gin.H{
			"message": dbErr.Error(),
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
			"message": "Id not int",
		})
		return
	}

	author := models.Author{ID: id}

	dbErr := author.Delete()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": dbErr.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Author deleted",
	})
}
