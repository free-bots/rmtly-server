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
	fmt.Println(applicationEntry)
	fmt.Println(applicationEntry.Exec)
	services.RunCommand(applicationEntry.Exec)
}
