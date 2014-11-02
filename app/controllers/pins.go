package controllers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"pin/app/models"
	"pin/app/routes"

	"github.com/revel/revel"
)

type Pins struct {
	*revel.Controller
}

func (c Pins) Index() revel.Result {

	revel.TRACE.Printf("%s", c.Params.Get("inputTitle"))
	revel.TRACE.Printf("%s", c.Params.Get("inputMemo"))

    pinView := models.Pin{}.All()

    log.Println(pinView)

	return c.Render(pinView)
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

	log.Println(inputTitle, inputMemo, outImageName)
    pin := models.Pin{Created: time.Now().UnixNano(), Title: inputTitle, Memo: inputMemo, Image: outImageName}
	log.Println(models.DB.NewRecord(pin))
    log.Println(models.DB.Create(&pin))
	log.Println(models.DB.Save(&pin))

	return c.Redirect(routes.Pins.Index())
}

func (c Pins) Show(id string) revel.Result {

    pin := models.Pin{}.Find_by_id(id)

    //TODO: not found
    /**
    if pin == nil {
        return c.NotFound("Pin %d does not exist", id)
    }
    /**/

    log.Println(pin)

	return c.Render(pin)
}

func (c Pins) Favorites() revel.Result {
	return c.Render()
}
