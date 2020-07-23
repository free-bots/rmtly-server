package parser

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func SplitLineToKeyValue(line string) (key string, value string, err error) {
	split := strings.Split(line, "=")
	if split == nil || len(split) <= 1 {
		return "", "", fmt.Errorf("spliting line to key value failed line:%s\n", line)
	}
	return split[0], split[1], nil
}

func IsLineComment(line string) bool {
	index := strings.Index(line, "#")
	return index == 0
}

func IsLineEmpty(line string) bool {
	line = strings.TrimSpace(line)
	return line == ""
}

func IsLineGroup(line string) bool {
	matched, err := regexp.MatchString("^\\[.*]", line)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return matched
}

func OnKey(line string, key string, callback func(key string, value string)) {
	currentKey, value, err := SplitLineToKeyValue(line)
	if err != nil {
		fmt.Printf("error finding key %s with line %s\n", key, line)
		return
	}
	if currentKey == key && value != "" {
		callback(key, value)
	}
}

func IsSemicolonList(line string) bool {
	index := strings.Index(line, ";")
	return index > -1
}

func GetSemicolonList(line string) []string {
	splitItems := strings.Split(line, ";")

	list := make([]string, 0)

	for _, value := range splitItems {
		if value != "" {
			list = append(list, value)
		}
	}

	return list
}

func IsColonList(line string) bool {
	index := strings.Index(line, ",")
	return index > -1
}

func GetColonList(line string) []string {
	splitItems := strings.Split(line, ",")

	list := make([]string, 0)

	for _, value := range splitItems {
		if value != "" {
			list = append(list, value)
		}
	}

	return list
}
