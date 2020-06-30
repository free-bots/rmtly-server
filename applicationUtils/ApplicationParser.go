package applicationUtils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"rmtly-server/interfaces"
	"strings"
)

// parser for .desktop files
func Parse(path string) *interfaces.ApplicationEntry {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer file.Close()

	applicationEntry := new(interfaces.ApplicationEntry)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case isLineEmpty(line):
			continue
		case isLineGroup(line):
		case isLineComment(line):
			applicationEntry.Comment = getValue(line)
		default:
			fmt.Printf("line not matched %s", line)
		}

		fmt.Println(splitLineToKeyValue(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func getValue(line string) string {
	_, value, splitErr := splitLineToKeyValue(line)
	if splitErr != nil {
		log.Fatal(splitErr)
	}
	return value
}

func splitLineToKeyValue(line string) (key string, value string, err error) {
	split := strings.Split(line, "=")
	if split == nil || len(split) <= 1 {
		return "", "", fmt.Errorf("spliting line to key value failed line:%s", line)
	}
	return split[0], split[1], nil
}

func isLineComment(line string) bool {
	index := strings.Index(line, "#")
	return index == 0
}

func isLineEmpty(line string) bool {
	line = strings.TrimSpace(line)
	return line == ""
}

func isLineGroup(line string) bool {
	// todo check if line contains [*****]
	return false
}

func splitToActions() {

}
