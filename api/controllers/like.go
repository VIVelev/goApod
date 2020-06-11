package controllers

import (
	"net/http"
	"strconv"

	"github.com/VIVelev/goApod/models"
	"github.com/gin-gonic/gin"
)

// LikeController - controls likes
type LikeController struct{}

// Create on POST /likes
func (LikeController) Create(c *gin.Context) {
	var like models.Like
	if err := c.ShouldBindJSON(&like); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if dbErr := like.Save(); dbErr != nil {
		c.JSON(dbErr.Code(), gin.H{
			"message": dbErr.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Added like",
	})
}

// DeleteByAuthorAndArticle on DELETE /likes/:authorId/:articleId
func (LikeController) DeleteByAuthorAndArticle(c *gin.Context) {
	authorID, err := strconv.Atoi(c.Param("authorId"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})

		return
	}
	articleID, err := strconv.Atoi(c.Param("articleId"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})

		return
	}

	like := models.Like{AuthorID: authorID, ArticleID: articleID}
	if dbErr := like.DeleteByAuthorAndArticle(); dbErr != nil {
		c.JSON(dbErr.Code(), gin.H{
			"message": dbErr.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Entity deleted",
	})
}
