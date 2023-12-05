package main

import (
	"sort"
	"strings"

	"golang.org/x/exp/slices"
)

var dict = map[string]string{}

func findAnagrams(words []string) map[string][]string {
	anagramSets := make(map[string][]string)

	dict = make(map[string]string)
	for _, word := range words {
		sortedWord := sortString(strings.ToLower(word))
		if !slices.Contains(anagramSets[sortedWord], word) {
			anagramSets[sortedWord] = append(anagramSets[sortedWord], word)
			if len(anagramSets[sortedWord]) == 1 {
				dict[sortedWord] = word
			}
		}
	}

	return anagramSets
}

func sortString(s string) string {
	symbols := strings.Split(s, "")
	sort.Strings(symbols)
	return strings.Join(symbols, "")
}

func getAnagramsMap(dictionary []string) map[string][]string {
	anagramSets := findAnagrams(dictionary)
	result := make(map[string][]string)

	for key, words := range anagramSets {
		sort.Strings(words)
		result[dict[key]] = words
	}

	return result
}
