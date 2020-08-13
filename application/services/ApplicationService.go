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
	configService "rmtly-server/config/services"
	"strings"
	"time"
)

const XdgDataDirs = "XDG_DATA_DIRS"
const LocalDir = "~/.local/share/"
const ApplicationDirName = "applications"

// cache
// application
var cachedApplications []*interfaces.ApplicationEntry
var lastApplicationCacheRefresh time.Time = time.Now()

// icons
var cachedImages map[string]interfaces.IconResponse
var lastIconCacheRefresh = time.Now()

func GetApplications() []*interfaces.ApplicationEntry {

	if !applicationCacheExpired() && cachedApplications != nil {
		return cachedApplications
	}

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

	refreshApplicationCache(applications)

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

	if !iconCacheExpired() && cachedImages != nil && iconCacheContains(applicationId) {
		value, _ := cachedImages[applicationId]
		return &value
	}

	response := new(interfaces.IconResponse)
	response.ApplicationId = applicationId

	application := GetApplicationById(applicationId)

	if application == nil {
		return nil
	}

	icon := applicationUtils.GetIconBase64(application.Icon)

	if icon == nil {
		return response
	}

	response.IconBase64 = *icon

	refreshIconCache(*response)

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

func applicationCacheExpired() bool {
	cacheDuration := configService.GetConfig().Application.CacheExpiresInMillis

	return cacheExpired(lastApplicationCacheRefresh, cacheDuration)
}

func refreshApplicationCache(applications []*interfaces.ApplicationEntry) {
	cachedApplications = applications
	lastApplicationCacheRefresh = time.Now()
}

func iconCacheExpired() bool {
	cacheDuration := configService.GetConfig().Image.CacheExpiresInMillis

	return cacheExpired(lastIconCacheRefresh, cacheDuration)
}

func refreshIconCache(icon interfaces.IconResponse) {

	if cachedImages == nil {
		cachedImages = make(map[string]interfaces.IconResponse)
	}

	if len(cachedImages) > configService.GetConfig().Image.MaxImagesInCache {
		cachedImages = nil
		cachedImages = make(map[string]interfaces.IconResponse)
	}

	cachedImages[icon.ApplicationId] = icon
	lastIconCacheRefresh = time.Now()
}

func iconCacheContains(applicationId string) bool {
	_, exists := cachedImages[applicationId]
	return exists
}

func cacheExpired(currentCacheTime time.Time, configDifference int) bool {
	timeDuration := time.Now().Sub(currentCacheTime).Milliseconds()

	return timeDuration > int64(configDifference)
}
