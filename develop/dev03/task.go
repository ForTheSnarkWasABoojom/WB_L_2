package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Example input: go run task.go -k=1 -r -u input.txt
func main() {
	column := flag.Int("k", 0, "Specifying a column to sort")
	numeric := flag.Bool("n", false, "Sort by numeric value")
	reverse := flag.Bool("r", false, "Sort in reverse order")
	unique := flag.Bool("u", false, "Don't print duplicate lines")
	flag.Parse()

	pathForFile, _ := os.Getwd()
	pathForFile += `\` + os.Args[len(os.Args)-1]
	fmt.Println(pathForFile)
	file, err := os.Open(pathForFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	if *column > 0 {
		sortByColumn(lines, *column, *numeric)
	}

	if *numeric {
		sortByNumericValue(lines)
	}

	if *unique {
		lines = removeDuplicates(lines)
	}

	if *reverse {
		reverseLines(lines)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

func sortByColumn(strings []string, column int, isNumeric bool) {
	sort.Slice(strings, func(i, j int) bool {
		columnI := getSubstring(strings[i], column)
		columnJ := getSubstring(strings[j], column)

		if isNumeric {
			numI, errI := strconv.Atoi(strings[i])
			numJ, errJ := strconv.Atoi(strings[j])

			if errI == nil && errJ == nil {
				return numI < numJ
			}
		}

		return columnI < columnJ
	})
}

func getSubstring(s string, column int) string {
	column--
	words := strings.Fields(s)

	if column < len(words) {
		return words[column]
	}

	return ""
}

func sortByNumericValue(strings []string) {
	sort.Slice(strings, func(i, j int) bool {
		numI, errI := strconv.Atoi(strings[i])
		numJ, errJ := strconv.Atoi(strings[j])

		if errI == nil && errJ == nil {
			return numI < numJ
		}

		return strings[i] < strings[j]
	})
}

func min(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

func reverseLines(strs []string) {
	last := len(strs) - 1
	for i := 0; i < len(strs)/2; i++ {
		strs[i], strs[last-i] = strs[last-i], strs[i]
	}
}

func removeDuplicates(lines []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, line := range lines {
		if !seen[line] {
			result = append(result, line)
			seen[line] = true
		}
	}
	return result
}
