package Table

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTablesService(c *gin.Context) []TableModel {
	return GetAllTablesRepo(c)
}

func GetOneTableService(c *gin.Context) (TableModel, error) {
	idTable := c.Param("id")
	id, err := strconv.Atoi(idTable)
	if err != nil {
		println("error")
	}
	return GetOneTableRepo(id, c)
}

func GetTablesFilterService(c *gin.Context) ([]TableModel, error) {
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

	return GetTablesFilterRepo(c, filter)
}
