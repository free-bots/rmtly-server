package applicationUtils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

	//applicationEntry := new(interfaces.ApplicationEntry)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case isLineEmpty(line):
			continue
		case isLineGroup(line):
			if isLineDesktopEntry(line) {
				parseEntry(scanner)
			}

			// check if is group or action group
			// if action group stop line parsing and continue group parsing

			fmt.Println(line)
			continue
		case isLineComment(line):
			continue
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
	matched, err := regexp.MatchString("^\\[.*]", line)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return matched
}

func isLineDesktopEntry(line string) bool {
	return line == "[Desktop Entry]"
}

func isLineDesktopAction(line string) bool {
	// todo check with regex
	return line == "[Desktop Action Edit]"
}

func getDesktopActionName(line string) string {
	return line
}

func splitToActions() {

}

func onKey(line string, key string, callback func(key string, value string)) {
	currentKey, value, err := splitLineToKeyValue(line)
	if err != nil {
		fmt.Printf("error finding key %s with line %s", key, line)
		return
	}
	if currentKey == key {
		callback(key, value)
	}
}

func parseEntry(scanner *bufio.Scanner) *interfaces.ApplicationEntry {
	fmt.Println("using entry parser")

	entry := new(interfaces.ApplicationEntry)

	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case isLineEmpty(line):
			continue
		case isLineGroup(line):
			return entry
		default:
			onKey(line, "Version", func(key string, value string) {
				//entry.Version = value
			})
			onKey(line, "Type", func(key string, value string) {
				entry.Type = value
			})
			onKey(line, "Name", func(key string, value string) {
				entry.Name = value
			})
			onKey(line, "Comment", func(key string, value string) {
				entry.Comment = value
			})
			onKey(line, "TryExec", func(key string, value string) {
				entry.TryExec = value
			})
			onKey(line, "Name", func(key string, value string) {
				entry.Name = value
			})
			onKey(line, "Exec", func(key string, value string) {
				entry.Exec = value
			})
			onKey(line, "Icon", func(key string, value string) {
				entry.Icon = value
			})
			onKey(line, "MimeType", func(key string, value string) {
				entry.MimeType = value
			})
			onKey(line, "Actions", func(key string, value string) {
			})

			fmt.Printf("line %s ignored\n", line)

		}
	}
	return nil
}

func parseAction(scanner *bufio.Scanner) *interfaces.Action {
	return nil
}
