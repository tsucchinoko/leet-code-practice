# maximum-average-subarray-1

## 問題の性質

- スライディングウィンドウ（固定長部分配列の和を効率的に管理する問題）
- 単純な線形走査で最適解が得られる（貪欲や二分探索は不要）
- 浮動小数点の出力許容誤差が 1e-5

## 問題要約

長さ n の整数配列 nums と整数 k が与えられる。連続する長さ k
の部分配列のうち、平均値が最大になるものを見つけ、その平均値を返す。答えは誤差
1e-5 以下であれば受理される。

例:

- nums = [1,12,-5,-6,50,3], k = 4 -> 最大平均は (12 + (-5) + (-6) + 50) / 4 =
  12.75
- nums = [5], k = 1 -> 5.0

## 制約

- n == nums.length
- 1 <= k <= n <= 10^5
- -10^4 <= nums[i] <= 10^4

時間計算量は O(n)、追加の空間は O(1) を目指すのが望ましい。

## 考え方

- 長さ k の連続部分配列の平均は、その部分配列の合計を k
  で割ったもの。したがって「平均が最大」は「合計が最大」と同値。
- 固定長 k
  の部分配列の合計を各位置で逐次計算するには、スライディングウィンドウ手法を使う：
  1. 最初の k 要素の合計 sum を初期化する。
  2. ウィンドウを右に1つずつスライドするとき、新しく入る要素を足し、出ていく要素を引くことで
     O(1) で合計を更新できる。
  3. 各位置での合計の最大値 maxSum を保持しておき、最終的に maxSum / k を返す。
- 負の値も含むが、合計の比較のみで良いため特別な扱いは不要。
- 浮動小数点は最終的な割り算でのみ使用する（合計は int64
  を使いオーバーフローを避ける）。n ≤ 1e5、|nums[i]| ≤ 1e4
  のため合計の絶対値は最大で 1e9 → int32でも足りるが安全のため int64 を使う。

## キーワード

- スライディングウィンドウ
- 部分配列（連続）
- 線形時間 O(n)
- int64 と浮動小数点変換

---

以下に Go 言語での実装（main.go）とテストコード（main_test.go）を示します。

main.go

```go
package main

import (
    "fmt"
)

func findMaxAverage(nums []int, k int) float64 {
    n := len(nums)
    if n == 0 || k == 0 {
        return 0.0
    }

    var sum int64 = 0
    for i := 0; i < k; i++ {
        sum += int64(nums[i])
    }
    maxSum := sum

    for i := k; i < n; i++ {
        sum += int64(nums[i])
        sum -= int64(nums[i-k])
        if sum > maxSum {
            maxSum = sum
        }
    }

    return float64(maxSum) / float64(k)
}

func main() {
    // 簡単な実行例
    fmt.Printf("%.5f\n", findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)) // 12.75000
}
```

main_test.go

```go
package main

import (
    "math"
    "testing"
)

func almostEqual(a, b, eps float64) bool {
    return math.Abs(a-b) <= eps
}

func TestFindMaxAverage(t *testing.T) {
    tests := []struct {
        nums []int
        k    int
        want float64
    }{
        {[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
        {[]int{5}, 1, 5.0},
        {[]int{0, 0, 0, 0}, 2, 0.0},
        {[]int{-1, -2, -3, -4}, 2, -1.5},
        {[]int{1, 2, 3, 4, 5}, 5, 3.0},
        {[]int{1, 2, 3, 4, 5}, 1, 5.0},
    }

    eps := 1e-5
    for _, tc := range tests {
        got := findMaxAverage(tc.nums, tc.k)
        if !almostEqual(got, tc.want, eps) {
            t.Errorf("findMaxAverage(%v, %d) = %v; want %v", tc.nums, tc.k, got, tc.want)
        }
    }
}
```

- テストは誤差 1e-5 を許容します。
- 実装は O(n) 時間、O(1) 追加空間です。
