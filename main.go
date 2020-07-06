package main

import (
	"fmt"
	"rmtly-server/applicationUtils"
	"rmtly-server/services"
)

func main() {
	fmt.Println("rmtly-server running...")

	const path = "./test.desktop"
	applicationEntry := applicationUtils.Parse(path, true)

	c := make(chan bool)
	go services.RunCommand(applicationEntry.Exec, c)

	fmt.Printf("running %s succesful %t", applicationEntry.Exec, <-c)
}
