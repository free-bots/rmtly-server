package main

import (
	"fmt"
	"rmtly-server/applicationUtils"
)

func main() {
	fmt.Println("rmtly-server running...")

	const path = "./test.desktop"
	applicationUtils.Parse(path)
}
