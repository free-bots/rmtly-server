package main

import (
	"fmt"
	"rmtly-server/applicationUtils"
)

func main() {
	fmt.Println("rmtly-server running...")

	const path = "./test.desktop"
	applicationEntry := applicationUtils.Parse(path, true)
	fmt.Println(applicationEntry)
	//err := exec.Command(applicationEntry.Exec).Run()
	//fmt.Println(err)
}
