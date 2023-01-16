package Reserve

import (
	"fmt"
	"net/http"
	"sanvic/Config"
	"sanvic/User"
	"time"

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

func GetReserveByUserRepo(u User.UserModel, c *gin.Context) (reserve []ReserveModel, err error) {
	// err = Config.DB.Where("id_user = ? AND is_confirmed <> 'pending'", u.Id).Order("updated_at desc").Limit(u.Noti).Find(&reserve).Error
	err = Config.DB.Raw("CALL getNotificationsClient(?,?)", u.Id, u.Noti).Find(&reserve).Error
	return reserve, err
}

func GetMyReservesRepo(u User.UserModel, c *gin.Context) (reserve []ReserveModel, err error) {
	err = Config.DB.Where("id_user = ?", u.Id).Order("updated_at desc").Find(&reserve).Error
	return reserve, err
}

func CreateReserveRepo(r *ReserveModel, c *gin.Context) (ReserveModel, error) {
	err := Config.DB.Create(r).Error
	return *r, err
}

func GetHoursRepo(t time.Time, id string, c *gin.Context) (r []ReserveModel, err error) {
	err = Config.DB.Where("id_table = ? AND dateini BETWEEN ? AND ?", id, t, t.AddDate(0, 0, 1)).Find(&r).Error
	return r, err
}
