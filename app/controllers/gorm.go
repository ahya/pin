package controllers

import (
	"pin/app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/revel/revel"
)

func InitDB() {
    models.DB, _ = gorm.Open("mysql", "hoge:hoge@/pin?charset=utf8&parseTime=True")

	models.DB.DB()

	models.DB.DB().Ping()
	models.DB.DB().SetMaxIdleConns(10)
	models.DB.DB().SetMaxOpenConns(100)

	models.DB.SingularTable(true)
}

func Migrate() {
    // Pin
	models.DB.CreateTable(&models.Pin{})
	//models.DB.DropTable(&models.Pin{})
	models.DB.DropTableIfExists(&models.Pin{})
	models.DB.AutoMigrate(&models.Pin{})
}
