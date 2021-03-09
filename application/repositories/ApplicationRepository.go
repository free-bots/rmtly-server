package repositories

import (
	GoppilcationEntry "github.com/free-bots/GopplicationEntry"
)

func FindAll() []*GoppilcationEntry.ApplicationEntry {
	return GoppilcationEntry.FindAllEntries()
}

func FindById(applicationId string) *GoppilcationEntry.ApplicationEntry {
	applications := FindAll()
	if applications == nil {
		return nil
	}

	for _, application := range applications {
		if application.Id != applicationId {
			continue
		}

		return application
	}

	return nil
}
