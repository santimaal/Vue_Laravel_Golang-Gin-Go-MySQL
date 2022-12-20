package Thematic

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllThematicsService(c *gin.Context) []ThematicModel {
	return GetAllThematicsRepo(c)
}

func GetOneThematicService(c *gin.Context) (ThematicModel, error) {
	idThematic := c.Param("id")
	var id int
	id, err := strconv.Atoi(idThematic)
	if err != nil {
		println("error")
	}

	return GetOneThematicRepo(id, c)
}

func GetThematicsInfinityService(c *gin.Context) ([]ThematicModel, error) {
	offset := c.Query("offset")
	limit := c.Query("limit")
	// var Data_limit string
	// Data_limit = offset + ", " + limit
	Data_limit := []string{offset, limit}
	return GetThematicsInfinityRepo(c, Data_limit)
}
