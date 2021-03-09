package services

import (
	"bytes"
	"fmt"
	GoppilcationEntry "github.com/free-bots/GopplicationEntry"
	"rmtly-server/application/applicationUtils"
	"rmtly-server/application/interfaces"
	"rmtly-server/application/repositories"
	configService "rmtly-server/config/services"
	"time"
)

// cache
// application
var cachedApplications []*GoppilcationEntry.ApplicationEntry
var lastApplicationCacheRefresh = time.Now()

// icons
var cachedImages map[string]interfaces.IconResponse
var lastIconCacheRefresh = time.Now()

func GetApplications() []*GoppilcationEntry.ApplicationEntry {
	return repositories.FindAll()
}

func GetApplicationById(applicationId string) *GoppilcationEntry.ApplicationEntry {
	return repositories.FindById(applicationId)
}

func GetApplicationsSortedBy(sortKey string) *interfaces.SortedApplicationResponse {
	applications := GetApplications()

	switch sortKey {
	case "category":
		categoryMap := make(map[string][]*GoppilcationEntry.ApplicationEntry)
		for _, application := range applications {
			for _, category := range application.Categories {
				mapValue, exists := categoryMap[category]
				if !exists {
					mapValue = make([]*GoppilcationEntry.ApplicationEntry, 0)
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

func GetApplicationIcon(applicationId string) *bytes.Buffer {
	application := GetApplicationById(applicationId)

	if application == nil {
		return nil
	}

	icon := applicationUtils.GetIconBuffer(application.Icon)
	return icon
}

func applicationCacheExpired() bool {
	cacheDuration := configService.GetConfig().Application.CacheExpiresInMillis

	return cacheExpired(lastApplicationCacheRefresh, cacheDuration)
}

func refreshApplicationCache(applications []*GoppilcationEntry.ApplicationEntry) {
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
