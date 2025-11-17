# container-with-most-water

## 問題の性質

この問題は「配列の両端を選んで作る長方形（幅 ×
高さ）」の面積の最大化を求める最適化問題です。探索空間は二つのインデックスの組み合わせで二次的に増えるが、特定の単調性を利用して線形時間で解けます。貪欲的かつ二分探索的な双方向ポインタ（two
pointers）戦略が適用される典型的な問題です。

## 問題要約

長さ n の整数配列 height が与えられる。各 i（0-indexed とする）について、点
(i, 0) と (i, height[i]) を結ぶ垂直線がある。2 本の線 i, j（i <
j）を選び、それらと x
軸で囲まれる長方形（コンテナ）が蓄えられる水量を考える。水量（面積）は幅 (j - i)
と高さ min(height[i], height[j])
の積である。傾けることはできない。可能な組み合わせの中で最大の水量を返せ。

例:

- height = [1,8,6,2,5,4,8,3,7] → 出力 49
- height = [1,1] → 出力 1

## 制約

- n == height.length
- 2 <= n <= 10^5
- 0 <= height[i] <= 10^4 （時間・メモリの厳密制限は明示されていないが、n が最大
  10^5 のため O(n^2) の解法は現実的でない。O(n) または O(n log n)
  を目指すべき。）

## 考え方

1. 面積の定義: 2 本の線 i < j による水量は \[ \text{area} = (j - i) \times
   \min(\text{height}[i], \text{height}[j]) \]
2. 全組み合わせを試すと O(n^2) だが、双方向ポインタ法で O(n) に削減可能:
   - 左ポインタ l を 0、右ポインタ r を n-1 に初期化する。
   - 現在の面積を計算して最大値を更新する。
   - 次にどちらのポインタを動かすかを決める。高さが低い側のポインタを内側へ移動するのが正しい決定：
     - 理由（直感）:
       面積は幅と高さの積。幅を狭めることで面積が増えるためには、縮めた後の高さが現在の
       min を上回らなければならない。現在 min
       をつくっている側（低い側）を動かして高い線を見つける可能性があるので、低い側を動かすのが唯一の有望な操作。高い側を動かしても
       min は変わらないか減るだけなので、幅が減ることで面積は改善しない。
   - 低い側のポインタを 1 つ内側に移動して探索を続ける。これを l >= r
     になるまで繰り返す。
3. この戦略は常に最適解を見つける（幅を縮める際に低い側を動かすことが安全であり、より良い解が存在するならその操作で見つかれる）ため、時間計算量は
   O(n)、空間計算量は O(1)。

## キーワード

- two pointers（双方向ポインタ）
- greedy（貪欲）
- O(n) 時間
- 面積最適化
- 配列インデックス

---

以下に Go 言語での実装を示します。実装ファイルは main.go、テストファイルは
main_test.go としてあります。テストは標準的なテストフレームワーク testing
を使い、例題といくつかの追加ケースを検証します。

main.go

```go
package main

import (
    "fmt"
)

func maxArea(height []int) int {
    n := len(height)
    if n < 2 {
        return 0
    }
    l, r := 0, n-1
    maxA := 0
    for l < r {
        h := height[l]
        if height[r] < h {
            h = height[r]
        }
        area := (r - l) * h
        if area > maxA {
            maxA = area
        }
        // move the smaller height pointer inward
        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }
    return maxA
}

func main() {
    examples := [][]int{
        {1, 8, 6, 2, 5, 4, 8, 3, 7},
        {1, 1},
    }
    for _, ex := range examples {
        fmt.Printf("height = %v -> maxArea = %d\n", ex, maxArea(ex))
    }
}
```

main_test.go

```go
package main

import "testing"

func TestMaxAreaExamples(t *testing.T) {
    tests := []struct {
        height []int
        want   int
    }{
        {[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
        {[]int{1, 1}, 1},
        {[]int{4, 3, 2, 1, 4}, 16},     // both ends
        {[]int{1, 2, 1}, 2},            // middle not best
        {[]int{0, 0, 0, 0}, 0},         // zeros
        {[]int{10000, 10000}, 10000},   // large values
        {[]int{1, 2, 4, 3}, 4},
    }

    for _, tt := range tests {
        got := maxArea(tt.height)
        if got != tt.want {
            t.Fatalf("maxArea(%v) = %d; want %d", tt.height, got, tt.want)
        }
    }
}
```

使用方法:

- main.go を実行すると例題の出力を表示します。
- go test でテストを実行してください。
