package User

import (
	"github.com/gin-gonic/gin"
)

func UserRouting(router *gin.RouterGroup) {
	router.GET("/", GetAllUsers)
	router.GET("/:id", GetUserByID)
}
