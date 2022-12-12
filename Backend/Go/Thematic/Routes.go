package Thematic

import (
	"github.com/gin-gonic/gin"
)

func ThematicRouting(router *gin.RouterGroup) {
	router.GET("/", GetAllThematics)
	router.GET("/:id", GetThematicByID)
	// router.GET("/infinite", GetThematicsInfinity)
}