package main

import (
	"reflect"
	"testing"
)

func TestGrep(t *testing.T) {
	tests := []struct {
		name          string
		lines         []string
		pattern       string
		options       Options
		expectedLines []string
	}{
		{
			name: "Simple match",
			lines: []string{
				"apple",
				"banana",
				"orange",
			},
			pattern: "banana",
			options: Options{},
			expectedLines: []string{
				"banana",
			},
		},
		{
			name: "Invert match",
			lines: []string{
				"apple",
				"banana",
				"orange",
			},
			pattern: "b",
			options: Options{invertMatch: true},
			expectedLines: []string{
				"apple",
				"orange",
			},
		},
		{
			name: "Case-insensitive match",
			lines: []string{
				"apple",
				"banana",
				"orange",
			},
			pattern: "aPPle",
			options: Options{ignoreCase: true},
			expectedLines: []string{
				"apple",
			},
		},
		{
			name: "Count lines",
			lines: []string{
				"apple",
				"banana",
				"orange",
			},
			pattern:       "banana",
			options:       Options{count: true},
			expectedLines: []string{"1"},
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resultLines := grep(test.lines, test.pattern, test.options)

			if !reflect.DeepEqual(resultLines, test.expectedLines) {
				t.Errorf("Expected %v, but got %v", test.expectedLines, resultLines)
			}
		})
	}
}
