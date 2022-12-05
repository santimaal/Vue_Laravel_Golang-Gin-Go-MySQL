package Models

import (
	"fmt"
	"second-api/Config"

	_ "github.com/go-sql-driver/mysql"
)

// GetAllTables Fetch all table data
func GetAllTables(table *[]Table) (err error) {
	if err = Config.DB.Find(table).Error; err != nil {
		return err
	}
	return nil
}

func GetTablesFilter(table *[]Table, filter []string) (err error) {
	if err = Config.DB.Raw("SELECT * FROM tables WHERE location LIKE ? AND capacity LIKE ? AND id_thematic LIKE ?", filter[0], filter[1], filter[2]).Scan(table).Error; err != nil {
		return err
	}
	return nil
}

// Createtable ... Insert New data
func CreateTable(table *Table) (err error) {
	if err = Config.DB.Create(table).Error; err != nil {
		return err
	}
	return nil
}

// GettableByID ... Fetch only one table by Id
func GetTableByID(table *Table, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(table).Error; err != nil {
		return err
	}
	return nil
}

// Updatetable ... Update table
func UpdateTable(table *Table, id string) (err error) {
	fmt.Println(table)
	Config.DB.Save(table)
	return nil
}

// Deletetable ... Delete table
func DeleteTable(table *Table, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(table)
	return nil
}
