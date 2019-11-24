package main

import (
	"net/http"

	"github.com/jonwzh/blog-sys/model"
	"github.com/jonwzh/blog-sys/controller"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Use Mysql dialect
)

func main() {
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	controller.StartUp()
	http.ListenAndServe(":8080", nil)
}