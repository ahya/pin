package controllers

import (
	"pin/app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/revel/revel"
)

var (
    DB gorm.DB
)

func InitDB() {
    DB, _ = gorm.Open("mysql", "hoge:hoge@/pin?charset=utf8&parseTime=True")

	DB.DB()

	DB.DB().Ping()
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	DB.SingularTable(true)
}

func Migrate() {
    // Pin
	DB.CreateTable(&models.Pin{})
	//DB.DropTable(&models.Pin{})
	DB.DropTableIfExists(&models.Pin{})
	DB.AutoMigrate(&models.Pin{})
}
