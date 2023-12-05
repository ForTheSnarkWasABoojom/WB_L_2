package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCut(t *testing.T) {
	tests := []struct {
		name           string
		fields         string
		delimiter      string
		onlySeparated  bool
		inputLines     []string
		expectedOutput string
	}{
		{
			name:           "Basic Test",
			fields:         "1",
			delimiter:      "\t",
			onlySeparated:  false,
			inputLines:     []string{"apple\torange\tbanana"},
			expectedOutput: "apple",
		},
		{
			name:           "Selecting Multiple Fields",
			fields:         "2 3",
			delimiter:      ",",
			onlySeparated:  true,
			inputLines:     []string{"apple,orange,banana"},
			expectedOutput: "orange,banana",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var fields []int
			if test.fields != "" {
				fieldStrs := strings.Split(test.fields, " ")
				for _, fieldStr := range fieldStrs {
					var field int
					fmt.Sscanf(fieldStr, "%d", &field)
					fields = append(fields, field)
				}
			}

			output := cut(fields, test.delimiter, test.onlySeparated, test.inputLines)

			if output != test.expectedOutput {
				t.Errorf("Expected output %v, but got %v", test.expectedOutput, output)
			}
		})
	}
}
