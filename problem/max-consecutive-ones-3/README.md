# max-consecutive-ones-3

## 問題の性質

スライディングウィンドウ（双方向ポインタ）を使った最長部分配列探索。連続する1の最大長を求めるために0を最大k個まで反転（0→1）できるという条件で、可変長のウィンドウを右に伸ばしつつ左端を調整して条件を満たす最長幅を求める問題。線形時間で解ける。

## 問題要約

長さ n の二値配列 nums と整数 k
が与えられる。配列中の連続する1の長さを最大化するために、最大で k 個の 0 を 1
に反転できる。反転操作は任意の位置の0を1にすることを意味する（部分配列内の0を数えて、それが
k
を超えないようにする）。反転を考慮したときに得られる、連続する1（すなわち部分配列の長さ）の最大値を返す。

## 制約

- 1 <= nums.length <= 10^5
- nums[i] は 0 か 1
- 0 <= k <= nums.length
- 時間計算量は O(n) を目指すのが現実的（n は配列長）
- 追加の空間は O(1) が可能（出力以外の定数追加メモリ）

## 考え方

- 目標：連続する1（反転で得られる1を含む）の最長部分列長を求める。
- スライディングウィンドウアプローチ：
  1. 左端 l = 0、右端 r を 0 から n-1 まで動かす（右に伸ばす）。
  2. ウィンドウ [l, r] 内で 0 の個数 zeroCount を保持する。
  3. r を進めていき、nums[r] が 0 なら zeroCount++。
  4. zeroCount が k を超えたら、左端 l を右に動かして zeroCount
     を減らす（nums[l] が 0 なら zeroCount--）、zeroCount <= k になるまで l
     を進める。
  5. 各ステップでウィンドウ幅 r - l + 1 を計算し、最大値を更新する。
- 正当性：
  - 任意の最適解はあるウィンドウ [l*, r*]
    に対応する。スライディングウィンドウはそのようなウィンドウを探索し、zeroCount
    が k を超えない最大幅を記録するため最適解に到達する。
- 計算量：
  - 左右のポインタはそれぞれ最長 n 回移動するため O(n)。
  - 使用メモリは定数（カウンタとインデックスのみ）。
- 特記事項：
  - k が 0
    のときは連続する1の最大長を求める単純な問題になるが、同じアルゴリズムで扱える。
  - 入力が大きい場合でも高速に動作する。

## キーワード

- スライディングウィンドウ
- 双方向ポインタ（two pointers）
- 連続部分配列（連続部分列）
- 線形時間 O(n)
- 定数空間 O(1)

---

以下に Go 言語での実装（main.go）とテストコード（main_test.go）を示します。

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func longestOnes(nums []int, k int) int {
    n := len(nums)
    left := 0
    zeroCount := 0
    maxLen := 0
    for right := 0; right < n; right++ {
        if nums[right] == 0 {
            zeroCount++
        }
        for zeroCount > k {
            if nums[left] == 0 {
                zeroCount--
            }
            left++
        }
        // window [left, right] is valid
        if cur := right - left + 1; cur > maxLen {
            maxLen = cur
        }
    }
    return maxLen
}

func main() {
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

    // 簡易入出力: 1行目に n と k、2行目にスペース区切りの nums を期待する。
    var n, k int
    if _, err := fmt.Fscan(in, &n, &k); err != nil {
        fmt.Fprintln(out, "input error")
        return
    }
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(in, &nums[i])
    }
    fmt.Fprintln(out, longestOnes(nums, k))
}
```

```go
package main

import (
    "reflect"
    "testing"
)

func TestLongestOnes(t *testing.T) {
    tests := []struct {
        nums []int
        k    int
        want int
    }{
        {[]int{1,1,1,0,0,0,1,1,1,1,0}, 2, 6},
        {[]int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}, 3, 10},
        {[]int{1,1,1,1}, 0, 4},
        {[]int{0,0,0,0}, 2, 2},
        {[]int{0,1,0,1,1,0,1}, 1, 4},
        {[]int{1}, 0, 1},
        {[]int{0}, 0, 0},
        {[]int{0}, 1, 1},
    }

    for _, tt := range tests {
        got := longestOnes(tt.nums, tt.k)
        if !reflect.DeepEqual(got, tt.want) {
            t.Errorf("longestOnes(%v, %d) = %d; want %d", tt.nums, tt.k, got, tt.want)
        }
    }
}
```
