package applicationUtils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

const (
	QUALITY = 512 // image size in pixel
)

func InitIconUtils() {
	gtk.Init(nil)
}

func GetIconBase64(iconName string) *string {
	var icon image.Image

	if strings.Index(iconName, string(os.PathSeparator)) >= 0 {

		fileInfo, err := os.Stat(iconName)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if fileInfo.IsDir() {
			fmt.Println("path is directory")
			return nil
		}

		file, err := os.Open(iconName)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		fileImage, _, err := image.Decode(file)
		if err != nil {
			fmt.Println(err)
		}

		icon = fileImage

	} else {
		icon = GetIcon(iconName)
	}

	if icon == nil {
		return nil
	}

	base64Icon := ImageToBase64(icon)

	return &base64Icon
}

func GetIcon(iconName string) *image.RGBA {

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
