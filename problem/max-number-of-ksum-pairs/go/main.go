package main

import "fmt"

func maxOperations(nums []int, k int) int {
	// 各要素の出現回数をカウント
	freq := make(map[int]int)
	ops := 0
	// nums を順に見て、補数が既に freq にあれば直ちにペアを作る。
	// そうでなければ自分を freq に登録して将来の補数待ちにする。
	for _, x := range nums {
		complement := k - x
		if freq[complement] > 0 {
			// 補数があればペアを作る（補数側のカウントを減らす）
			ops++
			freq[complement]--
		} else {
			// 補数が無いので自分のカウントを増やす
			freq[x]++
		}
	}

	return ops
}

func main() {
	nums := []int{1, 2, 3, 4}
	k := 5
	fmt.Println(maxOperations(nums, k))
}
