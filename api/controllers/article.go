package controllers

import (
	"net/http"
	"strconv"

	"github.com/VIVelev/goApod/models"
	"github.com/gin-gonic/gin"
)

// ArticleController struct
type ArticleController struct{}

// GetAll on GET /articles
func (ArticleController) GetAll(c *gin.Context) {
	articles, err := models.GetAllArticles()

	if err != nil {
		c.JSON(err.Code(), gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, articles)
}

// GetByID on GET /articles/:id
func (ArticleController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})

		return
	}

	article, customErr := models.GetArticleByID(id)
	if customErr != nil {
		c.JSON(customErr.Code(), gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, article)
}

// Create on POST /articles
func (ArticleController) Create(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if customErr := article.Save(); customErr != nil {
		c.JSON(customErr.Code(), gin.H{
			"message": customErr.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, article)
}

// DeleteByID on DELETE /articles/:id
func (ArticleController) DeleteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})

		return
	}

	if customErr := models.DeleteArticleByID(id); err != nil {
		c.JSON(customErr.Code(), gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
