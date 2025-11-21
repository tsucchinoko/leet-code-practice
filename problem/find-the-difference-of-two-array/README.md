# find-the-difference-of-two-array

## 問題の性質

集合・差集合の計算、配列操作、重複の排除。両配列の要素をそれぞれ集合（ユニークな値）として扱い、片方に存在しない値を抽出するという単純な集合演算問題です。計算量は各配列の長さに対して線形に近いアルゴリズムで解けます。

## 問題要約

0-indexed の整数配列 nums1 と nums2 が与えられる。出力は長さ2のリスト answer
で、

- answer[0] は nums1 に含まれるが nums2 には含まれない異なる整数のリスト
- answer[1] は nums2 に含まれるが nums1 には含まれない異なる整数のリスト

リスト内の順序は任意。重複値は1回だけ含める。

例:

- nums1 = [1,2,3], nums2 = [2,4,6] -> [[1,3],[4,6]]
- nums1 = [1,2,3,3], nums2 = [1,1,2,2] -> [[3],[]]

## 制約

- 1 <= nums1.length, nums2.length <= 1000
- -1000 <= nums1[i], nums2[i] <= 1000

これよりメモリ・時間ともに O(n) 程度の解法が十分に効率的。

## 考え方

1. 各配列について重複を取り除く（集合化）。Go では map[int]struct{}
   を使ってユニーク化するのが簡単で高速。
2. nums1 のユニーク要素を順に見て、nums2 の集合に存在しなければ answer[0]
   に追加。
3. 同様に nums2 のユニーク要素を順に見て、nums1 の集合に存在しなければ answer[1]
   に追加。
4. 各要素の存在確認は map のキー存在チェックで O(1)（期待値）。したがって全体は
   O(n1 + n2) の時間で解ける。
5. 出力順序は問われないので、そのまま追加して良い。

注意点:

- 元配列内の重複は1回だけ扱う。
- 値の範囲が狭いため（-1000〜1000）、固定長配列を使うことも可能だが、可読性と一般性のため
  map を使用する方法が良い。

## キーワード

- 集合 (set)
- 差集合 (difference)
- 重複除去 (deduplication)
- ハッシュマップ (hash map)
- O(n)

---

以下、Go 言語での実装を示します。ファイルを2つ用意しています。

- main.go —
  実装と簡単な標準入出力用の例（関数をエクスポートしてテストから呼べるようにします）
- main_test.go — テストコード（複数のケースを含む）

main.go

```go
package main

import "fmt"

// FindDifference returns a slice of two slices:
// - result[0]: distinct integers in nums1 not present in nums2
// - result[1]: distinct integers in nums2 not present in nums1
func FindDifference(nums1 []int, nums2 []int) [][]int {
    set1 := make(map[int]struct{})
    set2 := make(map[int]struct{})

    for _, v := range nums1 {
        set1[v] = struct{}{}
    }
    for _, v := range nums2 {
        set2[v] = struct{}{}
    }

    res1 := make([]int, 0)
    for v := range set1 {
        if _, ok := set2[v]; !ok {
            res1 = append(res1, v)
        }
    }

    res2 := make([]int, 0)
    for v := range set2 {
        if _, ok := set1[v]; !ok {
            res2 = append(res2, v)
        }
    }

    return [][]int{res1, res2}
}

func main() {
    // 簡単な実行例
    a := []int{1, 2, 3}
    b := []int{2, 4, 6}
    fmt.Println(FindDifference(a, b)) // 例: [[1 3] [4 6]] （順序は異なる可能性あり）
}
```

main_test.go

```go
package main

import (
    "reflect"
    "sort"
    "testing"
)

// helper to compare two slices of slices irrespective of inner order and outer order
func equalResult(a, b [][]int) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if len(a[i]) != len(b[i]) {
            return false
        }
        sa := append([]int(nil), a[i]...)
        sb := append([]int(nil), b[i]...)
        sort.Ints(sa)
        sort.Ints(sb)
        if !reflect.DeepEqual(sa, sb) {
            return false
        }
    }
    return true
}

func TestFindDifference(t *testing.T) {
    tests := []struct {
        nums1 []int
        nums2 []int
        want  [][]int
    }{
        {[]int{1, 2, 3}, []int{2, 4, 6}, [][]int{{1, 3}, {4, 6}}},
        {[]int{1, 2, 3, 3}, []int{1, 1, 2, 2}, [][]int{{3}, {}}},
        {[]int{1}, []int{1}, [][]int{{}, {}}},
        {[]int{0, -1, -1, 2}, []int{-1, 3}, [][]int{{0, 2}, {3}}},
        {[]int{1000, -1000}, []int{500, -1000}, [][]int{{1000}, {500}}},
    }

    for _, tc := range tests {
        got := FindDifference(tc.nums1, tc.nums2)
        if !equalResult(got, tc.want) {
            t.Errorf("FindDifference(%v, %v) = %v; want %v", tc.nums1, tc.nums2, got, tc.want)
        }
    }
}
```

実行方法:

- go test をリポジトリルートで実行するとテストが通ることを確認できます。
