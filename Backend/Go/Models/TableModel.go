package Models

type Table struct {
	Id          int    `json:"id"`
	Is_active   bool   `json:"is_active"`
	Capacity    int    `json:"capacity"`
	Location    string `json:"location"`
	Id_thematic int    `json:"id_thematic"`
}

func (b *Table) TableName() string {
	return "tables"
}
