package services

import (
	"fmt"
	"github.com/google/shlex"
	"os/exec"
	"rmtly-server/interfaces"
)

const DEFAULT_PATH = "/usr/share/applications"

func GetApplications() []*interfaces.ApplicationEntry {
	return make([]*interfaces.ApplicationEntry, 0)
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
