package main

import (
	"net/http"
	"second-api/Config"
	"second-api/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World"))
	})

	r := Routes.SetupRouter()
	//running
	r.Run(":3001")
}
