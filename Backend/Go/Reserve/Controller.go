package Reserve

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET ALL Reserves
func GetAllReserves(c *gin.Context) {
	var reserves []ReserveModel = GetAllReservesService(c)
	serializer := ReservesSerializer{c, reserves}
	c.JSON(http.StatusOK, serializer.Response())
}

// GET ONE Reserve
func GetReserveByID(c *gin.Context) {

	reserve, err := GetOneReserveService(c)

	if err != nil {
		c.JSON(http.StatusNotFound, "Reserve doesn't exist")
	} else {
		serializer := ReserveSerializer{c, reserve}
		c.JSON(http.StatusOK, serializer.Response())
	}

}

func CreateReserve(c *gin.Context) {
	reserve, err := CreateReserveService(c)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, reserve)
	}
}

func GetHours(c *gin.Context) {
	reserve, err := GetHoursService(c)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, reserve)
	}
}
