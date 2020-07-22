package applicationUtils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"rmtly-server/application/interfaces"
	"strconv"
)

func ParseIconThemeIndex(path string) *interfaces.IconTheme {

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

		for scanner.Scan() {
			line := scanner.Text()

			switch {
			case isLineEmpty(line):
				continue
			case isLineGroup(line):
				iconsGroups(*scanner, line, nil, nil)
				// todo parse icon
				continue
			case isLineComment(line):
				continue
			default:
				fmt.Printf("line not matched %s\n", line)
			}

		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	}

	return nil
}

func isLineIconGroup(line string) bool {
	return line == "[Icon Theme]"
}

func isLineScalableGroup(line string) bool {
	matched, err := regexp.Match("^\\[scalable/.*]", []byte(line))
	if err != nil {
		log.Fatal(err)
		return false
	}
	return matched
}

func parseIconDirectory(scanner bufio.Scanner) *interfaces.IconDirectory {
	return nil
}

func iconsGroups(scanner bufio.Scanner, line string, theme *interfaces.IconTheme, directories []*interfaces.IconDirectory) *interfaces.IconTheme {
	if isLineIconGroup(line) {

	}

	if isLineDirectoryGroup(line) {
		// todo parse directory data
	}

	iconTheme := new(interfaces.IconTheme)
	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case isLineEmpty(line):
			continue
		case isLineGroup(line):
			// return iconGroup
			continue
		case isLineComment(line):
			continue
		default:
			onKey(line, "Name", func(key string, value string) {
				iconTheme.Name = value
			})
			onKey(line, "Comment", func(key string, value string) {
				iconTheme.Comment = value
			})
			onKey(line, "Inherits", func(key string, value string) {
				if isColonList(value) {
					iconTheme.Inherits = getColonList(value)
				}
			})
			onKey(line, "Directories", func(key string, value string) {
				if isColonList(value) {
					iconTheme.Directories = getColonList(value)
				}
			})
			onKey(line, "ScaledDirectories", func(key string, value string) {
				if isColonList(value) {
					iconTheme.ScaledDirectories = getColonList(value)
				}
			})
			onKey(line, "Hidden", func(key string, value string) {
				hidden, err := strconv.ParseBool(value)
				if err != nil {
					log.Fatal(err)
					return
				}
				iconTheme.Hidden = hidden
			})
			onKey(line, "Example", func(key string, value string) {
				iconTheme.Example = value
			})
		}

	}

	return nil
}

func isLineDirectoryGroup(line string) bool {
	matched, err := regexp.Match("^\\[.*/.*]", []byte(line))
	if err != nil {
		log.Fatal(err)
		return false
	}
	return matched
}
