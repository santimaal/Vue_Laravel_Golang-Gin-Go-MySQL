package Table

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET ALL Tables
func GetAllTables(c *gin.Context) {
	var tables []TableModel = GetAllTablesService(c)
	serializer := TablesSerializer{c, tables}
	c.JSON(http.StatusOK, serializer.Response())
}

func GetTablesFilter(c *gin.Context) {
	table, err := GetTablesFilterService(c)
	if err != nil {
		c.JSON(http.StatusNotFound, "Table doesn't exist")
	} else {
		serializer := TablesSerializer{c, table}
		c.JSON(http.StatusOK, serializer.Response())
	}
}

// GET ONE Tables
func GetTableByID(c *gin.Context) {

	table, err := GetOneTableService(c)
	if err != nil {
		c.JSON(http.StatusNotFound, "Table doesn't exist")
	} else {
		serializer := TableSerializer{c, table}
		c.JSON(http.StatusOK, serializer.Response())
	}

}
