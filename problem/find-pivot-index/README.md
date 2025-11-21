# find-pivot-index

## 問題の性質

配列のあるインデックス i に対して、その左側（i より前）の要素の合計と右側（i
より後ろ）の要素の合計が等しくなるようなインデックス（ピボット）を探す問題。左端や右端も候補となりうる（その場合は存在しない側の和は
0 とみなす）。見つかれば最も左側のインデックスを返し、なければ -1 を返す。

## 問題要約

与えられた整数配列 nums
に対して、左側の合計と右側の合計が等しくなる最左のインデックス（pivot
index）を返す。存在しなければ -1。

例:

- nums = [1,7,3,6,5,6] -> 出力 3
- nums = [1,2,3] -> 出力 -1
- nums = [2,1,-1] -> 出力 0

## 制約

- 1 <= nums.length <= 10^4
- -1000 <= nums[i] <= 1000
- 計算上、合計は最大で 10^4 * 1000 = 10^7 程度（int 型で問題なし）

（LeetCode の同等問題: 1991）

## 考え方

1. 総和 total を最初に計算する。
2. 配列を左から走査し、現在の要素より左側の和を leftSum として管理する。
3. インデックス i で右側の和は total - leftSum - nums[i] で表せる。 もし leftSum
   == total - leftSum - nums[i] なら i がピボット。
4. 条件を満たしたらその最初の i を返す。最後まで見つからなければ -1。
5. 時間計算量 O(n)、追加の空間 O(1)。

理由（簡潔）:

- 総和を使うことで i
  ごとに右側の和を都度再計算する必要がなく、各要素を1回だけ見ればよい。

## キーワード

- 総和（prefix / suffix sums）
- 1回走査（single pass）
- 定数空間
- 左右の和の比較

---

以下に Go 実装とテストコードを示します。ファイル名は指定どおりです。

main.go

```go
package main

import "fmt"

func pivotIndex(nums []int) int {
    total := 0
    for _, v := range nums {
        total += v
    }

    leftSum := 0
    for i, v := range nums {
        if leftSum == total-leftSum-v {
            return i
        }
        leftSum += v
    }
    return -1
}

func main() {
    examples := [][]int{
        {1, 7, 3, 6, 5, 6},
        {1, 2, 3},
        {2, 1, -1},
    }
    for _, ex := range examples {
        fmt.Println(pivotIndex(ex))
    }
}
```

main_test.go

```go
package main

import "testing"

func TestPivotIndex(t *testing.T) {
    tests := []struct {
        nums []int
        want int
    }{
        {[]int{1, 7, 3, 6, 5, 6}, 3},
        {[]int{1, 2, 3}, -1},
        {[]int{2, 1, -1}, 0},
        {[]int{0}, 0},
        {[]int{-1,-1,-1,0,1,1}, 0},
        {[]int{1,-1,0}, 2},
        {[]int{1, -1, 1, -1, 1, -1}, 0}, // all prefixes match at 0
    }

    for _, tt := range tests {
        got := pivotIndex(tt.nums)
        if got != tt.want {
            t.Fatalf("pivotIndex(%v) = %d; want %d", tt.nums, got, tt.want)
        }
    }
}
```
