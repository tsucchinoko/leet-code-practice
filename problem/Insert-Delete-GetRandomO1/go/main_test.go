package main

import (
	"testing"
)

// Basic operations test
func TestRandomizedSet_Basic(t *testing.T) {
	rs := Constructor()

	if !rs.Insert(1) {
		t.Fatalf("Insert(1) expected true")
	}
	if rs.Remove(2) {
		t.Fatalf("Remove(2) expected false")
	}
	if !rs.Insert(2) {
		t.Fatalf("Insert(2) expected true")
	}

	// getRandom should return either 1 or 2
	v := rs.GetRandom()
	if v != 1 && v != 2 {
		t.Fatalf("GetRandom() expected 1 or 2, got %d", v)
	}

	if !rs.Remove(1) {
		t.Fatalf("Remove(1) expected true")
	}
	if rs.Insert(2) {
		t.Fatalf("Insert(2) expected false because 2 already present")
	}
	if rs.GetRandom() != 2 {
		t.Fatalf("GetRandom() expected 2 when only 2 remains")
	}
}

// Stress test to ensure no panic and basic invariants hold
func TestRandomizedSet_Stress(t *testing.T) {
	rs := Constructor()
	n := 1000
	// insert 0..n-1
	for i := 0; i < n; i++ {
		if !rs.Insert(i) {
			t.Fatalf("Insert(%d) expected true", i)
		}
	}
	// remove even numbers
	for i := 0; i < n; i += 2 {
		if !rs.Remove(i) {
			t.Fatalf("Remove(%d) expected true", i)
		}
	}
	// check remaining are odd numbers
	for i := 1; i < n; i += 2 {
		// try removing twice: first true, second false
		if !rs.Remove(i) {
			t.Fatalf("Remove(%d) expected true", i)
		}
		if rs.Remove(i) {
			t.Fatalf("Remove(%d) expected false on second removal", i)
		}
	}
}
