package controllers

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"
	"github.com/zuiurs/revel-sample-shop/app/models"
	"log"
)

var DB *gorm.DB

func InitDB() {
	driverName, ok := revel.Config.String("db.driver")
	if !ok {
		log.Panicln("db.driver is not set")
	}

	dsn, ok := revel.Config.String("db.spec")
	if !ok {
		log.Panicln("db.spec is not defined")
	}

	db, err := gorm.Open(driverName, dsn)
	if err != nil {
		log.Panicln("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.Merchandise{}, &models.User{}, &models.Role{})
	DB = db
}
