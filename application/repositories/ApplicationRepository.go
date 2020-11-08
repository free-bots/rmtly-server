package repositories

import (
	"fmt"
	"os"
	"path/filepath"
	"rmtly-server/application/applicationUtils"
	applicationParser "rmtly-server/application/applicationUtils/parser/application"
	"rmtly-server/application/interfaces"
)

func FindAll() []*interfaces.ApplicationEntry {
	applications := make([]*interfaces.ApplicationEntry, 0)
	for _, path := range applicationUtils.GetApplicationPaths() {

		path = filepath.Join(path, applicationUtils.ApplicationDirName)

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

		_ = file.Close()

		for _, name := range fileNames {
			application := applicationParser.Parse(filepath.Join(path, name), true)
			if application == nil {
				continue
			}

			applications = append(applications, application)
		}
	}

	return applications
}

func FindById(applicationId string) *interfaces.ApplicationEntry {
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
