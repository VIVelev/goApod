package controllers

import (
	"github.com/VIVelev/goApod/models"
	"github.com/gin-gonic/gin"
)

// TestController test
type TestController struct{}

// Test test
func (T *TestController) Test(c *gin.Context) {
	t := models.Test{c.FullPath()}
	c.JSON(200, gin.H{
		"path": t.GetInfo(),
	})
}
