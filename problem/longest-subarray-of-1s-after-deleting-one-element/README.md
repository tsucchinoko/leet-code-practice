# longest-subarray-of-1s-after-deleting-one-element

## 問題の性質

- スライディングウィンドウ（連続部分列）または双方向スキャンで解ける配列操作問題。
- 目的は「1
  をなるべく長く連続させる」ために配列からちょうど1つの要素を削除すること。
- 線形時間（O(n)）・定数追加空間（O(1)）で解ける典型的な中級問題。

## 問題要約

長さ n の二値配列 nums
が与えられる。配列からちょうど1つの要素を削除した後に得られる「1
のみからなる連続部分配列」の最大長を返す。 ただし、結果のサブアレイが空（1
が一つも無い）なら 0 を返す。

例：

- nums = [1,1,0,1] → 削除後に [1,1,1] を作れて出力 3
- nums = [1,1,1] → 必ず1つ削除するため出力 2

## 制約

- 1 <= nums.length <= 10^5
- nums[i] は 0 または 1
- 必ず1要素を削除する（配列長が1でも削除が必要になる状況を考慮）

## 考え方

アルゴリズムの直感（簡潔に）：

- 削除できるのは任意の1要素（0 でも 1 でも可）。0 を削除すると左右の 1
  の塊をつなげられる可能性がある。1 を削除すると単にその1を減らすだけ。
- 各位置の 0 を「境界」として、その左右に伸びる連続した 1
  の長さを足し合わせれば（左右の 1 の塊をつなげて）その 0
  を削除したときの得られる最大長になる。
- ただし配列が全て1のときは必ず1つ削除しなければならないので、結果は n-1。
- 線形スキャンで「直前の連続する1の長さ」と「現在の連続する1の長さ（現在連続の1ブロックの長さ）」を追跡し、0
  に遭遇したら前後の長さを合算して最大値を更新する。

具体的には：

- prev = 長さが直前に存在した（直前のブロック）1の長さ（0 を挟んだ前のブロック）
- curr = 現在走査中のブロックの 1 の長さ
- maxLen = 0
- 数列を左から走査：
  1. num が 1 なら curr++。
  2. num が 0 なら：
     - maxLen = max(maxLen, prev + curr) — この 0 を削除してつなげたときの長さ
     - prev = curr （現在のブロックが次の prev になる）
     - curr = 0
- 走査終了後に最後のブロックを考慮して maxLen = max(maxLen, prev + curr)
- ただし配列が全て1のときは結果を n-1 にする（上の合算では prev は 0 で curr = n
  となり prev+curr = n となるが削除必須なので n-1 を返す）
- また 1 が一つも無い（全て0）のときは prev = curr = 0 となるので 0
  を返す（削除しても1は出現しない）

計算量：

- 時間 O(n)、追加空間 O(1)

## キーワード

- スライディングウィンドウ
- 線形走査
- 連続部分列
- 定数空間

---

以下、Go 言語での実装とテストコードを示します。

ファイル: main.go

```go
package main

import "fmt"

func longestSubarray(nums []int) int {
    n := len(nums)
    prev, curr := 0, 0
    maxLen := 0
    for _, v := range nums {
        if v == 1 {
            curr++
        } else {
            if prev+curr > maxLen {
                maxLen = prev + curr
            }
            prev = curr
            curr = 0
        }
    }
    // 最後のブロックを考慮
    if prev+curr > maxLen {
        maxLen = prev + curr
    }
    // 全て1の場合は必ず1つ削除するので n-1
    if maxLen == n { // 全て1のときだけ成立する
        return n - 1
    }
    return maxLen
}

func main() {
    fmt.Println(longestSubarray([]int{1, 1, 0, 1}))                   // 3
    fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))   // 5
    fmt.Println(longestSubarray([]int{1, 1, 1}))                     // 2
}
```

ファイル: main_test.go

```go
package main

import "testing"

func TestLongestSubarray(t *testing.T) {
    tests := []struct {
        nums []int
        want int
    }{
        {[]int{1, 1, 0, 1}, 3},
        {[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
        {[]int{1, 1, 1}, 2},
        {[]int{0, 0, 0}, 0},
        {[]int{1}, 0},           // 1要素で1を削除すると空なので0
        {[]int{0}, 0},
        {[]int{1, 0, 1, 1, 0, 1, 1, 1, 0, 1}, 5},
        {[]int{1, 1, 0, 0, 1, 1, 1}, 4},
    }

    for _, tc := range tests {
        got := longestSubarray(tc.nums)
        if got != tc.want {
            t.Fatalf("nums=%v: got=%d want=%d", tc.nums, got, tc.want)
        }
    }
}
```
