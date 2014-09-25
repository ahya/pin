package controllers

import (
	"fmt"
	"log"
	"time"

	"pin/app/models"
	"pin/app/routes"
	"pin/app/viewmodels"

	"github.com/revel/revel"
)

type Pins struct {
	*revel.Controller
}

func (c Pins) Index() revel.Result {

	revel.TRACE.Printf("%s", c.Params.Get("inputTitle"))
	revel.TRACE.Printf("%s", c.Params.Get("inputMemo"))

	var pinList []models.Pin

	rows, _ := DbMap.Select(models.Pin{}, "select * from pin")
	for _, row := range rows {
		pin := row.(*models.Pin)

		// TODO: よくわからなかったのでmoels.Pinにbindingしなおしている
		pinList = append(pinList, models.Pin{Id: pin.Id, Created: pin.Created, Title: pin.Title, Memo: pin.Memo})
	}

	pinListViewModel := &viewModels.PinList{pinList}

	fmt.Println(pinListViewModel)

	return c.Render(pinListViewModel)
}

func (c Pins) New() revel.Result {

	return c.Render()
}

func (c Pins) Post(inputTitle string, inputMemo string) revel.Result {

	c.Validation.Required(inputTitle)
	c.Validation.Required(inputMemo)
	c.Validation.MaxSize(inputTitle, 20)
	c.Validation.MaxSize(inputMemo, 140)

	if c.Validation.HasErrors() {

		c.Validation.Keep()
		c.FlashParams()

		return c.Redirect(routes.Pins.New())
	}

	DbMap.Insert(&models.Pin{Created: time.Now().UnixNano(), Title: inputTitle, Memo: inputMemo})
	log.Println(inputTitle, inputMemo)

	return c.Redirect(routes.Pins.Index())
}

func (c Pins) Show() revel.Result {

	id := c.Params.Route["id"][0]

	return c.Render(id)
}
