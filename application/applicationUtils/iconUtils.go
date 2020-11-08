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
	icon := getIcon(iconName)

	if icon == nil {
		return nil
	}

	base64Icon := ImageToBase64(*icon)

	return base64Icon
}

func GetIconBuffer(iconName string) *bytes.Buffer {
	icon := getIcon(iconName)

	if icon == nil {
		return nil
	}

	imageBuffer := new(bytes.Buffer)

	err := png.Encode(imageBuffer, *icon)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return imageBuffer
}

func GetGtkIcon(iconName string) *image.Image {

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

	subImage := iconImage.SubImage(iconImage.Bounds())
	return &subImage
}

func ImageToBase64(image image.Image) *string {

	imageBuffer := bytes.Buffer{}

	err := png.Encode(&imageBuffer, image)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	encodedString := base64.StdEncoding.EncodeToString(imageBuffer.Bytes())

	withType := "data:image/png;base64," + encodedString

	return &withType
}

func isIconNameFilePath(iconName string) bool {
	return strings.Index(iconName, string(os.PathSeparator)) >= 0
}

func getIconFile(iconName string) (*os.File, error) {
	fileInfo, err := os.Stat(iconName)
	if err != nil {
		return nil, err
	}

	if fileInfo.IsDir() {
		return nil, fmt.Errorf("icon path is a directory")
	}

	file, err := os.Open(iconName)
	if err != nil {
		return nil, err
	}

	return file, err
}

func getIcon(iconName string) *image.Image {
	var icon *image.Image

	if isIconNameFilePath(iconName) {

		file, err := getIconFile(iconName)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		defer func() {
			if file != nil {
				_ = file.Close()
			}
		}()

		fileImage, _, err := image.Decode(file)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		icon = &fileImage

	} else {
		icon = GetGtkIcon(iconName)
	}

	return icon
}
