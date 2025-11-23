# asteroid-collision

## 問題の性質

- シミュレーション / スタックを使う典型的な問題
- 一方向（右向き）の正の値と左向きの負の値が衝突する場合のみ相互作用が発生する（その他は非干渉）
- 衝突は局所的かつ逐次的に解決でき、全体のシーケンス順序で決定可能

## 問題要約

- 配列 asteroids
  が与えられる。各要素は絶対値が大きさ、符号が方向（正＝右、負＝左）を示す。
- 各小惑星は同じ速度で移動するため、右向きの小惑星とその右側にある左向きの小惑星のみが衝突する可能性がある。
- 衝突時、より小さな小惑星は爆発して消滅する。大きさが同じなら両方とも消滅する。
- 最終的に残る小惑星の配列を返す。

## 制約

- 2 <= asteroids.length <= 10^4
- -1000 <= asteroids[i] <= 1000
- asteroids[i] != 0
- 入力サイズは最大で 10,000 要素なので、O(n) または O(n log n)
  のアルゴリズムが望ましい。メモリは出力配列分を含めて追加で O(n)
  程度まで許容可能。

## 考え方

- 衝突の条件：右向きの小惑星（正）とその右にいる左向きの小惑星（負）が出会う場合に発生する。配列を左から右へ走査すると、過去に出現した右向きの小惑星（スタックに保持）と現在の左向き小惑星との衝突を検討すればよい。
- スタックを使う方法（標準解法）：
  1. 空のスタックを用意する。
  2. 左から右へ asteroids を順に見る。現在の値 cur を考える。
  3. cur が正（右向き）の場合は衝突は起きないのでスタックに push。
  4. cur
     が負（左向き）の場合、スタックのトップに正の小惑星がある限り衝突が起きうるので以下を繰り返す：
     1. スタックが空、またはスタックトップが負（左向き）であれば衝突は起きないので
        cur をスタックに push してループを抜ける。
     2. スタックトップの絶対値が小さければ（|top| <
        |cur|）スタックトップが消える（pop）し、まだ衝突の可能性があるので繰り返す。
     3. 同じ大きさなら（|top| == |cur|）スタックトップを pop して cur
        も消える（push せずにループを抜ける）。
     4. スタックトップの方が大きければ（|top| > |cur|）cur が消える（push
        せずにループを抜ける）。
- この方法で各小惑星は最大1回 push と最大1回 pop されるため時間計算量は
  O(n)、追加メモリは O(n)（スタック）となる。

## キーワード

- スタック（Stack）
- シミュレーション
- 衝突判定（局所的）
- O(n) 時間

---

以下に Go 言語での具体的実装を示します。ファイル分割は要求どおり main.go と
main_test.go にしています。

main.go

```go
package main

import (
    "fmt"
)

func asteroidCollision(asteroids []int) []int {
    stack := make([]int, 0, len(asteroids))
    for _, cur := range asteroids {
        collided := false
        for len(stack) > 0 && stack[len(stack)-1] > 0 && cur < 0 {
            top := stack[len(stack)-1]
            if abs(top) < abs(cur) {
                // stack top explodes
                stack = stack[:len(stack)-1]
                // continue checking with next top
                continue
            } else if abs(top) == abs(cur) {
                // both explode
                stack = stack[:len(stack)-1]
                collided = true
                break
            } else {
                // cur explodes
                collided = true
                break
            }
        }
        if !collided {
            stack = append(stack, cur)
        }
    }
    return stack
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    examples := [][]int{
        {5, 10, -5},
        {8, -8},
        {10, 2, -5},
        {3, 5, -6, 2, -1, 4},
    }
    for _, ex := range examples {
        fmt.Printf("input: %v -> output: %v\n", ex, asteroidCollision(ex))
    }
}
```

main_test.go

```go
package main

import (
    "reflect"
    "testing"
)

func TestAsteroidCollision(t *testing.T) {


    tests := []struct {
        in  []int
        out []int
    }{
        {[]int{5, 10, -5}, []int{5, 10}},
        {[]int{8, -8}, []int{}},
        {[]int{10, 2, -5}, []int{10}},
        {[]int{3, 5, -6, 2, -1, 4}, []int{-6, 2, 4}},
        {[]int{-2, -1, 1, 2}, []int{-2, -1, 1, 2}}, // no collisions (same directions or separated)
        {[]int{1, -2, -2, -2}, []int{-2, -2, -2}},
        {[]int{2, 1, -1, -2}, []int{}},
    }

    for _, tc := range tests {
        got := asteroidCollision(tc.in)
        if !reflect.DeepEqual(got, tc.out) {
            t.Fatalf("input %v: expected %v, got %v", tc.in, tc.out, got)
        }
    }
}
```
