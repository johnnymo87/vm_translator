package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var comment = regexp.MustCompile(`//(.*)$`)
var pushpop = regexp.MustCompile(`^(?i)(push|pop) (\w+) (\d+)$`)

func Parse(lines []string) []string {
	var vm_lines []string
	for _, line := range lines {
		safe := comment.ReplaceAllString(line, "")
		_, err := Type(safe)
		if err == nil {
			vm_lines = append(vm_lines, safe)
		}
	}
	if len(vm_lines) == 0 {
		fmt.Println("Failed to parse any lines!\n")
		panic(vm_lines)
	}
	return vm_lines
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Type(command string) (string, error) {
	arithmetic := []string{"add", "sub", "neg"}
	logical := []string{"eq", "gt", "lt"}
	bitwise := []string{"and", "or", "not"}
	switch {
	case stringInSlice(command, arithmetic):
		return "ARITHMETIC", nil
	case stringInSlice(command, logical):
		return "LOGICAL", nil
	case stringInSlice(command, bitwise):
		return "BITWISE", nil
	case pushpop.MatchString(command):
		return "PUSHPOP", nil
	default:
		return "", errors.New("Unrecognized command type")
	}
}

func ReadPushPop(command string) (string, string, int, error) {
	switch {
	case pushpop.MatchString(command):
		matches := pushpop.FindStringSubmatch(command)
		comm := matches[1]
		segment := matches[2]
		num, err := strconv.Atoi(matches[3])
		if err != nil {
			panic(err)
		}
		return comm, segment, num, nil
	default:
		return "", "", 0, errors.New("Command not of type PUSHPOP")
	}
}
