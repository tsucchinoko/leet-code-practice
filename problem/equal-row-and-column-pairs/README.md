# equal-row-and-column-pairs

## 問題の性質

- 行（row）と列（column）の「配列」が完全に一致するかを数える組合せ問題。
- 比較は要素の順序も含めた完全一致（等しい配列）を要求する。
- 正方行列（n x n）。

## 問題要約

0-indexed の n x n 整数行列 grid が与えられる。 行 ri と列 cj の組 (ri, cj)
のうち、行 ri と列 cj
の要素列が同じ（同じ値を同じ順序で持つ）となるペアの総数を返す。

例:

- grid = [[3,2,1],[1,7,6],[2,7,7]] → 出力 1（Row 2 = [2,7,7] と Column 1 =
  [2,7,7] の1組）
- grid = [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]] → 出力 3

## 制約

- n == grid.length == grid[i].length
- 1 <= n <= 200
- 1 <= grid[i][j] <= 10^5

制約より、O(n^3)（各行と各列を要素比較で逐一比較する手法）は n=200 で最悪 200^3
= 8,000,000 要素比較となり現実的に許容範囲だが、より効率的に O(n^2)〜O(n^2 log
n) にできる。

## 考え方

目的は「各行ベクトル」と「各列ベクトル」の一致ペアを数えること。一般的なアプローチは次の通り。

1. 各行を一意に表現（キー化）する:
   - 行を文字列化（例: 要素を区切り文字で連結）するか、固定長の配列をそのまま
     map のキーに使える型へ変換する（Go ではスライスは map
     のキーに直接使えないため文字列化や配列化が必要）。
   - 文字列化する際は区切り文字に注意（値の範囲が大きい場合でも衝突しないように安全な区切りを用いる）。
2. 各行キーの出現回数を map で数える（行ごとにカウント）。
3. 各列も同様にキー化して、そのキーが行の map
   にあればそのカウント分だけ結果に足す。

計算量:

- 行のキー作成: O(n^2)
- 列のキー作成と照合: O(n^2)
- 全体: O(n^2) 時間、O(n^2) 空間（キー化用）

実装上の注意:

- Go ではスライスを map のキーにできないため、行/列を string
  にシリアライズする（例えばバイナリ的にエンコード、または区切り付き文字列）。要素の最大値が
  10^5 なので区切りに ',' や '#' を使えば安全。
- もう一つの手法は map[string]int の代わりに map[[maxN]int]int
  のような固定長配列をキーにすることだが n は動的なので文字列化が汎用的。

問題の性質（要約）:

- 比較操作のコストが主要因で、キー化して頻度を数えることで二重ループ＋要素比較を避けられる。
- ハッシュ（map）を使って一致判定を高速化する典型的なパターン。

## キーワード

- ハッシュ（map）
- キー化 / シリアライズ
- 行列の列取り出し
- 頻度カウント
- 時間計算量 O(n^2)

---

以下に Go 言語での具体的実装を示します。ファイルは main.go（実装）と
main_test.go（テスト）として出力します。

main.go

```go
package main

import (
    "fmt"
    "strings"
)

func equalPairs(grid [][]int) int {
    n := len(grid)
    rowCount := make(map[string]int, n)

    // 行をキー化してカウント
    for i := 0; i < n; i++ {
        // セパレータにカンマを使う（値は正の整数なので安全）
        parts := make([]string, n)
        for j := 0; j < n; j++ {
            parts[j] = fmt.Sprintf("%d", grid[i][j])
        }
        key := strings.Join(parts, ",")
        rowCount[key]++
    }

    // 各列をキー化して rowCount に存在すれば加算
    res := 0
    for j := 0; j < n; j++ {
        parts := make([]string, n)
        for i := 0; i < n; i++ {
            parts[i] = fmt.Sprintf("%d", grid[i][j])
        }
        key := strings.Join(parts, ",")
        if c, ok := rowCount[key]; ok {
            res += c
        }
    }

    return res
}

func main() {
    // 簡単なデモ
    g1 := [][]int{{3,2,1},{1,7,6},{2,7,7}}
    fmt.Println(equalPairs(g1)) // 1

    g2 := [][]int{{3,1,2,2},{1,4,4,5},{2,4,2,2},{2,4,2,2}}
    fmt.Println(equalPairs(g2)) // 3
}
```

main_test.go

```go
package main

import "testing"

func TestEqualPairs(t *testing.T) {
    tests := []struct {
        grid [][]int
        want int
    }{
        {
            grid: [][]int{{3,2,1},{1,7,6},{2,7,7}},
            want: 1,
        },
        {
            grid: [][]int{{3,1,2,2},{1,4,4,5},{2,4,2,2},{2,4,2,2}},
            want: 3,
        },
        {
            grid: [][]int{{1}},
            want: 1,
        },
        {
            grid: [][]int{{1,2},{2,1}},
            want: 0,
        },
        {
            grid: [][]int{{1,1,1},{1,1,1},{1,1,1}},
            want: 9, // 全ての行と列が [1,1,1] なので 3*3 = 9
        },
    }

    for _, tt := range tests {
        if got := equalPairs(tt.grid); got != tt.want {
            t.Fatalf("equalPairs(%v) = %d; want %d", tt.grid, got, tt.want)
        }
    }
}
```
