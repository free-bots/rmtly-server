package applicationUtils

import (
	"fmt"
	"math"
	"os"
	"rmtly-server/application/interfaces"
)

func GetCurrentTheme() *interfaces.IconTheme {
	return nil
}

func GetHiColorTheme() *interfaces.IconTheme {
	return nil
}

func FindIcon(icon string, size float32, scale string) *string {
	// todo check if icon is file path


	currentTheme :=  GetCurrentTheme()
	if currentTheme != nil {
		fileName := FindIconHelper(icon, size, scale, *currentTheme)
		if fileName != nil {
			return fileName
		}
	}

	hiColorTheme := GetHiColorTheme()
	if hiColorTheme != nil {
		fileName := FindIconHelper(icon, size, scale, *hiColorTheme)
		if fileName != nil {
			return fileName
		}
	}

	return LookupFallbackIcon(icon)
}

func FindIconHelper(icon string, size float32, scale string, theme interfaces.IconTheme) *string {
	fileName := LookUpIcon(icon, size, scale, theme)
	if fileName != nil {
		return fileName
	}

	if true { // todo theme has parrent
		parents := make([]interfaces.IconTheme, 0)
		for _, value := range parents {
			fileName = FindIconHelper(icon, size, scale, value)
			if fileName != nil {
				return fileName
			}
		}
	}

	return nil
}

func LookUpIcon(iconName string, size float32, scale string, theme interfaces.IconTheme) *string {
	for _, themeSubDir := range make([]string, 0) { // todo sub directoryies of theme
		for _, directory := range make([]string, 0) { // todo directory in $(basename list)
			for _, extension := range make([]string, 0) { // extension in ("png", "svg", "xpm")
				if DirectoryMatchesSize(themeSubDir, size, scale) {
					fileName := fmt.Sprintf("%s/%s/%s/iconname.extension", directory, theme, themeSubDir) // todo 					fileName := fmt.Sprintf("%s/%s/%s/iconname.extension", directory, theme, themeSubDir )
					file, _ := os.Open(fileName)
					if file != nil{
						// todo file exists
						return &fileName
					}
				}
			}
		}
	}
	minimal_size := math.MaxFloat32
	for _, subDir := range theme.Directories{
		for each directory in $(basename list) {
			for extension in("png", "svg", "xpm") {
				filename = directory /$(themename) / subdir / iconname.extension
				if exist filename
				and
				DirectorySizeDistance(subdir, size, scale) < minimal_size{
					closest_filename = filename
					minimal_size = DirectorySizeDistance(subdir, size, scale)
				}
			}
		}
	}
	if closest_filename set
	return closest_filename
	return nil
}

func LookupFallbackIcon(iconName string) *string {
	for each directory in $(basename list) {
		for extension in("png", "svg", "xpm") {
			if exists directory / iconname.extension
			return directory / iconname.extension
		}
	}
	return nil
}

func DirectoryMatchesSize(subDir string, iconSize float32, iconScale string) bool {
	read Type and size data from subdir
	if Scale != iconscale
	return false
	if Type is Fixed
	return Size == iconsize
	if Type is Scaled
	return MinSize <= iconsize <= MaxSize
	if Type is Threshold
	return Size-Threshold <= iconsize <= Size+Threshold
}

func DirectorySizeDistance(subDir string, iconSize float32, iconScale string) *string {
	read
	Type
	and
	size
	data
	from
	subdir
	if Type is
	Fixed
	return abs(Size*Scale - iconsize*iconscale)
	if Type is
	Scaled
	if iconsize*iconscale < MinSize*Scale
	return MinSize*Scale - iconsize*iconscale
	if iconsize*iconscale > MaxSize*Scale
	return iconsize*iconscale - MaxSize*Scale
	return 0
	if Type is
	Threshold
	if iconsize*iconscale < (Size-Threshold)*Scale
	return MinSize*Scale - iconsize*iconscale
	if iconsize*iconsize > (Size+Threshold)*Scale
	return iconsize*iconsize - MaxSize*Scale
	return 0
}

func FindBestIcon(iconList []string, size float32, scale string) *string {
	currentTheme := GetCurrentTheme()
	if currentTheme != nil {
		fileName := FindBestIconHelper(iconList, size, scale, *currentTheme)
		if fileName != nil {
			return fileName
		}
	}


	hiColor := GetHiColorTheme()
	if  hiColor != nil {
		fileName := FindBestIconHelper(iconList, size, scale, *hiColor)
		if fileName != nil {
			return fileName
		}
	}

	for _, icon := range iconList {
		fileName := LookupFallbackIcon(icon)
		if fileName != nil {
			return fileName
		}
	}
	return nil
}

func FindBestIconHelper(iconList []string, size float32, scale string, theme interfaces.IconTheme) *string {
	for _, icon := range iconList {
		fileName := LookUpIcon(icon, size, scale, theme)

		if fileName != nil {
			return fileName
		}
	}

	if theme.HasThemesParents() {

		//parents = theme.parents
		parents := make([]interfaces.IconTheme, 0)
		for _, parent := range parents {
			fileName := FindBestIconHelper(iconList, size, scale, parent)
			if fileName != nil {
				return fileName
			}
		}
	}

	return nil
}


