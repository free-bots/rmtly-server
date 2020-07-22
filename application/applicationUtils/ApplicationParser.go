package applicationUtils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"rmtly-server/application/interfaces"
	"strconv"
	"strings"
)

// parser for .desktop files
func Parse(path string, removeExecFields bool) *interfaces.ApplicationEntry {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	var applicationEntry *interfaces.ApplicationEntry
	var actions = make([]*interfaces.Action, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case isLineEmpty(line):
			continue
		case isLineGroup(line):
			applicationEntry, actions = parseGroup(scanner, removeExecFields, line, nil, actions)
			continue
		case isLineComment(line):
			continue
		default:
			fmt.Printf("line not matched %s\n", line)
		}

		fmt.Println(splitLineToKeyValue(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if applicationEntry != nil {
		applicationEntry.Actions = actions

		splitFileName := strings.Split(file.Name(), "/")
		length := len(splitFileName)
		applicationEntry.Id = splitFileName[length-1]
	}

	return applicationEntry
}

func isLineDesktopEntry(line string) bool {
	return line == "[Desktop Entry]"
}

func isLineDesktopAction(line string) bool {
	matched, err := regexp.MatchString("^\\[Desktop Action.*]", line)
	if err != nil {
		log.Fatal(err)
		return false
	}
	// todo check with regex
	return matched
}

func parseEntry(scanner *bufio.Scanner, removeExecFields bool) (*interfaces.ApplicationEntry, *string) {
	fmt.Println("using entry parser")

	entry := new(interfaces.ApplicationEntry)

	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case isLineEmpty(line):
			continue
		case isLineComment(line):
			continue
		case isLineGroup(line):
			return entry, &line
		default:
			onKey(line, "Version", func(key string, value string) {
				float, err := strconv.ParseFloat(value, 32)
				if err != nil {
					fmt.Println(err)
					return
				}

				entry.Version = float32(float)
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
				if removeExecFields {
					entry.Exec = removeExecFieldCodes(value)
				} else {
					entry.Exec = value
				}
			})
			onKey(line, "Icon", func(key string, value string) {
				entry.Icon = value
			})
			onKey(line, "MimeType", func(key string, value string) {
				if isSemicolonList(value) {
					entry.MimeType = getSemicolonList(value)
				}
			})
			onKey(line, "Actions", func(key string, value string) {
			})
			onKey(line, "Categories", func(key string, value string) {
				if isSemicolonList(value) {
					entry.Categories = getSemicolonList(value)
				}
			})

			//fmt.Printf("line %s ignored\n", line)

		}
	}
	return entry, nil
}

func parseAction(scanner *bufio.Scanner, removeExecFields bool) (*interfaces.Action, *string) {

	var action = new(interfaces.Action)

	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case isLineEmpty(line):
			continue
		case isLineComment(line):
			continue
		case isLineGroup(line):
			return action, &line
		default:
			onKey(line, "Name", func(key string, value string) {
				action.Name = value
			})
			onKey(line, "Exec", func(key string, value string) {
				if removeExecFields {
					action.Exec = removeExecFieldCodes(value)
				} else {
					action.Exec = value
				}
			})
			onKey(line, "Icon", func(key string, value string) {
				action.Icon = value
			})
		}
	}

	return action, nil
}

func parseGroup(scanner *bufio.Scanner, removeExecFields bool, line string, applicationEntry *interfaces.ApplicationEntry, actions []*interfaces.Action) (*interfaces.ApplicationEntry, []*interfaces.Action) {
	if isLineDesktopEntry(line) {
		var stopLine *string
		applicationEntry, stopLine = parseEntry(scanner, removeExecFields)
		if stopLine != nil {
			return parseGroup(scanner, removeExecFields, *stopLine, applicationEntry, actions)
		}
	}

	if isLineDesktopAction(line) {
		fmt.Println("action")
		var newAction, stopLine = parseAction(scanner, removeExecFields)
		if newAction != nil {
			actions = append(actions, newAction)
		}
		if stopLine != nil {
			return parseGroup(scanner, removeExecFields, *stopLine, applicationEntry, actions)
		}
	}

	return applicationEntry, actions
}

func removeExecFieldCodes(value string) string {
	return strings.NewReplacer(
		"%f", "",
		"%F", "",
		"%u", "",
		"%U", "",
		"%d", "",
		"%D", "",
		"%n", "",
		"%N", "",
		"%i", "",
		"%c", "",
		"%k", "",
		"%v", "",
		"%m", "").Replace(value)
}
