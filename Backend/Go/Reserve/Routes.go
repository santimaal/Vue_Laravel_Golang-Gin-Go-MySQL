package Reserve

import (
	"github.com/gin-gonic/gin"
)

func ReserveRouting(router *gin.RouterGroup) {
	router.GET("/", GetAllReserves)
	router.GET("/:id", GetReserveByID)
}