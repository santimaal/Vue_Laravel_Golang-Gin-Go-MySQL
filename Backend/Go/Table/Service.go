package Table

import (
	"github.com/gin-gonic/gin"
)

func GetAllTablesService(c *gin.Context) []TableModel {
	return GetAllTablesRepo(c)
}

func GetOneTableService(id int, c *gin.Context) (TableModel, error) {
	return GetOneTableRepo(id, c)
}

func GetTablesFilterService(c *gin.Context, filter []string) {
	return GetTablesFilterRepo(c, filter)
}
