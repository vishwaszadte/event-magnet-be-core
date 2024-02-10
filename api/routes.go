package api

import (
	"github.com/gin-gonic/gin"
	userController "github.com/vishwaszadte/event-magnet-be-core/api/controllers/user"
)

func RegisterRoutes(r *gin.Engine) {

	r.POST("/users", userController.CreateUser)
}
