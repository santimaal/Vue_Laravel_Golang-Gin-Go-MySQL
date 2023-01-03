package Reserve

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllReservesService(c *gin.Context) []ReserveModel {
	return GetAllReservesRepo(c)
}

func GetOneReserveService(c *gin.Context) (ReserveModel, error) {
	idReserve := c.Param("id")
	id, err := strconv.Atoi(idReserve)
	if err != nil {
		println("error")
	}
	return GetOneReserveRepo(id, c)
}
