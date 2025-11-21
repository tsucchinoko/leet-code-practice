package main

import "fmt"

func largestAltitude(gain []int) int {
	maxAltitude, currentAltitude := 0, 0
	for _, altitudeGain := range gain {
		currentAltitude += altitudeGain
		if currentAltitude > maxAltitude {
			maxAltitude = currentAltitude
		}
	}
	return maxAltitude
}

func main() {
	// Example usage
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))         // 1
	fmt.Println(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2})) // 0
}
