package main

import (
	"testing"
)

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		a, b string
		want string
	}{
		{"abc", "def", "adbecf"},
		{"ab", "pqrs", "apbqrs"},
		{"abcd", "pq", "apbqcd"},
	}

	for _, tt := range tests {
		got := mergeAlternately(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("mergeAlternately(%q, %q) = %q; want %q", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestMergeAlternately2(t *testing.T) {
	tests := []struct {
		a, b string
		want string
	}{
		{"abc", "def", "adbecf"},
		{"ab", "pqrs", "apbqrs"},
		{"abcd", "pq", "apbqcd"},
	}

	for _, tt := range tests {
		got := mergeAlternately2(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("mergeAlternately(%q, %q) = %q; want %q", tt.a, tt.b, got, tt.want)
		}
	}
}
