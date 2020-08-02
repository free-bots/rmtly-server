package services

import (
	"fmt"
	"github.com/google/shlex"
	"os"
	"os/exec"
	"rmtly-server/application/applicationUtils"
	application2 "rmtly-server/application/applicationUtils/parser/application"
	"rmtly-server/application/interfaces"
)

const DefaultPath = "/usr/share/applications"

func GetApplications() []*interfaces.ApplicationEntry {
	applications := make([]*interfaces.ApplicationEntry, 0)

	file, err := os.Open(DefaultPath)

	if err != nil {
		return nil
	}

	fileNames, err := file.Readdirnames(0)

	if err != nil {
		return nil
	}

	for _, name := range fileNames {
		application := application2.Parse(DefaultPath+string(os.PathSeparator)+name, true)
		applications = append(applications, application)
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

	// todo check if application is path

	icon := applicationUtils.GetIcon(application.Icon)

	if icon == nil {
		return response
	}

	response.IconBase64 = applicationUtils.ImageToBase64(icon)
	return response
}
