package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"

	"fyne.io/fyne/v2"
)

func (app *Config) pricesTab() *fyne.Container {
	return nil
}

func (app *Config) getChart() *fyne.Canvas {
	return nil
}

func (app *Config) downloadFile(URL, filename string) error {
	// get response bytes from callinga url
	response, err := app.HTTPClient.Get(URL)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New("received wrong response code when downloading file")
	}

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	img, _, err := image.Decode(bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	out, err := os.Create(fmt.Sprintf("./%s", filename))
	if err != nil {
		return err
	}

	err = png.Encode(out, img)
	if err != nil {
		return err
	}

	return nil
}