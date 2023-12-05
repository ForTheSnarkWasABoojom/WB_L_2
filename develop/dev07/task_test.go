package main

import (
	"testing"
	"time"
)

func TestOrFunction(t *testing.T) {
	tests := []struct {
		name     string
		channels []<-chan interface{}
		expected bool
	}{
		{
			name: "All Channels Remain Open",
			channels: []<-chan interface{}{
				make(chan interface{}),
				make(chan interface{}),
				make(chan interface{}),
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isClosed(or(test.channels...))
			if result != test.expected {
				t.Errorf("Ожидалось: %v, Получено: %v", test.expected, result)
			}
		})
	}
}

func isClosed(ch <-chan interface{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func sig7(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}
