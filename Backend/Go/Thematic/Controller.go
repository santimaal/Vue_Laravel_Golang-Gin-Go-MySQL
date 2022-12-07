package Thematic

import (
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
