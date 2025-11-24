package main

import "testing"

func TestRecentCounter_Example(t *testing.T) {
	rc := Constructor()
	if got := rc.Ping(1); got != 1 {
		t.Fatalf("Ping(1) = %d, want 1", got)
	}
	if got := rc.Ping(100); got != 2 {
		t.Fatalf("Ping(100)= %d, want 2", got)
	}
	if got := rc.Ping(3001); got != 3 {
		t.Fatalf("Ping(3001) = %d, want 3", got)
	}
	if got := rc.Ping(3002); got != 3 {
		t.Fatalf("Ping(3002) = %d, want 3", got)
	}
}

func TestRecentCounter_Multiple(t *testing.T) {
	rc := Constructor()
	seq := []int{1, 2000, 3000, 4000, 7000}
	wants := []int{1, 2, 3, 3, 1}
	for i, v := range seq {
		if got := rc.Ping(v); got != wants[i] {
			t.Fatalf("Ping(%d) = %d, want %d", v, got, wants[i])
		}
	}
}
