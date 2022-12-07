package Table


type TableModel struct {
	Id        uint   `json:"id"`
	Is_active bool   `json:"is_active"`
	Capacity  uint   `json:"capacity"`
	Location  string `json:"location"`
	Id_thematic uint`gorm:"ForeignKey:Id" json:"id_thematic"`
	// Created_at string `json:"created_at"`
	// Updated_at string `json:"updated_at"`
}
type ThematicModel struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Img  string `json:"img"`
}

func (b *TableModel) TableName() string {
	return "tables"
}
