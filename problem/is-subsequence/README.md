# is-subsequence

## 問題の性質

文字列操作、探索、貪欲法。基本問題は単一判定（s が t
の部分列かどうか）で、入力サイズは小さいがフォローアップでは多数の s
を効率的に判定するため前処理（インデックス出現テーブル）を使う典型問題になる。計算量とメモリのトレードオフが重要。

## 問題要約

2つの文字列 s と t が与えられる。s の文字を順序を保ったまま t
から削除操作で取り出せるなら true を返し、そうでなければ false
を返す。「部分列(subsequence)」は連続である必要はないが順序は保たれる。

例:

- s = "abc", t = "ahbgdc" → true
- s = "axc", t = "ahbgdc" → false

フォローアップ: 判定したい s が非常に多数（k ≥ 10^9）ある場合、どう変更するか。

## 制約

- 0 <= s.length <= 100
- 0 <= t.length <= 10^4
- s, t は小文字英字のみ
- 標準ケース: 単一判定（1 回の s）
- フォローアップ: 多数の異なる s に対する多数回の判定（高速化が必要）

## 考え方

1. 基本（単一判定） — 貪欲2ポインタ:
   - i を s の現在位置、j を t の現在位置として両方 0 から開始。
   - j を先に進めつつ t[j] が s[i] と一致したら i++ して次の文字を探す。
   - 最終的に i が s.length に達していれば true、そうでなければ false。
   - 時間計算量 O(|t| + |s|)（実装上は O(|t|)）、追加空間 O(1)。

2. フォローアップ（多数の s を高速判定） — 事前インデックス（next ポインタ /
   遷移テーブル）:
   - t を前処理して、各位置 pos と各文字 c に対して、pos
     から右へ進んだとき最初に現れる c の位置 next[pos][c] を作る。
   - 実装方法:
     - next を (|t|+1) x 26 の配列で持つ。next[pos][c] は pos から c
       が現れる最小インデックス（存在しなければ -1 または |t| を示す）。
     - 構築は逆走査で O(|t| * 26)。
   - 各 s の判定は pos=0 から s の文字を順に next[pos][c]
     を見て更新していき、存在しない（-1）なら false、最後まで見つかれば true。
   - 各判定は O(|s|)（|s| ≤ 100）。前処理は O(|t| * 26) とメモリ
     O((|t|+1)*26)（|t| ≤ 10^4 なので現実的）。
   - これで多数の s に対して高速に応答可能（k
     が非常に大きくても各判定は短時間）。

3. 他の手法:
   - 文字ごとの出現インデックスリスト（map<char,
     []indices>）を作り、各文字で二分探索して次の位置を探す方法。前処理 O(|t|)
     と空間 O(|t|)。各 s 判定は O(|s| log n)（n:平均出現数）。
   - next テーブルは定数時間遷移できる点でより高速（ただしメモリはやや大きい）。

ポイント:

- s が短く t が長い場合、単純2ポインタが最もシンプルで十分。
- 多数判定（フォローアップ）は前処理 + O(|s|) 判定の手法が有利。
- 制約内ではいずれの方法も安全に動作する。

## キーワード

- 部分列 (subsequence)
- 2ポインタ (two pointers)
- 貪欲法 (greedy)
- 前処理 (preprocessing)
- next 配列 / 遷移テーブル
- 二分探索による位置検索

---

以下に Go の実装を示す。ファイル構成:

- main.go — 単純な部分列判定（2ポインタ方式）と、フォローアップ用の事前処理型
  SubsequenceChecker を含む。
- main_test.go —
  単体テスト。基本ケースとフォローアップの複数問い合わせケースを確認する。

両ファイルとも go test で実行できるようにしてあります。

```main.go
package main

import (
    "fmt"
)

// IsSubsequenceSimple returns true if s is a subsequence of t using two-pointer greedy.
func IsSubsequenceSimple(s, t string) bool {
    si, ti := 0, 0
    sn, tn := len(s), len(t)
    for si < sn && ti < tn {
        if s[si] == t[ti] {
            si++
        }
        ti++
    }
    return si == sn
}

// SubsequenceChecker precomputes next positions for t to answer many queries efficiently.
type SubsequenceChecker struct {
    // next[pos][c] = next index >= pos where character c appears in t, or -1 if none
    next [][]int
    n    int
}

// NewSubsequenceChecker builds the next table for string t.
func NewSubsequenceChecker(t string) *SubsequenceChecker {
    n := len(t)
    // next size: (n+1) x 26
    next := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        next[i] = make([]int, 26)
        for j := 0; j < 26; j++ {
            next[i][j] = -1
        }
    }
    // fill from back: at position n (beyond end) all are -1
    for i := n - 1; i >= 0; i-- {
        // copy row i+1 to i
        for c := 0; c < 26; c++ {
            next[i][c] = next[i+1][c]
        }
        // set current character
        ch := t[i] - 'a'
        next[i][ch] = i
    }
    return &SubsequenceChecker{next: next, n: n}
}

// IsSubsequence checks whether s is subsequence of original t using precomputed table.
func (sc *SubsequenceChecker) IsSubsequence(s string) bool {
    pos := 0
    for i := 0; i < len(s); i++ {
        c := s[i] - 'a'
        if pos > sc.n {
            return false
        }
        nxt := sc.next[pos][c]
        if nxt == -1 {
            return false
        }
        pos = nxt + 1
    }
    return true
}

func main() {
    // Example usage
    fmt.Println(IsSubsequenceSimple("abc", "ahbgdc")) // true
    fmt.Println(IsSubsequenceSimple("axc", "ahbgdc")) // false

    sc := NewSubsequenceChecker("ahbgdc")
    fmt.Println(sc.IsSubsequence("abc")) // true
    fmt.Println(sc.IsSubsequence("axc")) // false
}
```

```main_test.go
package main

import "testing"

func TestIsSubsequenceSimple(t *testing.T) {
    tests := []struct {
        s, t string
        want bool
    }{
        {"abc", "ahbgdc", true},
        {"axc", "ahbgdc", false},
        {"", "ahbgdc", true},
        {"a", "", false},
        {"", "", true},
        {"ace", "abcde", true},
        {"aec", "abcde", false},
    }
    for _, tc := range tests {
        got := IsSubsequenceSimple(tc.s, tc.t)
        if got != tc.want {
            t.Fatalf("IsSubsequenceSimple(%q, %q) = %v; want %v", tc.s, tc.t, got, tc.want)
        }
    }
}

func TestSubsequenceChecker(t *testing.T) {
    tstr := "ahbgdc"
    sc := NewSubsequenceChecker(tstr)

    tests := []struct {
        s   string
        want bool
    }{
        {"abc", true},
        {"axc", false},
        {"", true},
        {"ahbgdc", true},
        {"ahbgdcz", false},
        {"hb", true},
    }

    for _, tc := range tests {
        got := sc.IsSubsequence(tc.s)
        if got != tc.want {
            t.Fatalf("NewSubsequenceChecker(%q).IsSubsequence(%q) = %v; want %v", tstr, tc.s, got, tc.want)
        }
    }
}

// Benchmark-style test for many queries simulated (not using testing.B).
func TestManyQueries(t *testing.T) {
    tstr := "abacbabcdelmnopqrsabctuvwxyzab" // sample t
    sc := NewSubsequenceChecker(tstr)

    // simulate many small s queries
    queries := []string{"abc", "abct", "mnop", "zz", "ab", "qrst", "abctu"}
    for _, q := range queries {
        _ = sc.IsSubsequence(q) // ensure no panic and runs fast
    }
}
```
