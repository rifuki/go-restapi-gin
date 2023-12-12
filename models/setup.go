package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/db_go_restapi_gin"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	DB = db
}
