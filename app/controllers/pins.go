package controllers

import "github.com/revel/revel"

type Pins struct {
	*revel.Controller
}

func (c Pins) Index() revel.Result {

    revel.TRACE.Printf("%s", c.Params.Get("inputTitle"))
    revel.TRACE.Printf("%s", c.Params.Get("inputMemo"))

	return c.Render()
}

func (c Pins) New() revel.Result {

	return c.Render()
}

func (c Pins) Show() revel.Result {

	return c.Render()
}
