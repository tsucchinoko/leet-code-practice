package main

import (
	"reflect"
	"testing"
)

func toByteSlice(ss []string) []byte {
	b := make([]byte, len(ss))
	for i := range ss {
		b[i] = ss[i][0]
	}
	return b
}

func sliceEqualBytes(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestCompressExamples(t *testing.T) {
	tests := []struct {
		input   []byte
		want    []byte
		wantLen int
	}{
		{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, []byte{'a', '2', 'b', '2', 'c', '3'}, 6},
		{[]byte{'a'}, []byte{'a'}, 1},
		{[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, []byte{'a', 'b', '1', '2'}, 4},
		{[]byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}, []byte{'a', '1', '0'}, 3}, // 10 a's
		{[]byte{'a', 'b', 'c'}, []byte{'a', 'b', 'c'}, 3},
	}

	for _, tt := range tests {
		// copy input to avoid mutating test case data across runs
		inp := make([]byte, len(tt.input))
		copy(inp, tt.input)
		gotLen := compress(inp)
		if gotLen != tt.wantLen {
			t.Fatalf("compress(%q) returned length %d, want %d", tt.input, gotLen, tt.wantLen)
		}
		if !sliceEqualBytes(inp[:gotLen], tt.want) {
			t.Fatalf("compress(%q) = %q (len %d), want %q (len %d)", tt.input, inp[:gotLen], gotLen, tt.want, tt.wantLen)
		}
	}
}

func TestCompressRandom(t *testing.T) {
	// basic additional sanity: repeating groups and singletons
	inp := []byte{'z', 'z', 'z', 'y', 'x', 'x', '1', '1', '1', '1'}
	want := []byte{'z', '3', 'y', 'x', '2', '1', '4'}
	gotLen := compress(inp)
	if gotLen != len(want) {
		t.Fatalf("length got %d want %d", gotLen, len(want))
	}
	if !reflect.DeepEqual(inp[:gotLen], want) {
		t.Fatalf("got %q want %q", inp[:gotLen], want)
	}
}
