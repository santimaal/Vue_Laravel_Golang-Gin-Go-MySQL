package Models

import (
	"fmt"
	"second-api/Config"

	_ "github.com/go-sql-driver/mysql"
)

// GetAllThematics Fetch all Thematic data
func GetAllThematics(thematic *[]Thematic) (err error) {
	if err = Config.DB.Find(thematic).Error; err != nil {
		return err
	}
	return nil
}

// CreateThematic ... Insert New data
func CreateThematic(thematic *Thematic) (err error) {
	if err = Config.DB.Create(thematic).Error; err != nil {
		return err
	}
	return nil
}

// GetThematicByID ... Fetch only one Thematic by Id
func GetThematicByID(thematic *Thematic, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(thematic).Error; err != nil {
		return err
	}
	return nil
}

// UpdateThematic ... Update Thematic
func UpdateThematic(thematic *Thematic, id string) (err error) {
	fmt.Println(thematic)
	Config.DB.Save(thematic)
	return nil
}

// DeleteThematic ... Delete Thematic
func DeleteThematic(thematic *Thematic, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(thematic)
	return nil
}
