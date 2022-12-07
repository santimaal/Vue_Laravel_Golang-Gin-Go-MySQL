package main

import (
	"fmt"
	"sanvic/Config"
	"sanvic/Routes"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()

	if err != nil {
		fmt.Println("Status:", err)
	}
	// Config.DB.AutoMigrate(&Models.User{})

	r := Routes.SetupRouter()
	r.Run(":3001")
}
