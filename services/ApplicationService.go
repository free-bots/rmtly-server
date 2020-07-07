package services

import (
	"fmt"
	"github.com/google/shlex"
	"os"
	"os/exec"
	"rmtly-server/applicationUtils"
	"rmtly-server/interfaces"
)

const DEFAULT_PATH = "/usr/share/applications"

func GetApplications() []*interfaces.ApplicationEntry {
	applications := make([]*interfaces.ApplicationEntry, 0)

	file, err := os.Open(DEFAULT_PATH)

	if err != nil {
		return nil
	}

	fileNames, err := file.Readdirnames(0)

	if err != nil {
		return nil
	}

	for _, name := range fileNames {
		application := applicationUtils.Parse(DEFAULT_PATH+string(os.PathSeparator)+name, true)
		applications = append(applications, application)
	}

	return applications
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
