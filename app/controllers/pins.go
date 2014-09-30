package controllers

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		pinList = append(pinList, models.Pin{Id: pin.Id, Created: pin.Created, Title: pin.Title, Memo: pin.Memo, Image: pin.Image})
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

	// TODO: アップした画像を書き出す
	var outImageName string
	if c.Params.Files["upImage"] != nil {

		upImage := c.Params.Files["upImage"][0]
		outImageName = fmt.Sprintf("%d", time.Now().UnixNano()) + ".jpg"

		outImage, err := os.Create("./public/uploads/" + outImageName)
		log.Println(outImage)

		if err != nil {
			log.Println(err)
		}
		writer := bufio.NewWriter(outImage)

		image, _ := upImage.Open()
		reader := bufio.NewReader(image)
		bufSize := 4 * 1024 * 1024
		buf := make([]byte, bufSize)
		for {
			n, err := reader.Read(buf)
			if err != nil {
				break
			}
			_, err = writer.Write(buf[:n])
			if err != nil {
				log.Println(err)
				break
			}
		}
		writer.Flush()
	} else {
		log.Println("Image is empty")
	}

	DbMap.Insert(&models.Pin{Created: time.Now().UnixNano(), Title: inputTitle, Memo: inputMemo, Image: outImageName})
	log.Println(inputTitle, inputMemo, outImageName)

	return c.Redirect(routes.Pins.Index())
}

func (c Pins) Show() revel.Result {

	id := c.Params.Route["id"][0]

	return c.Render(id)
}
