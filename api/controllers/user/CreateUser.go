package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishwaszadte/event-magnet-be-core/api/models"
	"github.com/vishwaszadte/event-magnet-be-core/configs"
)

// CreateUser creates a user using the data provided in the JSON request body.
//
// c *gin.Context
// Return type: None
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := configs.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
