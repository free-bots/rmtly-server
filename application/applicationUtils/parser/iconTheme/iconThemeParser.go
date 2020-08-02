package iconTheme

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"rmtly-server/application/applicationUtils/parser"
	"rmtly-server/application/interfaces"
	"strconv"
)

func ParseIconThemeIndex(themeFolder string) *interfaces.IconTheme {

	file, err := os.Open(themeFolder + string(os.PathSeparator) + "index.theme")
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

	var theme *interfaces.IconTheme

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case parser.IsLineEmpty(line):
			continue
		case parser.IsLineGroup(line):
			theme = iconsGroups(scanner, line, nil, nil)
			continue
		case parser.IsLineComment(line):
			continue
		default:
			fmt.Printf("line not matched %s\n", line)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	}

	if theme != nil {
		theme.RootFolder = themeFolder
	}

	return theme
}

func isLineIconGroup(line string) bool {
	return line == "[Icon Theme]"
}

func parseIconDirectory(scanner *bufio.Scanner) (*interfaces.IconDirectory, *string) {
	directory := new(interfaces.IconDirectory)

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case parser.IsLineEmpty(line):
			continue
		case parser.IsLineGroup(line):
			return directory, &line
		case parser.IsLineComment(line):
			continue
		default:
			parser.OnKey(line, "Size", func(key string, value string) {
				size, err := strconv.ParseInt(value, 10, 32)
				if err != nil {
					log.Fatal(err)
					return
				}

				directory.Size = int(size)
			})
			parser.OnKey(line, "Scale", func(key string, value string) {
				scale, err := strconv.ParseInt(value, 10, 32)
				if err != nil {
					log.Fatal(err)
					return
				}

				directory.Scale = int(scale)
			})
			parser.OnKey(line, "Context", func(key string, value string) {
				directory.Context = value
			})
			parser.OnKey(line, "Type", func(key string, value string) {
				directory.Type = value
			})
			parser.OnKey(line, "MaxSize", func(key string, value string) {
				maxSize, err := strconv.ParseInt(value, 10, 32)
				if err != nil {
					log.Fatal(err)
					return
				}

				directory.MaxSize = int(maxSize)
			})
			parser.OnKey(line, "MinSize", func(key string, value string) {
				minSize, err := strconv.ParseInt(value, 10, 32)
				if err != nil {
					log.Fatal(err)
					return
				}

				directory.MinSize = int(minSize)
			})
			parser.OnKey(line, "Threshold", func(key string, value string) {
				threshold, err := strconv.ParseInt(value, 10, 32)
				if err != nil {
					log.Fatal(err)
					return
				}

				directory.Threshold = int(threshold)
			})
		}

	}

	return directory, nil
}

func iconsGroups(scanner *bufio.Scanner, line string, theme *interfaces.IconTheme, directories []*interfaces.IconDirectory) *interfaces.IconTheme {
	if isLineIconGroup(line) {
		iconTheme, stopLine := parseIconGroup(scanner)
		if stopLine != nil {
			return iconsGroups(scanner, *stopLine, iconTheme, directories)
		}
	}

	if isLineDirectoryGroup(line) {
		iconDirectory, stopLine := parseIconDirectory(scanner)
		if iconDirectory != nil && theme != nil {
			theme.DirectoriesData = append(theme.DirectoriesData, *iconDirectory)
		}

		if stopLine != nil {
			return iconsGroups(scanner, *stopLine, theme, directories)
		}
	}

	return theme
}

func parseIconGroup(scanner *bufio.Scanner) (*interfaces.IconTheme, *string) {
	iconTheme := new(interfaces.IconTheme)
	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case parser.IsLineEmpty(line):
			continue
		case parser.IsLineGroup(line):
			// return iconGroup
			return iconTheme, &line
		case parser.IsLineComment(line):
			continue
		default:
			parser.OnKey(line, "Name", func(key string, value string) {
				iconTheme.Name = value
			})
			parser.OnKey(line, "Comment", func(key string, value string) {
				iconTheme.Comment = value
			})
			parser.OnKey(line, "Inherits", func(key string, value string) {
				if parser.IsColonList(value) {
					iconTheme.Inherits = parser.GetColonList(value)
				}
			})
			parser.OnKey(line, "Directories", func(key string, value string) {
				if parser.IsColonList(value) {
					iconTheme.Directories = parser.GetColonList(value)
				}
			})
			parser.OnKey(line, "ScaledDirectories", func(key string, value string) {
				if parser.IsColonList(value) {
					iconTheme.ScaledDirectories = parser.GetColonList(value)
				}
			})
			parser.OnKey(line, "Hidden", func(key string, value string) {
				hidden, err := strconv.ParseBool(value)
				if err != nil {
					log.Fatal(err)
					return
				}
				iconTheme.Hidden = hidden
			})
			parser.OnKey(line, "Example", func(key string, value string) {
				iconTheme.Example = value
			})
		}

	}

	return iconTheme, nil
}

func isLineDirectoryGroup(line string) bool {
	matched, err := regexp.Match("^\\[.*/.*]", []byte(line))
	if err != nil {
		log.Fatal(err)
		return false
	}
	return matched
}
