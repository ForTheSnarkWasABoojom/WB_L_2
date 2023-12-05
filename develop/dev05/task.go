package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func grep(lines []string, pattern string, options Options) []string {
	matchingLines := make([]string, 0)

	for i, line := range lines {
		if options.ignoreCase {
			line = strings.ToLower(line)
			pattern = strings.ToLower(pattern)
		}

		if options.fixed && line == pattern {
			matchingLines = append(matchingLines, line)
		}
		if options.invertMatch {
			if !strings.Contains(line, pattern) {
				matchingLines = append(matchingLines, line)
			}
		} else if options.lineNumber {
			if strings.Contains(line, pattern) {
				lineNum := strconv.Itoa(i)
				matchingLines = append(matchingLines, lineNum)
			}
		}

		if strings.Contains(line, pattern) && !options.invertMatch {
			if options.beforeContext > 0 {
				beforeStart := max(0, i-options.beforeContext)
				beforeEnd := max(0, i)
				matchingLines = append(matchingLines, lines[beforeStart:beforeEnd]...)
			}
			matchingLines = append(matchingLines, line)

			if options.afterContext > 0 {
				afterStart := min(len(lines), i+1)
				afterEnd := min(len(lines), i+1+options.afterContext)
				matchingLines = append(matchingLines, lines[afterStart:afterEnd]...)
			}
		}

	}
	if options.count {
		res := make([]string, 1)
		res[0] = strconv.Itoa(len(matchingLines))

		return res
	} else {
		return matchingLines
	}
}

type Options struct {
	afterContext  int
	beforeContext int
	context       int
	count         bool
	ignoreCase    bool
	invertMatch   bool
	fixed         bool
	lineNumber    bool
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var options Options

	flag.IntVar(&options.afterContext, "A", 0, "Print N lines after each match")
	flag.IntVar(&options.beforeContext, "B", 0, "Print N lines before each match")
	flag.IntVar(&options.context, "C", 0, "Print N lines after and before each match")
	flag.BoolVar(&options.count, "c", false, "Print only a count of matching lines")
	flag.BoolVar(&options.ignoreCase, "i", false, "IgnoreCase")
	flag.BoolVar(&options.invertMatch, "v", false, "Invert the sense of matching")
	flag.BoolVar(&options.fixed, "F", false, "Fixed, exact matching")
	flag.BoolVar(&options.lineNumber, "n", false, "Print line number with output")
	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		fmt.Println("Usage: grep [options] pattern file")
		os.Exit(1)
	}

	pattern := args[0]
	filePath := args[1]

	if options.context > 0 {
		options.afterContext = options.context
		options.beforeContext = options.context
	}
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading file:", scanner.Err())
		os.Exit(1)
	}

	grep(lines, pattern, options)

}
