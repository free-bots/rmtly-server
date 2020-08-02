package applicationUtils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"image"
	"image/color"
	"image/png"
)

const (
	QUALITY = 512 // image size in pixel
)

func GetIcon(iconName string) *image.RGBA {
	gtk.Init(nil)

	theme, err := gtk.IconThemeGetDefault()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	buff, err := theme.LoadIcon(iconName, QUALITY, 0)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	pixels := buff.GetPixels()

	iconImage := image.NewRGBA(image.Rect(0, 0, buff.GetHeight(), buff.GetWidth()))

	x := 0
	y := 0

	for index := 0; index <= len(pixels)-buff.GetNChannels(); index += buff.GetNChannels() {
		r := pixels[index]
		g := pixels[index+1]
		b := pixels[index+2]
		a := pixels[index+3]

		iconImage.Set(x, y, color.RGBA{R: r, G: g, B: b, A: a})

		x++

		if x >= buff.GetWidth() {
			x = 0
			y++
		}
	}

	return iconImage
}

func ImageToBase64(image image.Image) string {
	imageBuffer := bytes.Buffer{}

	err := png.Encode(&imageBuffer, image)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	encodedString := base64.StdEncoding.EncodeToString(imageBuffer.Bytes())

	return "data:image/png;base64," + encodedString
}
