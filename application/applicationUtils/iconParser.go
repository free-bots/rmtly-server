package applicationUtils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"rmtly-server/application/interfaces"
)

func ParseIconTheme(path string) *interfaces.IconTheme {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

	}

	return nil
}
