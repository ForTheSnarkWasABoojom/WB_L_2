package main

import (
	"reflect"
	"testing"
)

func TestGetSubstring(t *testing.T) {
	tests := []struct {
		input  string
		column int
		output string
	}{
		{"a b c d", 1, "a"},
		{"a b c d", 2, "b"},
		{"a b c d", 4, "d"},
		{"a b c d", 5, ""},
	}

	for _, test := range tests {
		result := getSubstring(test.input, test.column)
		if result != test.output {
			t.Errorf("Expected %s for column %d, but got %s", test.output, test.column, result)
		}
	}
}

func TestSortByColumn(t *testing.T) {
	tests := []struct {
		input     []string
		column    int
		isNumeric bool
		output    []string
	}{
		{[]string{"3 a", "1 b", "2 c"}, 1, false, []string{"1 b", "2 c", "3 a"}},
		{[]string{"3 a", "1 b", "2 c"}, 1, true, []string{"1 b", "2 c", "3 a"}},
	}

	for _, test := range tests {
		sortByColumn(test.input, test.column, test.isNumeric)
		if !reflect.DeepEqual(test.input, test.output) {
			t.Errorf("Expected %v, but got %v", test.output, test.input)
		}
	}
}

func TestSortByNumericValue(t *testing.T) {
	tests := []struct {
		input  []string
		output []string
	}{
		{[]string{"3", "1", "2"}, []string{"1", "2", "3"}},
	}

	for _, test := range tests {
		sortByNumericValue(test.input)
		if !reflect.DeepEqual(test.input, test.output) {
			t.Errorf("Expected %v, but got %v", test.output, test.input)
		}
	}
}

func TestReverseLines(t *testing.T) {
	tests := []struct {
		input  []string
		output []string
	}{
		{[]string{"a", "b", "c"}, []string{"c", "b", "a"}},
		{[]string{"1", "2", "3"}, []string{"3", "2", "1"}},
	}

	for _, test := range tests {
		reverseLines(test.input)
		if !reflect.DeepEqual(test.input, test.output) {
			t.Errorf("Expected %v, but got %v", test.output, test.input)
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input  []string
		output []string
	}{
		{[]string{"a", "b", "c", "a", "b"}, []string{"a", "b", "c"}},
		{[]string{"1", "2", "3", "1", "2"}, []string{"1", "2", "3"}},
	}

	for _, test := range tests {
		result := removeDuplicates(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("Expected %v, but got %v", test.output, result)
		}
	}
}
