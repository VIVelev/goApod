package controllers

import (
	"net/http"
	"strconv"

	"github.com/VIVelev/goApod/models"
	"github.com/gin-gonic/gin"
)

// AuthorController - manages authors
type AuthorController struct{}

// GetAuthors on GET /authors
func (*AuthorController) GetAuthors(c *gin.Context) {
	authors, err := models.GetAllAuthors()
	if err != nil {
		c.JSON(err.Code(), gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, authors)
}

// GetAuthor on GET /authors/:id
func (*AuthorController) GetAuthor(c *gin.Context) {
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
func (*AuthorController) GetAuthorByName(c *gin.Context) {
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

// CreateAuthor on POST /authors
func (*AuthorController) CreateAuthor(c *gin.Context) {
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

// UpdateAuthor on PUT /authors/:id
func (*AuthorController) UpdateAuthor(c *gin.Context) {
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

// DeleteAuthor on DELETE /authors/:id
func (*AuthorController) DeleteAuthor(c *gin.Context) {
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
