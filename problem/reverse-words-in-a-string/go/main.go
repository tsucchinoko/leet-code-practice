package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// 1. Trim leading/trailing spaces
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return ""
	}

	// 2. Split by spaces (fields handles multiple spaces)
	words := strings.Fields(s) // Fields splits by any unicode whitespace and collapses multiple spaces

	// 3. Reverse slice of words in-place
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// 4. Join with single space
	return strings.Join(words, " ")
}

// reverseRange reverses b[i:j] inclusive (i and j are indices)
func reverseRange(b []byte, i, j int) {
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
}

// reverseWordsInPlace takes a string, converts to []byte, and performs in-place algorithm.
// Returns the resulting string with words reversed and spaces normalized.
func reverseWordsInPlace(s string) string {
	// Convert to mutable []byte
	b := []byte(s)
	n := len(b)
	if n == 0 {
		return ""
	}

	// 1) Reverse whole array
	reverseRange(b, 0, n-1)

	// 2) Traverse, reverse each word (they are reversed because of step 1),
	//    and write them left-justified with a single space between words.
	write := 0
	i := 0

	for i < n {
		// Skip spaces
		for i < n && b[i] == ' ' {
			i++
		}
		if i >= n {
			break
		}
		// j: end of the current word (inclusive)
		j := i
		for j < n && b[j] != ' ' {
			j++
		}
		// word is in b[i:j-1]
		// reverse this word to restore correct char order
		reverseRange(b, i, j-1)

		// If not the first word, write a single space before the word
		if write != 0 {
			b[write] = ' '
			write++
		}

		// Copy the word to the write position (may overlap; copy forward)
		for k := i; k < j; k++ {
			b[write] = b[k]
			write++
		}

		// Move i to j (next potential word)
		i = j
	}

	// Cut off any trailing bytes beyond write
	return string(b[:write])
}

func main() {
	examples := []string{
		"the sky is blue",
		"  hello world  ",
		"a good   example",
	}
	for _, e := range examples {
		fmt.Printf("Input: %q -> Output: %q\n", e, reverseWords(e))
		fmt.Printf("Input: %q -> Output: %q\n", e, reverseWordsInPlace(e))
	}
}
