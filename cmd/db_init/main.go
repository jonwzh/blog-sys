package main

import (
	"log"

	"github.com/jonwzh/blog-sys/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	log.Println("DB Init ...")
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	db.DropTableIfExists(model.User{}, model.Post{})
	db.CreateTable(model.User{}, model.Post{})

	users := []model.User{
		{
			Username: "Jonathan",
			PasswordHash: model.GeneratePasswordHash("123456"),
			Posts: []model.Post{
				{Body: "Nice day in Vancouver."},
			},
		},
		{
			Username: "Wang",
			PasswordHash: model.GeneratePasswordHash("123456"),
			Posts: []model.Post{
				{
					Body: "Nice day in Toronto.",
				},
				{
					Body: "Welcome to Canada!",
				},
			},
		},
	}

	for _, u := range(users) {
		db.Debug().Create(&u)
	}
}