package main

import (
	"testing"
)

func TestGetBaseURL(t *testing.T) {
	tests := []struct {
		url      string
		expected string
	}{
		{"http://localhost:8080/page", "http://localhost:8080"},
	}

	for _, test := range tests {
		result := getBaseURL(test.url)
		if result != test.expected {
			t.Errorf("For URL %s expected %s, but got %s", test.url, test.expected, result)
		}
	}
}
