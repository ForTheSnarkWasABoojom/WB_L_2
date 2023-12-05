package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

func unpackString(s string) (string, error) {
	var result string
	runeSlice := []rune(s)
	length := len(runeSlice)

	for i := 0; i < length; i++ {
		char := runeSlice[i]

		if unicode.IsDigit(char) {
			if i+1 == length {
				return "", errors.New("некорректная строка")
			}
			if unicode.IsDigit(runeSlice[i+1]) {
				return "", errors.New("некорректная строка")
			}
			count, err := strconv.Atoi(string(char))
			if err != nil {
				return "", errors.New("некорректная строка")
			}

			for j := 0; j < count-1; j++ {
				result += string(runeSlice[i-1])
			}

		} else {
			result += string(char)
		}
	}

	return result, nil
}

func main() {
	testCases := []string{"a4bc2d5e", "abcd", "45", ""}
	for _, testCase := range testCases {
		result, err := unpackString(testCase)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			fmt.Printf("%s => %s\n", testCase, result)
		}
	}
}
