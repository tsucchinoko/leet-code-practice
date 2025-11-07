# kids-with-the-greatest-number-of-candies

## 問題の性質

配列操作・比較（シミュレーション）。各要素に同じ値を加えたとき、その要素が配列内の最大値以上になるかを判定する単純な線形走査問題。計算量は入力長に対して線形で解ける（O(n)）。安定・決定的問題で特殊なデータ構造は不要。

## 問題要約

n 人の子どもがいて、それぞれ candies[i] 個のキャンディを持っている。あなたは
extraCandies
個の追加キャンディを持っており、ある子に全て与えた場合にその子が持つキャンディ数が配列内で最大値（複数可）に達するかどうかを判定して、長さ
n の boolean 配列を返す問題。

入力例:

- candies = [2,3,5,1,3], extraCandies = 3
- 出力: [true,true,true,false,true]

## 制約

- n == candies.length
- 2 <= n <= 100
- 1 <= candies[i] <= 100
- 1 <= extraCandies <= 50

これらはすべて小さい定数範囲なので、メモリ・時間ともに余裕がある。

## 考え方

1. 最初に全体の最大値 maxCandies を求める（配列を一度走査）。
2. 各子 i について、candies[i] + extraCandies >= maxCandies かを判定する。
   - もし真なら result[i] = true、そうでなければ false。
3. この方法で配列を2回走査する（最大値取得と各要素判定）か、1回で最大値取得後にもう1回判定する。いずれにせよ計算量は
   O(n)、追加メモリは結果配列分のみ O(n)。

性質のポイント（簡潔）:

- 線形時間・定数追加メモリ（出力除く）
- 比較のみ、整数オーバーフローの心配なし（制約が小さい）
- 並列化や順序の依存性はない

## キーワード

- 配列走査
- 最大値
- 比較
- O(n)

---

以下に Go 言語での具体的な実装を示します。ファイル名は指示どおり
main.go、テストは main_test.go としてください。

main.go

```go
package main

import "fmt"

// kidsWithCandies returns a boolean slice where result[i] is true if
// candies[i] + extraCandies >= max(candies).
func kidsWithCandies(candies []int, extraCandies int) []bool {
    if len(candies) == 0 {
        return []bool{}
    }

    max := candies[0]
    for _, v := range candies {
        if v > max {
            max = v
        }
    }

    res := make([]bool, len(candies))
    for i, v := range candies {
        res[i] = (v + extraCandies) >= max
    }
    return res
}

func main() {
    // 簡単なデモ
    candies := []int{2, 3, 5, 1, 3}
    extra := 3
    fmt.Println(kidsWithCandies(candies, extra)) // [true true true false true]
}
```

main_test.go

```go
package main

import (
    "reflect"
    "testing"
)

func TestKidsWithCandies(t *testing.T) {
    tests := []struct {
        candies      []int
        extraCandies int
        want         []bool
    }{
        {
            candies:      []int{2, 3, 5, 1, 3},
            extraCandies: 3,
            want:         []bool{true, true, true, false, true},
        },
        {
            candies:      []int{4, 2, 1, 1, 2},
            extraCandies: 1,
            want:         []bool{true, false, false, false, false},
        },
        {
            candies:      []int{12, 1, 12},
            extraCandies: 10,
            want:         []bool{true, false, true},
        },
        {
            candies:      []int{5, 5, 5},
            extraCandies: 0,
            want:         []bool{true, true, true},
        },
        {
            candies:      []int{1, 2},
            extraCandies: 50,
            want:         []bool{true, true},
        },
    }

    for _, tc := range tests {
        got := kidsWithCandies(tc.candies, tc.extraCandies)
        if !reflect.DeepEqual(got, tc.want) {
            t.Fatalf("candies=%v extra=%d => got=%v want=%v", tc.candies, tc.extraCandies, got, tc.want)
        }
    }
}
```
