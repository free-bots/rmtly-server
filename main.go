package main

import (
	"fmt"
	"rmtly-server/applicationUtils"
)

func main() {
	fmt.Println("rmtly-server running...")

	const path = "./test.desktop"
	fmt.Println(applicationUtils.Parse(path))
}
