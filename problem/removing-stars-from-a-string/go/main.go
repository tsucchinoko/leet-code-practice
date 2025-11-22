package main

func removeStars(s string) string {
	stack := make([]byte, 0, len(s))
	for i := range s {
		if s[i] == '*' {
			// pop
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			// push
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
