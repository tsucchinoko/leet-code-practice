package main

// カウント（バケット）法による h-index 計算（線形時間）
func hIndex(citations []int) int {
	n := len(citations)
	counts := make([]int, n+1)
	for _, c := range citations {
		if c >= n {
			counts[n]++
		} else {
			counts[c]++
		}
	}
	total := 0
	for h := n; h >= 0; h-- {
		total += counts[h]
		if total >= h {
			return h
		}
	}
	return 0
}

// 降順ソート法による h-index 計算
func hIndex2(citations []int) int {
	a := make([]int, len(citations))
	copy(a, citations)
	// sort descending
	for i := range a {
		for j := i + 1; j < len(a); j++ {
			if a[j] > a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	h := 0
	for i, v := range a {
		if v >= i+1 {
			h = i + 1
		} else {
			break
		}
	}
	return h
}
