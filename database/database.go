package database

import (
	"fmt"
	"log"

	"example.com/user/hello/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func init() {
	configuration := config.GetConfig()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configuration.DB_USERNAME,
		configuration.DB_PASSWORD,
		configuration.DB_HOST,
		configuration.DB_PORT,
		configuration.DB_NAME)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
	})

	if err != nil {
		log.Println(dsn)
		log.Println(err)
		log.Fatal("database configuration load error.")
	}

	DB = db
	// db.Logger.LogMode(logger.Info)

	/* 	sqldb, _ := db.DB()
	   	sqldb.SetMaxIdleConns(10)
	   	sqldb.SetMaxOpenConns(140) */
}
