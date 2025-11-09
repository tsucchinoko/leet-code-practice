package main

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		flowerbed []int
		n         int
		want      bool
	}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
		{[]int{0}, 1, true},
		{[]int{0, 0}, 1, true},
		{[]int{0, 0, 0}, 2, true},
		{[]int{1, 0, 0, 0, 0, 1}, 2, false},
		{[]int{0, 0, 1, 0, 0}, 2, true},
		{[]int{1, 0, 1, 0, 1, 0, 1}, 1, false},
		{[]int{0, 0, 0, 0, 0}, 3, true}, // 植えられる最大は3 (positions 0,2,4)
		{[]int{0, 0, 0, 0, 0}, 4, false},
		{[]int{1, 0, 0, 0, 0}, 2, true}, // positions 2 and 4 (but 4 is edge)
		{[]int{1, 0, 0, 0, 0}, 3, false},
	}

	for i, tt := range tests {
		// テスト関数はスライスを破壊するのでコピーして渡す
		bedCopy := make([]int, len(tt.flowerbed))
		copy(bedCopy, tt.flowerbed)
		got := canPlaceFlowers(bedCopy, tt.n)
		if got != tt.want {
			t.Fatalf("case %d: flowerbed=%v n=%d want=%v got=%v", i, tt.flowerbed, tt.n, tt.want, got)
		}
	}
}
