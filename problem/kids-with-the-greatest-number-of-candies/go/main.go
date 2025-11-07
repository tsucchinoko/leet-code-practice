package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	if len(candies) == 0 {
		return []bool{}
	}

	maxCandies := candies[0]
	for _, v := range candies {
		if v > maxCandies {
			maxCandies = v
		}
	}

	res := make([]bool, len(candies))
	for i, v := range candies {
		res[i] = (v + extraCandies) >= maxCandies
	}
	return res
}

func main() {
	// 簡単なデモ
	candies := []int{2, 3, 5, 1, 3}
	extra := 3
	fmt.Println(kidsWithCandies(candies, extra)) // [true true true false true]
}
