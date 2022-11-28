package Models

type Thematic struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Img      string `json:"img"`
}

func (b *Thematic) TableName() string {
	return "thematics"
}
