package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	words := []string{"listen", "silent", "enlist", "hello", "world", "lentils"}
	expectedResult := map[string][]string{
		"hello":   {"hello"},
		"lentils": {"lentils"},
		"listen":  {"enlist", "listen", "silent"},
		"world":   {"world"},
	}

	result := getAnagramsMap(words)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v, but got %v", expectedResult, result)
	}
}
