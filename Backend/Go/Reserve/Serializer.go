package Reserve

import (
	"github.com/gin-gonic/gin"
	// "fmt"
)

type ReserveResponse struct {
	Id           uint `json:"id"`
	Id_table     uint `json:"id_table"`
	Id_user      uint `json:"Id_user"`
	Is_confirmed bool `json:"is_confirmed"`
}

type ReserveSerializer struct {
	C       *gin.Context
	reserve ReserveModel
}

type ReservesSerializer struct {
	C        *gin.Context
	reserves []ReserveModel
}

func (ss *ReserveSerializer) Response() ReserveResponse {
	response := ReserveResponse{
		Id:           ss.reserve.Id,
		Id_table:     ss.reserve.Id_table,
		Id_user:      ss.reserve.Id_user,
		Is_confirmed: ss.reserve.Is_confirmed,
	}

	return response
}

func (ss *ReservesSerializer) Response() []ReserveResponse {
	response := []ReserveResponse{}

	for _, reserve := range ss.reserves {
		serializer := ReserveSerializer{ss.C, reserve}
		response = append(response, serializer.Response())
	}

	return response
}
