package Thematic

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GET ALL Thematics
func GetAllThematics(c *gin.Context) {
	var thematics []ThematicModel = GetAllThematicsService(c)
	println(thematics)
	serializer := ThematicsSerializer{c, thematics}
	c.JSON(http.StatusOK, serializer.Response())
}

//GET ONE Thematics
func GetThematicByID(c *gin.Context) {
	idThematic := c.Param("id")
	fmt.Println(c.Param("prueba"))
	var thematic ThematicModel
	var id int
	id, err := strconv.Atoi(idThematic)
	if err != nil {
		println("error")
	}

	thematic, err = GetOneThematicService(id, c)

	if err != nil {
		c.JSON(http.StatusNotFound, "Thematic doesn't exist")
	} else {
		serializer := ThematicSerializer{c, thematic}
		c.JSON(http.StatusOK, serializer.Response())
	}
}

//GET Thematics with LIMIT for InfinityScrool
func GetThematicsInfinity(c *gin.Context) {

	offset := c.Query("offset")
	limit := c.Query("limit")
	fmt.Println(offset)
	fmt.Println(limit)
	// var Data_limit string
	// Data_limit = offset + ", " + limit
	Data_limit := []string{offset, limit}
	var thematic []ThematicModel = GetThematicsInfinityService(c, Data_limit)
	c.JSON(http.StatusOK, thematic)
}
