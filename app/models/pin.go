package models

import (

    "log"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var (
    DB gorm.DB
)

type Pin struct {
	Id, Created        int64
	Title, Memo, Image string
}

func (m Pin) All() PinList {

    var pins []Pin
    DB.Find(&pins)

    pinList := PinList{Pins: pins}

    log.Println(pins)
    log.Println(pinList)

    return pinList
}
