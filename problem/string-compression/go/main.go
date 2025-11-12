package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	n := len(chars)
	write := 0
	i := 0

	for i < n {
		ch := chars[i]
		j := i
		for j < n && chars[j] == ch {
			j++
		}
		count := j - i

		// write chars in-place
		chars[write] = ch
		write++

		// if count > 1, write the digits of count
		if count > 1 {
			s := strconv.Itoa(count)
			for k := 0; k < len(s); k++ {
				chars[write] = s[k]
				write++
			}
		}

		// move to next group
		i = j
	}
	return write
}

func main() {
	// example run
	arr := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	newLen := compress(arr)
	fmt.Printf("newLen=%d, compressed=%q\n", newLen, arr[:newLen])
}
