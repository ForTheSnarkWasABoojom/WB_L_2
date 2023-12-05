package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func cut(fields []int, delimiter string, onlySeparated bool, inputLines []string) string {
	var result string
	for _, line := range inputLines {
		if !onlySeparated || strings.Contains(line, delimiter) {
			parts := strings.Split(line, delimiter)
			outputParts := make([]string, 0)

			for _, field := range fields {
				if field > 0 && field <= len(parts) {
					outputParts = append(outputParts, parts[field-1])
				}
			}

			result += strings.Join(outputParts, delimiter)
		}
	}
	return result
}

func main() {
	var fieldsStr string
	var delimiter string
	var onlySeparated bool

	flag.StringVar(&fieldsStr, "f", "1", "Fields to select")
	flag.StringVar(&delimiter, "d", "\t", "Delimiter")
	flag.BoolVar(&onlySeparated, "s", false, "Only lines with delimiter")
	flag.Parse()

	fieldsStr = strings.TrimSpace(fieldsStr)
	fields := make([]int, 0)

	if fieldsStr != "" {
		fieldsStr = strings.Replace(fieldsStr, " ", ",", -1)
		for _, fieldStr := range strings.Split(fieldsStr, ",") {
			var field int
			fmt.Sscanf(fieldStr, "%d", &field)
			fields = append(fields, field)
		}
	} else {
		fmt.Println("No fields specified.")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	inputLines := make([]string, 0)

	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading input:", scanner.Err())
		os.Exit(1)
	}

	fmt.Println(cut(fields, delimiter, onlySeparated, inputLines))
}
