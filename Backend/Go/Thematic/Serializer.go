package Thematic

import (
	"github.com/gin-gonic/gin"
)

type ThematicResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Img string `json:"img"`
}

type ThematicSerializer struct {
	C *gin.Context
	thematic ThematicModel
}

type ThematicsSerializer struct {
	C *gin.Context
	thematics []ThematicModel
}

func (ss *ThematicSerializer) Response() ThematicResponse {
	response := ThematicResponse{
		Id:    ss.thematic.Id,
		Name:  ss.thematic.Name,
		Img: ss.thematic.Img,
	}

	return response
}

func (ss *ThematicsSerializer) Response() []ThematicResponse {
	response := []ThematicResponse{}

	for _, thematic := range ss.thematics {
		serializer := ThematicSerializer{ss.C, thematic}
		response = append(response, serializer.Response())
	}

	return response
}
