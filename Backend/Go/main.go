package main

import (
	"second-api/Config"
	"second-api/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()

	r := Routes.SetupRouter()
	//running
	r.Run()
}
