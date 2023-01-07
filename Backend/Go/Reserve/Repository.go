package Reserve

import (
	"fmt"
	"net/http"
	"sanvic/Config"

	"github.com/gin-gonic/gin"
)

func GetAllReservesRepo(c *gin.Context) []ReserveModel {

	var reserves []ReserveModel

	if err := Config.DB.Find(&reserves).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println("Status:", err)
	}

	return reserves

}

func GetOneReserveRepo(id int, c *gin.Context) (ReserveModel, error) {
	var reserve ReserveModel
	err := Config.DB.Where("id = ?", id).Find(&reserve).Error
	return reserve, err
}

func CreateReserveRepo(r *ReserveModel, c *gin.Context) (ReserveModel, error) {
	err := Config.DB.Create(r).Error
	return *r, err
}
