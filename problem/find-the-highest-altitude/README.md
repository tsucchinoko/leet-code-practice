# find-the-highest-altitude

## 問題の性質

配列操作、累積和（prefix
sum）、線形走査。各要素は隣接する2点間の高度差を表すので、出発点高度0から順に累積していき最高値を求めるだけの単純な問題。時間計算量は
O(n)、追加の空間は O(1) で解ける。

## 問題要約

長さ n の整数配列 gain が与えられる。gain[i] は点 i から点 i+1
への高度の増分（正なら上昇、負なら下降）を表す。点0の高度は 0。全ての点（0 から
n まで、合計 n+1 点）の高度を累積して求めたときの最高高度を返す。

例:

- gain = [-5,1,5,0,-7] のとき高度列は [0,-5,-4,1,1,-6] で最高高度は 1。

## 制約

- n == gain.length
- 1 <= n <= 100
- -100 <= gain[i] <= 100

制約より、オーバーフロー等の心配は不要で、単純な線形走査で十分高速に動作する。

## 考え方

1. 変数 current を 0 で初期化し、出発点の高度を表す。
2. 変数 maxAlt を 0 で初期化し、これが最高高度となる。
3. gain 配列を先頭から順に走査する。
   1. current に gain[i] を加える（次の点の高度）。
   2. current が maxAlt より大きければ maxAlt を更新する。
4. 走査終了後に maxAlt を返す。

この問題は累積和をとって最大値を求めるだけなので、実装は非常に簡潔。初期値として
0 を忘れないこと（出発点の高度を考慮する必要がある）。

## キーワード

- 累積和（prefix sum）
- 線形走査（single pass）
- 最大値更新（running maximum）

---

以下に Go 言語での実装（main.go）とテストコード（main_test.go）を示します。

main.go

```go
package main

import "fmt"

// highestAltitude returns the highest altitude reached given gain array.
func highestAltitude(gain []int) int {
    current := 0
    maxAlt := 0
    for _, g := range gain {
        current += g
        if current > maxAlt {
            maxAlt = current
        }
    }
    return maxAlt
}

func main() {
    // Example usage
    fmt.Println(highestAltitude([]int{-5, 1, 5, 0, -7})) // 1
    fmt.Println(highestAltitude([]int{-4, -3, -2, -1, 4, 3, 2})) // 0
}
```

main_test.go

```go
package main

import "testing"

func TestHighestAltitude(t *testing.T) {
    tests := []struct {
        gain []int
        want int
    }{
        {[]int{-5, 1, 5, 0, -7}, 1},
        {[]int{-4, -3, -2, -1, 4, 3, 2}, 0},
        {[]int{1,2,3,4,5}, 15},      // monotonically increasing
        {[]int{-1,-2,-3,-4,-5}, 0},  // monotonically decreasing, start is max
        {[]int{0,0,0,0}, 0},         // all zeros
        {[]int{5,-2,3,-1,2}, 7},     // mixed
    }

    for _, tt := range tests {
        got := highestAltitude(tt.gain)
        if got != tt.want {
            t.Fatalf("highestAltitude(%v) = %d; want %d", tt.gain, got, tt.want)
        }
    }
}
```
