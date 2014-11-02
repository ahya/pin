package models

import (
    "log"
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
