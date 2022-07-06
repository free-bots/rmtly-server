package applicationUtils

import (
	"bytes"
	"fmt"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	QUALITY = 512 // image size in pixel
)

func InitIconUtils() {
	gtk.Init()
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
	icon := getTheme().LookupIcon(
		iconName,
		make([]string, 0),
		QUALITY,
		1,
		gtk.TextDirNone,
		gtk.IconLookupPreload,
	)

	file, err := os.Open(icon.File().Path())
	if err != nil {
		return nil
	}

	defer file.Close()

	if strings.Contains(file.Name(), ".svg") {
		all, err := ioutil.ReadAll(file)
		if err != nil {
			return nil
		}

		toPng, err := svgToPng(all)
		if err != nil {
			return nil
		}

		return &toPng
	}

	decodedImage, _, err := image.Decode(file)
	if err != nil {
		return nil
	}

	return &decodedImage
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

func getTheme() *gtk.IconTheme {
	return gtk.IconThemeGetForDisplay(gdk.DisplayGetDefault())
}

func svgToPng(input []byte) (image.Image, error) {
	cmd := exec.Command("rsvg-convert", "--height", strconv.Itoa(QUALITY), "--width", strconv.Itoa(QUALITY))
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stdin = strings.NewReader(string(input))

	if err := cmd.Run(); err != nil {
		return nil, err
	}
	img, err := png.Decode(&out)
	if err != nil {
		return nil, err
	}
	return img, nil
}
