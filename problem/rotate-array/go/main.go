package main

// rotate rotates nums to the right by k steps in-place.
func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	// k > n の可能性があるため
	k = k % n
	if k == 0 {
		return
	}
	reverse := func(a []int, i, j int) {
		for i < j {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}
	// reverse all
	reverse(nums, 0, n-1)
	// reverse first k
	reverse(nums, 0, k-1)
	// reverse remaining
	reverse(nums, k, n-1)
}
