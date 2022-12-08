package Table

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET ALL Tables
func GetAllTables(c *gin.Context) {
	var tables []TableModel = GetAllTablesService(c)
	serializer := TablesSerializer{c, tables}
	c.JSON(http.StatusOK, serializer.Response())
}

func GetTablesFilter(c *gin.Context) {

	location := c.Query("location")
	thematic := c.Query("id_thematic")
	capacity := c.Query("capacity")
	search := c.Query("search")

	if len(location) < 1 {
		location = "%"
	}
	if len(thematic) < 1 {
		thematic = "%"
	}
	if len(capacity) < 1 {
		capacity = "%"
	}
	if len(search) < 1 {
		search = "%"
	} else {
		search = "%" + search + "%"
	}
	filter := []string{location, capacity, thematic, search}

	var table []TableModel = GetTablesFilterService(c, filter)
	c.JSON(http.StatusOK, table)
}

// GET ONE Tables
func GetTableByID(c *gin.Context) {
	idTable := c.Param("id")
	var table TableModel
	var id int
	id, err := strconv.Atoi(idTable)
	if err != nil {
		println("error")
	}

	table, err = GetOneTableService(id, c)

	if err != nil {
		c.JSON(http.StatusNotFound, "Table doesn't exist")
	} else {
		serializer := TableSerializer{c, table}
		c.JSON(http.StatusOK, serializer.Response())
	}

}
