package controllers

import (
	"net/http"
	"strconv"
	"time"

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

	article, dbErr := models.GetArticleByID(id)
	if dbErr != nil {
		c.JSON(dbErr.Code(), gin.H{
			"message": dbErr.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, article)
}

// GetAPOD on GET /apod
func (ArticleController) GetAPOD(c *gin.Context) {
	articles, err := models.GetArticleByDate(time.Now().Format("01-02-2006"))
	if err != nil {
		c.JSON(err.Code(), gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, articles)
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

	if dbErr := article.Save(); dbErr != nil {
		c.JSON(dbErr.Code(), gin.H{
			"message": dbErr.Error(),
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

	if dbErr := models.DeleteArticleByID(id); err != nil {
		c.JSON(dbErr.Code(), gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
