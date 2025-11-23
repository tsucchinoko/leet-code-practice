package main

import (
	"testing"
)

func TestDecodeString(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"10[a]", "aaaaaaaaaa"},
		{"1[abc]", "abc"},
		{"2[3[a]b]", "aaabaaab"},
		{"3[z2[y]]", "zyyzyyzyy"},
	}

	for _, c := range cases {
		got := decodeString(c.in)
		if got != c.out {
			t.Fatalf("decodeString(%q) == %q, want %q", c.in, got, c.out)
		}
	}
}
