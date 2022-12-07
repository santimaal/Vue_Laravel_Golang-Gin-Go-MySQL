package Table

import (
	"github.com/gin-gonic/gin"
)

type TableResponse struct {
	Id          uint   `json:"id"`
	Is_active   bool   `json:"is_active"`
	Capacity    uint   `json:"capacity"`
	Location    string `json:"location"`
	Id_thematic uint   `json:"id_thematic"`
}

type TableSerializer struct {
	C     *gin.Context
	table TableModel
}

type TablesSerializer struct {
	C      *gin.Context
	tables []TableModel
}

func (ss *TableSerializer) Response() TableResponse {
	response := TableResponse{
		Id:          ss.table.Id,
		Is_active:   ss.table.Is_active,
		Capacity:    ss.table.Capacity,
		Location:    ss.table.Location,
		Id_thematic: ss.table.Id_thematic,
	}

	return response
}

func (ss *TablesSerializer) Response() []TableResponse {
	response := []TableResponse{}

	for _, table := range ss.tables {
		serializer := TableSerializer{ss.C, table}
		response = append(response, serializer.Response())
	}

	return response
}
