package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))
	for _, cur := range asteroids {
		collided := false

		for len(stack) > 0 && stack[len(stack)-1] > 0 && cur < 0 {
			top := stack[len(stack)-1]
			if abs(top) < abs(cur) {
				// stack top exploded
				stack = stack[:len(stack)-1]
				// continue checking next asteroid
				continue
			} else if abs(top) == abs(cur) {
				// both asteroids exploded
				stack = stack[:len(stack)-1]
				collided = true
				break
			} else {
				// current asteroid exploded
				collided = true
				break
			}
		}

		if !collided {
			stack = append(stack, cur)
		}
	}

	return stack
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	examples := [][]int{
		{5, 10, -5},
		{8, -8},
		{10, 2, -5},
		{3, 5, -6, 2, -1, 4},
	}
	for _, ex := range examples {
		fmt.Printf("input: %v -> output: %v\n", ex, asteroidCollision(ex))
	}
}
