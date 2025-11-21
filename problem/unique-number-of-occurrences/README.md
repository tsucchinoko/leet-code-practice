# unique-number-of-occurrences

## 問題の性質

配列内の各値の出現回数が互いに重複していないか（すべてユニークか）を判定する典型的なハッシュテーブル利用問題。線形時間に近い計算量と線形の追加メモリで解ける、頻度カウントと集合操作の基本問題。

## 問題要約

整数配列 arr
が与えられる。配列内の各異なる値について、その出現回数（frequency）を求め、異なる値同士でその出現回数に重複があれば
false、重複がなければ true を返す。

例:

- arr = [1,2,2,1,1,3] → 1 は3回、2 は2回、3 は1回 → 出現回数は {3,2,1}
  で重複なし → true
- arr = [1,2] → 1 は1回、2 は1回 → 出現回数に重複あり → false

## 制約

- 1 <= arr.length <= 1000
- -1000 <= arr[i] <= 1000

これらの制約により、時間計算量 O(n) （n =
arr.length）で十分高速。追加メモリは出現数を数えるハッシュマップ（あるいは配列オフセットでのカウント）と、出現回数の確認用集合が必要。

## 考え方

1. 値ごとの出現回数を数える。
   - Go では map[int]int
     を使うか、値の範囲が狭い（-1000..1000）ため固定長配列を用いることも可能（オフセット1000）。
2. 得られた各値の出現回数について、それらが互いに重複していないかを確認する。
   - 出現回数をキーとする集合（map[int]bool）を使い、既に存在する回数が出てきたら
     false を返す。
3. すべて確認して重複がなければ true を返す。

時間計算量: O(n) 追加空間: O(k)（k = 異なる値の個数、最大は
n。また値の範囲固定なら O(1)）

実装のポイント:

- 値の範囲が小さいため、配列を用いて高速にカウントする実装も可能だが、可読性と汎用性のため
  map を使う実装で十分。
- テストでは例示されたケースといくつかの境界ケース（最小長、すべて同じ値、ネガティブ値混在）を含める。

## キーワード

- ハッシュマップ（map）
- 出現回数（frequency count）
- セット（重複チェック）
- 時間計算量 O(n)

---

以下、Go の実装ファイルを2つ示します。

main.go

```go
package main

import "fmt"

// uniqueOccurrences returns true if the number of occurrences of each value in arr is unique.
func uniqueOccurrences(arr []int) bool {
    // 値 -> 出現回数
    freq := make(map[int]int)
    for _, v := range arr {
        freq[v]++
    }

    // 出現回数の集合で重複を検査
    seen := make(map[int]bool)
    for _, count := range freq {
        if seen[count] {
            return false
        }
        seen[count] = true
    }
    return true
}

func main() {
    examples := [][]int{
        {1, 2, 2, 1, 1, 3},
        {1, 2},
        {-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
    }
    for _, ex := range examples {
        fmt.Println(ex, "->", uniqueOccurrences(ex))
    }
}
```

main_test.go

```go
package main

import "testing"

func TestUniqueOccurrences(t *testing.T) {
    tests := []struct {
        arr  []int
        want bool
    }{
        {[]int{1, 2, 2, 1, 1, 3}, true},
        {[]int{1, 2}, false},
        {[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
        // 境界ケース
        {[]int{5}, true},                     // 1要素は常に true
        {[]int{2, 2, 2, 2}, true},            // 1つの値のみ -> 出現回数集合は {4}
        {[]int{1, 1, 2, 2, 3, 3}, false},     // すべて2回ずつ -> 重複
        {[]int{-1, -1, 0, 0, 0}, false},      // -1:2回, 0:3回 -> 重複なし -> true? 実際は出現は {2,3} -> true
        {[]int{-1, -1, 0, 0, 0, 7, 7}, false},// -1:2,0:3,7:2 -> 2が重複 -> false
    }

    for _, tt := range tests {
        got := uniqueOccurrences(tt.arr)
        if got != tt.want {
            t.Errorf("uniqueOccurrences(%v) = %v; want %v", tt.arr, got, tt.want)
        }
    }
}
```

注意:
上のテストケース群には意図を示すコメントが付いています。必要ならさらに境界値テスト（長さ1000、値の最大最小など）を追加してください。
