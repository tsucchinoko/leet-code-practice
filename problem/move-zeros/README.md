# move-zeros

## 問題の性質

配列操作（インプレース）。安定性（非ゼロ要素の相対順序を維持）を保ちながら、配列内の0を末尾に移動する。追加配列を使わず、定数追加メモリで行う。

## 問題要約

整数配列 nums が与えられる。配列内のすべての 0
を配列の末尾に移動し、非ゼロ要素の相対順序は維持すること。配列はその場（in-place）で変更し、追加配列のコピーは作らないこと。

例:

- 入力: nums = [0,1,0,3,12] 出力: [1,3,12,0,0]
- 入力: nums = [0] 出力: [0]

## 制約

- 1 <= nums.length <= 10,000
- -2^31 <= nums[i] <= 2^31 - 1
- 追記: メモリは O(1)（定数追加メモリ）を目指す。配列は in-place で変更する。

## 考え方

主な方針は2ポインタ（またはスロー＆ファスト）テクニックを使うこと。

基本アルゴリズム（2ポインタ、安定）:

1. write インデックスを 0 に初期化する（非ゼロ要素を書き込む位置）。
2. nums を順に走査する読み取りインデックス i を 0 から n-1 まで進める。
3. nums[i] が非ゼロなら、nums[write] = nums[i] として書き、write
   をインクリメントする。
   - このとき i と write
     が異なる場合にのみ代入することで不要な書き換えを減らせる（write==i
     のときはそのまま）。
4. 走査が終わったら、write 以降の位置をすべて 0 にする。

操作回数最小化の観点:

- 非ゼロ要素の総数を k とすると、非ゼロ要素は最大 k 回の書き込み、残りの n-k
  個の 0 を書き込むため n-k 回の代入が必要。不要な自己代入を避ければ十分効率的。
- 時間計算量は O(n)、追加メモリは O(1)。

別解候補:

- 2ポインタで非ゼロとゼロを見つけてスワップする方法もある（swap
  を使うと非ゼロの順序は維持される）。ただしスワップは書き込みが2回になることがあるため、不要な書き込み回数が増える点に注意。

問題の性質（簡潔に）:

- 安定な配列変換
- インプレース、定数追加メモリ
- 走査一回で解ける（O(n) 時間）

## キーワード

- Two pointers（2ポインタ）
- In-place
- Stable (相対順序維持)
- O(n) time, O(1) space
- Write/Read pointer

---

以下、Go言語での具体的実装とテストコードを示します。

main.go（実装）:

```go
package main

import "fmt"

// moveZeroes moves all zeros in nums to the end while maintaining the relative order
// of the non-zero elements. It modifies nums in-place.
func moveZeroes(nums []int) {
    write := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            if i != write {
                nums[write] = nums[i]
            }
            write++
        }
    }
    for write < len(nums) {
        nums[write] = 0
        write++
    }
}

func main() {
    examples := [][]int{
        {0, 1, 0, 3, 12},
        {0},
        {1, 0, 2, 0, 0, 3},
        {0, 0, 0},
        {1, 2, 3},
    }

    for _, ex := range examples {
        nums := make([]int, len(ex))
        copy(nums, ex)
        moveZeroes(nums)
        fmt.Println(nums)
    }
}
```

main_test.go（テスト）:

```go
package main

import (
    "reflect"
    "testing"
)

func TestMoveZeroes(t *testing.T) {
    tests := []struct {
        input  []int
        expect []int
    }{
        {[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
        {[]int{0}, []int{0}},
        {[]int{1, 0, 2, 0, 0, 3}, []int{1, 2, 3, 0, 0, 0}},
        {[]int{0, 0, 0}, []int{0, 0, 0}},
        {[]int{1, 2, 3}, []int{1, 2, 3}},
        {[]int{0, 1, 2, 0, 3, 0, 4}, []int{1, 2, 3, 4, 0, 0, 0}},
    }

    for _, tt := range tests {
        input := make([]int, len(tt.input))
        copy(input, tt.input)
        moveZeroes(input)
        if !reflect.DeepEqual(input, tt.expect) {
            t.Errorf("moveZeroes(%v) = %v; want %v", tt.input, input, tt.expect)
        }
    }
}
```

補足:

- テストは標準的な table-driven
  テスト形式で、入力配列は意図せず変更されないようにコピーしてから関数を呼んでいます（関数自体は
  in-place 変更します）。
- moveZeroes の実装は不要な自己代入を避け、書き込み回数を最小限にしています。
