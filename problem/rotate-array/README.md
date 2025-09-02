# rotate-array

## 問題要約

整数配列 `nums` が与えられる。配列を右に `k` ステップ回転（シフト）せよ。 右に 1
ステップ回転するとは、配列の各要素がインデックスを 1
つ増やし、最後の要素が先頭に来る操作を指す。`k` は非負整数。

例:

- 入力: `nums = [1,2,3,4,5,6,7]`, `k = 3` 出力: `[5,6,7,1,2,3,4]`
- 入力: `nums = [-1,-100,3,99]`, `k = 2` 出力: `[3,99,-1,-100]`

目的は配列を指定の回数だけ右に回転した結果を返すこと（できればインプレースで、追加メモリ
O(1) を使って実装）。

## 制約

- 1 <= nums.length <= 10^5
- -2^31 <= nums[i] <= 2^31 - 1
- 0 <= k <= 10^5

時間・空間に関する期待:

- 可能ならば O(n) 時間、O(1) 追加空間での解法を目指す。

## 考え方

ここでは代表的なアプローチを複数紹介し、最後に Go の実装で O(1)
追加空間の方法（反転法）を提示する。

1. 余剰を使った簡単な方法（追加配列を使う）

- k が配列長 n を越える場合があるので、まず `k = k % n` とする。
- 新しい配列 `b` を長さ n で用意し、各要素を `b[(i+k)%n] = nums[i]` と配置する。
- `b` を `nums` にコピーして終了。
- 時間 O(n)、追加空間 O(n)（簡潔で実装容易）。

2. 要素を逐次回転させる（ループで1ずつ回す）

- 右に1回回転する操作を k 回繰り返す方法。
- 各回転は O(n) 時間なので合計 O(n*k) となり、k が大きいと非現実的。

3. 循環置換を使う方法（サイクル処理）

- ある位置から始め、要素を適切な位置へ移動させていくことで追加配列を使わずに済む。
- 各要素は一度だけ移動され、時間 O(n)、追加空間 O(1)
  となるが、実装はやや複雑で、移動済み判定とサイクル検出が必要。

4. 反転（reverse）を使う方法 — 推奨（シンプルで O(1) 追加空間）

- 手順（n = len(nums), k %= n）:

1. 配列全体を反転する
2. 先頭から k 要素を反転する
3. 残りの n-k 要素を反転する

- なぜ動くかの直感:
- 例えば nums = [1,2,3,4,5,6,7], k=3:
  - 全体反転 -> [7,6,5,4,3,2,1]
  - 先頭 k=3 を反転 -> [5,6,7,4,3,2,1]
  - 残りを反転 -> [5,6,7,1,2,3,4]
- 各反転は O(n) 全体で O(n)、追加配列を使わないので追加空間 O(1)。
- 実装は `reverse(nums, start, end)` 補助関数を用意して三度呼ぶだけで済む。

次に、Go 言語での具体的な実装とテストコードを示します。

```go
package main

import (
	"fmt"
	"reflect"
)

// rotate rotates nums to the right by k steps in-place.
func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
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

func main() {
	// テストケース
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{1}, 0, []int{1}},
		{[]int{1, 2}, 3, []int{2, 1}}, // k > n
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}}, // k == n
	}

	for i, tc := range tests {
		// コピーして関数を呼ぶ（破壊的なので）
		arr := make([]int, len(tc.nums))
		copy(arr, tc.nums)
		rotate(arr, tc.k)
		if !reflect.DeepEqual(arr, tc.expected) {
			fmt.Printf("Test %d FAILED: input=%v k=%d got=%v expected=%v\n", i+1, tc.nums, tc.k, arr, tc.expected)
		} else {
			fmt.Printf("Test %d PASSED: input=%v k=%d => %v\n", i+1, tc.nums, tc.k, arr)
		}
	}
}
```

ポイントまとめ:

- 反転（reverse）法は実装が簡潔でバグが少なく、時間 O(n)、追加空間 O(1)
  を達成できる。
- `k` を `k % n` にしておくことを忘れない（k が配列長を超える場合に備える）。
