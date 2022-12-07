package Table

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GET ALL Tables
func GetAllTables(c *gin.Context) {
	var tables []TableModel = GetAllTablesService(c)
	serializer := TablesSerializer{c, tables}
	c.JSON(http.StatusOK, serializer.Response())
}

//GET ONE Tables
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
