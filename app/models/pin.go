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

func (p Pin) Find_by_id(id string) Pin {

    var pin Pin
    DB.First(&pin, id)

    return pin
}
