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

// articleRequest struct
type articleRequest struct {
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
	Text     string `json:"text"`
	AuthorID string `json:"authorId"`
	EventID  string `json:"eventId"`
}

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

// GetTopArticles - GET on /top-articles?limit=
func (ArticleController) GetTopArticles(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	articles, dbErr := models.GetTopArticles(limit)
	if dbErr != nil {
		c.JSON(dbErr.Code(), gin.H{
			"message": dbErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, articles)
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
	var articleReq articleRequest
	if err := c.ShouldBindJSON(&articleReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	authorID, err := strconv.Atoi(articleReq.AuthorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	eventID, err := strconv.Atoi(articleReq.EventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	article := models.Article{
		Title:    articleReq.Title,
		ImageURL: articleReq.ImageURL,
		Text:     articleReq.Text,
		AuthorID: authorID,
		EventID:  eventID,
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
