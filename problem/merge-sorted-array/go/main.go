package main

// merge は 2つの昇順配列 nums1 と nums2 をマージし、結果を nums1 に格納します。
//
// 引数:
//   - nums1: 長さ m+n のスライス。先頭 m 個が有効値で、末尾 n 個は 0（空きスペース）
//   - m: nums1 の有効な要素数
//   - nums2: 長さ n のスライス。すべて有効値
//   - n: nums2 の要素数
//
// 備考:
//   - あらかじめ確保された値を入れ替えるだけなので、メモリアロケーションは発生しない
//     -
func merge(nums1 []int, m int, nums2 []int, n int) {
	// nums1の有効末尾インデックス, nums2の末尾インデックス, nums1の全長末尾インデックス
	i, j, k := m-1, n-1, m+n-1

	// どちらかの配列が空になるまで
	for i >= 0 && j >= 0 {
		// 2つの配列を比較し、大きい方をnums1に詰める
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}

		k--
	}

	// nums2に要素が残っている場合にはnums1に詰める
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
