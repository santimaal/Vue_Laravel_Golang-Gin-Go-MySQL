package Controllers

import (
	"fmt"
	"net/http"
	"second-api/Models"

	"github.com/gin-gonic/gin"
)

// GetThematics ... Get all Thematics
func GetThematics(c *gin.Context) {
	var thematic []Models.Thematic
	err := Models.GetAllThematics(&thematic)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, thematic)
	}
}

// CreateMesa ... Create Mesa
func CreateThematic(c *gin.Context) {
	var thematic Models.Thematic
	c.BindJSON(&thematic)
	err := Models.CreateThematic(&thematic)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, thematic)
	}
}

// GetMesaByID ... Get the Mesa by id
func GetThematicByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var thematic Models.Thematic
	err := Models.GetThematicByID(&thematic, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, thematic)
	}
}

// UpdateMesa ... Update the Mesa information
func UpdateThematic(c *gin.Context) {
	var thematic Models.Thematic
	id := c.Params.ByName("id")
	err := Models.GetThematicByID(&thematic, id)
	if err != nil {
		c.JSON(http.StatusNotFound, thematic)
	}
	c.BindJSON(&thematic)
	err = Models.UpdateThematic(&thematic, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, thematic)
	}
}

// DeleteMesa ... Delete the Mesa
func DeleteThematic(c *gin.Context) {
	var thematic Models.Thematic
	id := c.Params.ByName("id")
	err := Models.DeleteThematic(&thematic, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id " + id: "is deleted"})
	}
}
