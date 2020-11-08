package services

import (
	"fmt"
	"github.com/google/shlex"
	"os/exec"
	"rmtly-server/application/interfaces"
	"rmtly-server/application/repositories"
	notificationService "rmtly-server/notification/services"
	"time"
)

func Execute(applicationId string, request interfaces.ExecuteRequest) *interfaces.ExecuteResponse {
	application := repositories.FindById(applicationId)

	c := make(chan bool)
	go func() {
		if request.ExecuteDelay > 0 {
			time.Sleep(time.Duration(request.ExecuteDelay) * time.Millisecond)
		}
		runCommand(application.Exec, c)
	}()
	notificationService.SendAsync(application.Name, "executed by rmtly-server")

	response := new(interfaces.ExecuteResponse)
	response.Application = application
	return response
}

func runCommand(command string, c chan bool) {
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
