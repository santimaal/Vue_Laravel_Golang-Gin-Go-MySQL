package Thematic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET ALL Thematics
func GetAllThematics(c *gin.Context) {
	var thematics []ThematicModel = GetAllThematicsService(c)
	println(thematics)
	serializer := ThematicsSerializer{c, thematics}
	c.JSON(http.StatusOK, serializer.Response())
}

// GET ONE Thematics
func GetThematicByID(c *gin.Context) {
	thematic, err := GetOneThematicService(c)
	if err != nil {
		c.JSON(http.StatusNotFound, "Thematic doesn't exist")
	} else {
		serializer := ThematicSerializer{c, thematic}
		c.JSON(http.StatusOK, serializer.Response())
	}
}

// GET Thematics with LIMIT for InfinityScrool
func GetThematicsInfinity(c *gin.Context) {
	thematic, err := GetThematicsInfinityService(c)
	if err != nil {
		c.JSON(http.StatusNotFound, "Thematic doesn't exist")
	} else {
		// serializer := ThematicsSerializer{c, thematic}
		// c.JSON(http.StatusOK, serializer.Response())
		c.JSON(http.StatusOK, thematic)
	}
}
