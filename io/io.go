package io

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func ReadLines(filename string) ([]string, error) {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		return lines, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func WriteLines(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	return writer.Flush()
}

func ParsePath(path string) string {
	regex := regexp.MustCompile(`(\S+)\.\w+$`)
	if !regex.MatchString(path) {
		fmt.Println("Path '%v' does not point to a file\n", path)
		panic(path)
	}
	result := regex.FindStringSubmatch(path)
	return result[len(result)-1]
}
