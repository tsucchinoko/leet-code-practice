package main

// アプローチ1
func removeElement(nums []int, val int) int {
	k := 0
	for _, num := range nums {
		if num != val {
			nums[k] = num
			k++
		}
	}
	return k
}

// アプローチ2
func removeElement2(nums []int, val int) int {
	n := len(nums)
	for i := 0; i < n; {
		if nums[i] == val {
			nums[i] = nums[n-1] // 末尾と交換
			n--                 // 有効範囲を縮める
		} else {
			i++
		}
	}
	return n
}
