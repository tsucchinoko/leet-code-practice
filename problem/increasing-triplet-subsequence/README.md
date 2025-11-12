# increasong-triplet-subsequence

## 問題の性質

増加する長さ3の部分列（インデックス順で i < j < k かつ nums[i] < nums[j] <
nums[k]）が配列に存在するかを判定する決定問題。部分列は連続である必要はない。入力長が大きく、時間・空間効率が重要。

## 問題要約

与えられた整数配列 nums
に対して、インデックスの増加順で要素が単調増加する長さ3の部分列（トリプル）が存在すれば
true、なければ false を返す。

例:

- [1,2,3,4,5] → true
- [5,4,3,2,1] → false
- [2,1,5,0,4,6] → true （0 < 4 < 6）

## 制約

- 1 <= nums.length <= 5 * 10^5
- -2^31 <= nums[i] <= 2^31 - 1
- 要求: 可能なら O(n) 時間、O(1) 追加空間で解く（Follow-up）

## 考え方

アルゴリズムの要点（簡潔に）:

- 長さ3の増加部分列があるかを線形走査で判定できる。
- 走査中に「最小値 small」と「small より大きいが今のところの第2候補
  mid」を保持する。
- イデア:
  1. small を非常に大きな値で初期化、mid も非常に大きな値で初期化する。
  2. 配列を左から右へ走査する。
  3. 現在値 x を見て:
     - x <= small のとき: small = x（より小さな最小値を更新）
     - else if x <= mid のとき: mid = x（small < x <= mid、2番目候補を更新）
     - else: small < mid < x の状況で、長さ3の増加部分列を見つけたので true
       を返す
- なぜこれで良いか:
  - small は左側での最小の候補（i）を保持する。mid は small より大きく、i
    より右にある j の最小候補を保持する。
  - もし x > mid が現れれば、small < mid < x
    の順序でインデックス関係も保持される（走査順により）。
- 証明の概略:
  - small, mid は走査を通じて左からの有効な候補を維持する。mid が存在する時点で
    small は mid の左にあり small < mid。以降で mid より大きい x
    が出ればインデックス順も保証される。

計算量:

- 時間: O(n)
- 追加空間: O(1)

注意点:

- 等号の扱いに注意（増加は厳密な <）。したがって x == small は small
  を更新しても mid を更新しない、x == mid も mid を更新する処理で問題ない（<=
  を使うことで安定に動作する）。

## キーワード

- 貪欲法（Greedy）
- 線形走査（One-pass）
- 定数空間（O(1) space）
- 増加部分列（Increasing subsequence）
- 334 (LeetCode)

---

以下に Go 言語での具体的実装を示します。ファイルは main.go と main_test.go
です。

main.go

```go
package main

import "fmt"

// increasingTriplet returns true if there exists i < j < k such that nums[i] < nums[j] < nums[k].
func increasingTriplet(nums []int) bool {
    if len(nums) < 3 {
        return false
    }

    // Initialize small and mid to large values
    const INF = int(^uint(0) >> 1) // max int
    small, mid := INF, INF

    for _, x := range nums {
        if x <= small {
            small = x
        } else if x <= mid {
            mid = x
        } else {
            // small < mid < x
            return true
        }
    }
    return false
}

func main() {
    tests := [][]int{
        {1, 2, 3, 4, 5},
        {5, 4, 3, 2, 1},
        {2, 1, 5, 0, 4, 6},
        {2, 2, 2, 2},
        {1, 1, 2, 2, 3},
    }
    for _, t := range tests {
        fmt.Println(t, "=>", increasingTriplet(t))
    }
}
```

main_test.go

```go
package main

import "testing"

func TestIncreasingTriplet(t *testing.T) {
    cases := []struct {
        nums []int
        want bool
    }{
        {[]int{1, 2, 3, 4, 5}, true},
        {[]int{5, 4, 3, 2, 1}, false},
        {[]int{2, 1, 5, 0, 4, 6}, true},
        {[]int{2, 2, 2, 2}, false},
        {[]int{1, 1, 2, 2, 3}, true},
        {[]int{1, 3, 2}, false},
        {[]int{1, 2}, false},
        {[]int{}, false},
        {[]int{0, -1, 2, -2, 3}, true},
    }

    for _, c := range cases {
        got := increasingTriplet(c.nums)
        if got != c.want {
            t.Errorf("increasingTriplet(%v) = %v; want %v", c.nums, got, c.want)
        }
    }
}
```

上記実装は O(n) 時間、O(1) 追加空間で動作します。
