package main

import (
	"pin/app/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/revel/revel"
)

func main() {
	db, _ := gorm.Open("mysql", "hoge:hoge@/pin?charset=utf8&parseTime=True")
	db.AutoMigrate(&models.Pin{})
}
