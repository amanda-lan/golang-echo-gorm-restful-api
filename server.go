package main

import (
	"example.com/user/hello/database"
	"example.com/user/hello/model"
	"example.com/user/hello/route"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = database.DB
}

func main() {

	db.AutoMigrate(&model.Article{})

	e := route.Init()

	e.Logger.Fatal(e.Start(":1323"))

}
