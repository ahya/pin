package controllers

import "github.com/revel/revel"

type Pins struct {
	*revel.Controller
}

type Pin struct {
    Title, Memo string
}

func (c Pins) Index() revel.Result {

    revel.TRACE.Printf("%s", c.Params.Get("inputTitle"))
    revel.TRACE.Printf("%s", c.Params.Get("inputMemo"))

    // TODO: sample
    pinList := []Pin{
        Pin{Title: "たいとる1", Memo: "めもめも1"},
        Pin{Title: "たいとる2", Memo: "めもめも2"},
        Pin{Title: "たいとる3", Memo: "めもめも3"},
    }

	return c.Render(pinList)
}

func (c Pins) New() revel.Result {

	return c.Render()
}

func (c Pins) Show() revel.Result {

    // TODO: IDを取得して表示してみたい
    //var id string = c.Params.Get("id")

	return c.Render()
}
