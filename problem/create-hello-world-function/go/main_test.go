package main

import (
	"testing"
)

func TestCreateHelloWorld(t *testing.T) {
	test := []struct {
		name     string
		args     []any
		expected string
	}{
		{
			name:     "Hello World",
			args:     []any{},
			expected: "Hello World!",
		},
		{
			name:     "Hello World with Name",
			args:     []any{"John", 1, false},
			expected: "Hello World!",
		},
	}
	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := createHelloWorld(tc.args...)
			if result() != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, result())
			}
		})
	}
}
