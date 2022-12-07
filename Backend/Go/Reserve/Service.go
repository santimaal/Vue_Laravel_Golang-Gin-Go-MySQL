package Reserve

import (
	"github.com/gin-gonic/gin"
)

func GetAllReservesService(c *gin.Context) []ReserveModel {
	return GetAllReservesRepo(c)
}

func GetOneReserveService(id int, c *gin.Context) (ReserveModel, error) {
	return GetOneReserveRepo(id, c)
}
