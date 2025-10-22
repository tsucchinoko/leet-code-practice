package main

import (
	"fmt"
)

// lengthOfLastWord returns the length of the last word in s.
// A word is a maximal substring consisting of non-space characters.
func lengthOfLastWord(s string) int {
	n := len(s)
	i := n - 1

	// skip trailing spaces
	for i >= 0 && s[i] == ' ' {
		i--
	}

	// count non-space characters
	length := 0
	for i >= 0 && s[i] != ' ' {
		length++
		i--
	}

	return length
}

func main() {
	examples := []string{
		"Hello World",
		"   fly me   to   the moon  ",
		"luffy is still joyboy",
	}

	for _, ex := range examples {
		fmt.Printf("Input: %q -> Output: %d\n", ex, lengthOfLastWord(ex))
	}
}
