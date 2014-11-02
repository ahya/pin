package models

type Pin struct {
	Id, Created        int64
	Title, Memo, Image string
}

func (p Pin) All() PinView {

    var pins []Pin
    DB.Find(&pins)

    return PinView{Pins: pins}
}
