package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestProcessCommand(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"echo Hello World", "Hello World\n"},
		{"", ""},
		{"pwd", "D:\\PROJECTS\\Golang\\WB_2\\WB_L_2\\develop\\dev08\n"},
		{"cd /", ""},
		{"unknownCommand", "Unexpected error: exec: \"unknownCommand\": executable file not found in %PATH%\n"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			output := captureOutput(func() {
				processCommand(tc.input)
			})

			if output != tc.output {
				t.Errorf("Expected output %q, but got %q", tc.output, output)
			}
		})
	}
}

func captureOutput(f func()) string {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
