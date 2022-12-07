package Thematic

type ThematicModel struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Img string `json:"img"`
	// Created_at string `json:"created_at"`
	// Updated_at string `json:"updated_at"`
}

func (b *ThematicModel) TableName() string {
	return "thematics"
}
