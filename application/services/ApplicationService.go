package services

import (
	"fmt"
	"github.com/google/shlex"
	"os"
	"os/exec"
	"path/filepath"
	"rmtly-server/application/applicationUtils"
	application2 "rmtly-server/application/applicationUtils/parser/application"
	"rmtly-server/application/interfaces"
	"strings"
)

const XdgDataDirs = "XDG_DATA_DIRS"
const LocalDir = "~/.local/share/"
const ApplicationDirName = "applications"

func GetApplications() []*interfaces.ApplicationEntry {
	applications := make([]*interfaces.ApplicationEntry, 0)

	for _, path := range getApplicationPaths() {

		path = filepath.Join(path, ApplicationDirName)

		fileInfo, err := os.Stat(path)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if !fileInfo.IsDir() {
			continue
		}

		file, err := os.Open(path)

		if err != nil {
			return nil
		}

		fileNames, err := file.Readdirnames(0)

		if err != nil {
			return nil
		}

		for _, name := range fileNames {
			application := application2.Parse(filepath.Join(path, name), true)
			if application == nil {
				continue
			}
			applications = append(applications, application)
		}

	}

	return applications
}

func GetApplicationById(applicationId string) *interfaces.ApplicationEntry {
	applications := GetApplications()
	if applications == nil {
		return nil
	}

	for _, application := range applications {
		if application == nil {
			continue
		}

		if application.Id == applicationId {
			return application
		}
	}

	return nil
}

func GetApplicationsSortedBy(sortKey string) *interfaces.SortedApplicationResponse {
	applications := GetApplications()

	switch sortKey {
	case "category":
		categoryMap := make(map[string][]*interfaces.ApplicationEntry)
		for _, application := range applications {
			for _, category := range application.Categories {
				mapValue, exists := categoryMap[category]
				if !exists {
					mapValue = make([]*interfaces.ApplicationEntry, 0)
				}
				categoryMap[category] = append(mapValue, application)
			}
		}

		sortedResponse := new(interfaces.SortedApplicationResponse)
		sortedResponse.SortedBy = sortKey
		for key, value := range categoryMap {
			sortedItem := new(interfaces.SortedValue)
			sortedItem.SortedValue = key
			sortedItem.ApplicationEntries = value

			sortedResponse.Values = append(sortedResponse.Values, *sortedItem)
		}

		return sortedResponse
	default:
		fmt.Println("key not found")
		return nil
	}
}

func RunCommand(command string, c chan bool) {
	args, err := shlex.Split(command)
	if err != nil {
		fmt.Println(err)
		return
	}

	if args == nil || len(args) == 0 {
		c <- false
		return
	}

	if len(args) > 1 {
		err = exec.Command(args[0], args[1:]...).Run()
	} else if len(args) == 1 {
		err = exec.Command(args[0]).Run()
	}

	if err != nil {
		fmt.Println(err)
		c <- false
		return
	}

	c <- true
}

func GetIconOfApplication(applicationId string) *interfaces.IconResponse {

	response := new(interfaces.IconResponse)
	response.ApplicationId = applicationId

	application := GetApplicationById(applicationId)

	if application == nil {
		return response
	}

	icon := applicationUtils.GetIconBase64(application.Icon)

	if icon == nil {
		return response
	}

	response.IconBase64 = *icon
	return response
}

func getApplicationPaths() []string {
	applicationPaths := os.Getenv(XdgDataDirs)
	paths := strings.Split(applicationPaths, ":")

	if strings.Index(applicationPaths, LocalDir) >= 0 {
		paths = append(paths, LocalDir)
	}

	return append(paths)
}
