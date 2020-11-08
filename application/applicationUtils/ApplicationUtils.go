package applicationUtils

import (
	"os"
	"strings"
)

const XdgDataDirs = "XDG_DATA_DIRS"
const LocalDir = "~/.local/share/"
const ApplicationDirName = "applications"

func GetApplicationPaths() []string {
	applicationPaths := os.Getenv(XdgDataDirs)
	paths := strings.Split(applicationPaths, ":")

	if strings.Index(applicationPaths, LocalDir) >= 0 {
		paths = append(paths, LocalDir)
	}

	return append(paths)
}
