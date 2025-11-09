package main

import (
	"fmt"
)

const (
	EMPTY   = 0
	PLANTED = 1
)

func canPlantAt(f []int, i int) bool {
	// すでに植えられている
	if f[i] != EMPTY {
		return false
	}
	// 一つ前が植えられている
	if i > 0 && f[i-1] != EMPTY {
		return false
	}

	// 一つ後ろが植えられている
	if i < len(f)-1 && f[i+1] != EMPTY {
		return false
	}

	// 両隣が空いている
	return true
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	for i, _ := range flowerbed {
		if !canPlantAt(flowerbed, i) {
			continue
		}
		flowerbed[i] = PLANTED
		n--
		if n == 0 {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1)) // true
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2)) // false
}
