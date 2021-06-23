package main

import (
	"example.com/hello/database"
	"example.com/hello/model"
	"example.com/hello/controller"
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
