package Reserve

import (
	"fmt"
	"sanvic/User"
	"strconv"
	"time"

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

func CreateReserveService(c *gin.Context) (ReserveModel, error) {
	var r ReserveModel
	c.BindJSON(&r)
	usr, _ := c.Get("my_user_model")
	u, ok := usr.(User.UserModel)
	if !ok {
		fmt.Println("No se ha podido convertir")
	}
	r.Id_user = u.Id
	r.Datefin = r.Dateini.Add(time.Hour * time.Duration(1))
	return CreateReserveRepo(&r, c)
}

func GetHoursService(c *gin.Context) ([]ReserveModel, error) {
	type times struct {
		Dateini time.Time `json:"dateini"`
	}
	var t times
	// test := c.Param("hour")
	c.BindJSON(&t)
	fmt.Println(t.Dateini)
	return GetHoursRepo(t.Dateini, c)
}
