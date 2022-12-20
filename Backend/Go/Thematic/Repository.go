package Thematic

import (
	"fmt"
	"net/http"
	"sanvic/Config"

	"github.com/gin-gonic/gin"
)

func GetAllThematicsRepo(c *gin.Context) []ThematicModel {

	var thematics []ThematicModel

	if err := Config.DB.Find(&thematics).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println("Status:", err)
	}

	return thematics

}

func GetOneThematicRepo(id int, c *gin.Context) (ThematicModel, error) {

	var thematic ThematicModel
	err := Config.DB.Where("id = ?", id).Find(&thematic).Error
	return thematic, err
}

func GetThematicsInfinityRepo(c *gin.Context, Data_limit []string) ([]ThematicModel, error) {
	var thematic []ThematicModel
	err := Config.DB.Limit(Data_limit[1]).Offset(Data_limit[0]).Find(&thematic).Error
	return thematic, err
}
