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

func GetReserveByUserService(c *gin.Context) ([]ReserveModel, error) {
	usr, _ := c.Get("my_user_model")
	u, ok := usr.(User.UserModel)
	if !ok {
		fmt.Println("No se ha podido convertir")
	}
	return GetReserveByUserRepo(u, c)
}

func GetMyReservesService(c *gin.Context) ([]ReserveModel, error) {
	usr, _ := c.Get("my_user_model")
	u, ok := usr.(User.UserModel)
	if !ok {
		fmt.Println("No se ha podido convertir")
	}
	return GetMyReservesRepo(u, c)
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
	r.Is_confirmed = "pending"
	r.Datefin = r.Dateini.Add(time.Hour * time.Duration(1))
	return CreateReserveRepo(&r, c)
}

func GetHoursService(c *gin.Context) ([]ReserveModel, error) {
	// type times struct {
	// 	Dateini time.Time `json:"dateini"`
	// }
	t := c.Param("date")
	id := c.Param("id")

	date, error := time.Parse("2006-01-02", t)
	if error != nil {
		fmt.Println(error)
	}

	// c.BindJSON(&t)
	fmt.Println(date)
	return GetHoursRepo(date, id, c)
}
