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
	// var done ReserveModel
	// c.BindJSON(&done)
	// c.JSON(http.StatusOK, done.Dateini)
	// unaHoraMas := done.Dateini.Add(time.Hour * time.Duration(10))
	// c.JSON(http.StatusOK, unaHoraMas)

	reserve, err := CreateReserveService(c)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, reserve)
	}
	// var err error
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, "Can't reserve it")
	// } else {
	// 	c.JSON(http.StatusOK, "reserve created")
	// }
}
