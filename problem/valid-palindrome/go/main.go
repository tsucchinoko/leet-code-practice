package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// isLetterOrDigit returns true if b is an ASCII letter or digit.
func isLetterOrDigit(b byte) bool {
	return (b >= '0' && b <= '9') ||
		(b >= 'A' && b <= 'Z') ||
		(b >= 'a' && b <= 'z')
}

// normalizeASCII converts an ASCII uppercase letter to lowercase;
// leaves digits and lowercase letters unchanged.
func normalizeASCII(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}

// advanceLeft moves left index to the next alphanumeric character or beyond right.
func advanceLeft(s string, left, right int) int {
	for left <= right && !isLetterOrDigit(s[left]) {
		left++
	}
	return left
}

// advanceRight moves right index to the previous alphanumeric character or before left.
func advanceRight(s string, left, right int) int {
	for right >= left && !isLetterOrDigit(s[right]) {
		right--
	}
	return right
}

// isPalindrome checks if s is a palindrome using two-pointer technique.
// Assumes ASCII input; ignores non-alphanumeric characters and is case-insensitive.
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		left = advanceLeft(s, left, right)
		right = advanceRight(s, left, right)

		if left >= right {
			break
		}

		if normalizeASCII(s[left]) != normalizeASCII(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

// trimTrailingNewline removes a single trailing newline or CRLF from input.
func trimTrailingNewline(s string) string {
	s = strings.TrimSuffix(s, "\n")
	s = strings.TrimSuffix(s, "\r")
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil && len(line) == 0 {
		return
	}
	input := trimTrailingNewline(line)
	fmt.Println(isPalindrome(input))
}
