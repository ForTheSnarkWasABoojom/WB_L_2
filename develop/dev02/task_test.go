package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      error
	}{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"45", "", errors.New("некорректная строка")},
	}

	for _, test := range tests {
		result, err := unpackString(test.input)
		assert.Equal(t, test.expected, result)
		assert.Equal(t, test.err, err)
	}
}
