package Table

import (
	"fmt"
	"net/http"
	"sanvic/Config"

	"github.com/gin-gonic/gin"
)

func GetAllTablesRepo(c *gin.Context) []TableModel {

	var tables []TableModel

	if err := Config.DB.Find(&tables).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println("Status:", err)
	}

	return tables

}

func GetTablesFilterRepo(c *gin.Context filter []string)  []TableModel {
	var table []TableModel

	if err = Config.DB.Raw("SELECT * FROM tables WHERE location LIKE ? AND capacity LIKE ? AND id_thematic LIKE ?", filter[0], filter[1], filter[2]).Scan(table).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println("Status:", err)	
	}
	return table
}

func GetOneTableRepo(id int, c *gin.Context) (TableModel, error) {

	var table TableModel

	err := Config.DB.Where("id = ?", id).Find(&table).Error
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	return table, err
}
