# find-the-index-of-the-first-occurrence-in-a-string

## 問題の性質

文字列検索（部分文字列探索）。与えられた小さな文字列（needle）が大きな文字列（haystack）内に現れる最初の位置を求める典型的な問題。部分一致・完全一致の判定と位置の返却を行う。

## 問題要約

- 入力：文字列 haystack（大きい方）と needle（探す方）
- 出力：needle が haystack
  に最初に現れる開始インデックス（0-based）。存在しなければ -1。
- 例：
  - haystack = "sadbutsad", needle = "sad" → 出力 0（最初の出現位置）
  - haystack = "leetcode", needle = "leeto" → 出力 -1（存在しない）

## 制約

- 1 <= haystack.length, needle.length <= 10^4
- haystack と needle は小文字英字のみ
- 制約から考えて、最悪ケースでも O(n*m) の単純一致は 10^8
  程度の文字比較で収まる場合が多いが、最悪入力で時間がかかる可能性あり
- メモリは入力サイズに対して追加で大きな領域を用いない方針が望ましい

## 考え方

1. 単純なスライディングウィンドウ（naive）
   - haystack の各開始位置 i について、needle
     の各文字と比較して全一致するか調べる。
   - 実装は直感的で簡単。平均的な入力では高速だが、最悪ケース（例: haystack と
     needle が長いかつ多くの共有接頭辞を持つ場合）では O(n*m) の時間。
2. 改良アルゴリズム（必要なら）
   - KMP（Knuth–Morris–Pratt）：部分一致テーブル（LPS）を事前に作って不必要な比較をスキップ。時間計算量
     O(n + m)。
   - Rabin–Karp：ハッシュを使って高速に候補を絞る。ハッシュ衝突処理が必要。
3. 本問題の制約（最大
   10^4）を鑑みると、単純実装で十分でありコードの可読性が重要。ただし面接や高性能要求なら
   KMP を推奨。

簡潔な特徴：

- 可変長のパターン検索、最初の一致を返す
- 早期リターンが可能（最初に見つけたら即終了）
- アルゴリズム選択は入力の性質とパフォーマンス要件による

## キーワード

- 部分文字列探索
- スライディングウィンドウ
- Naive string search（単純一致）
- KMP（Knuth–Morris–Pratt）
- Rabin–Karp
- LPS（Longest Proper Prefix which is also Suffix）

---

以下に Go
言語での具体的な実装を示します。まずは分かりやすく安全な単純一致（スライディングウィンドウ）実装を提供します。必要なら
KMP 実装も追加できます。

ファイル構成：

- main.go — 実装
- main_test.go — テスト（標準ライブラリ testing を使用）

main.go

```go
package main

import (
    "fmt"
)

// strStr returns the index of the first occurrence of needle in haystack, or -1 if not present.
func strStr(haystack string, needle string) int {
    hn, nn := len(haystack), len(needle)
    if nn == 0 {
        return 0
    }
    if nn > hn {
        return -1
    }

    // Sliding window naive search
    for i := 0; i <= hn-nn; i++ {
        j := 0
        for j < nn && haystack[i+j] == needle[j] {
            j++
        }
        if j == nn {
            return i
        }
    }
    return -1
}

func main() {
    fmt.Println(strStr("sadbutsad", "sad"))   // 0
    fmt.Println(strStr("leetcode", "leeto"))  // -1
}
```

main_test.go

```go
package main

import "testing"

func TestStrStr(t *testing.T) {
    tests := []struct {
        haystack string
        needle   string
        want     int
    }{
        {"sadbutsad", "sad", 0},
        {"sadbutsad", "but", 3},
        {"leetcode", "leeto", -1},
        {"a", "a", 0},
        {"a", "b", -1},
        {"aaa", "aa", 0},        // first occurrence
        {"mississippi", "issip", 4},
        {"", "", 0},            // although constraints forbid empty, be robust
        {"abc", "", 0},
        {"", "a", -1},
    }

    for _, tt := range tests {
        got := strStr(tt.haystack, tt.needle)
        if got != tt.want {
            t.Errorf("strStr(%q, %q) = %d; want %d", tt.haystack, tt.needle, got, tt.want)
        }
    }
}
```

注意：

- テストには空文字列のケースを含めています（問題制約では長さ >=1
  だが、実装の堅牢性のため）。空の needle に対しては 0
  を返すのが一般的な定義です。
- 性能が問題となる入力が想定されるなら、KMP
  実装に置き換えてください。必要であれば KMP 実装も提供します。
