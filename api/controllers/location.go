package controllers

import (
	"net/http"
	"strconv"

	"github.com/VIVelev/goApod/models"
	"github.com/gin-gonic/gin"
)

// LocationController struct
type LocationController struct{}

// locationRequest struct
type locationRequest struct {
	Lat  string `json:"lat"`
	Long string `json:"long"`
}

// GetByID on /locations/:id
func (LocationController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})

		return
	}

	location, dbErr := models.GetLocationByID(id)
	if dbErr != nil {
		c.JSON(dbErr.Code(), gin.H{
			"message": dbErr.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, location)
}

// Create on POST /locations
func (LocationController) Create(c *gin.Context) {
	var locReq locationRequest
	if err := c.ShouldBindJSON(&locReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	lat, err := strconv.ParseFloat(locReq.Lat, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	long, err := strconv.ParseFloat(locReq.Long, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	location := models.Location{
		Lat:  lat,
		Long: long,
	}
	if dbErr := location.Save(); dbErr != nil {
		c.JSON(dbErr.Code(), gin.H{
			"message": dbErr.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, location)
}
