package Reserve

import "time"

type ReserveModel struct {
	Id           uint      `json:"id"`
	Id_table     uint      `gorm:"ForeignKey:Id" json:"id_table"`
	Id_user      uint      `gorm:"ForeignKey:Id" json:"id_user"`
	Is_confirmed string    `json:"is_confirmed"`
	Dateini      time.Time `json:"dateini"`
	Datefin      time.Time `json:"datefin"`
	// Created_at string `json:"created_at"`
	// Updated_at string `json:"updated_at"`
}

func (b *ReserveModel) TableName() string {
	return "reserves"
}
