package database

import (
	"github.com/beintil/zametka-mysql-restapi/zametka"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:password@tcp(localhost:9999)/name_database"))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&zametka.Zametka{})
	
	DB = db
}