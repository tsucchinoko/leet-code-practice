# greatest-common-divisor-of-strings

## 問題の性質

文字列操作、数学的な最大公約数（GCD）的概念の応用。2つの文字列に共通する「周期（繰り返し単位）」のうち最長のものを求める問題。計算量は文字列長に依存するが、効率よく文字列比較と
GCD 計算を組み合わせて解ける。

## 問題要約

2つの文字列 str1, str2 が与えられる。文字列 t が s を割り切る（t divides
s）とは、s が t を1回以上繰り返して得られることを意味する。すなわち s = t + t +
... + t（t を繰り返し連結したもの）である。str1 と str2
の両方を割り切るような文字列 x
のうち、最長（かつ最大の）ものを返す。該当するものがなければ空文字列 "" を返す。

例：

- str1 = "ABCABC", str2 = "ABC" → "ABC"
- str1 = "ABABAB", str2 = "ABAB" → "AB"
- str1 = "LEET", str2 = "CODE" → ""

## 制約

- 1 <= str1.length, str2.length <= 1000
- str1 と str2 は英大文字のみ（A–Z）
- 出力は str1 と str2
  の両方を繰り返しで構成できる最長の共通基底文字列（存在しなければ空文字列）
- 時間計算量は O(n + m) 〜 O(n*m) 程度を目標（n,m はそれぞれの長さ）

## 考え方

1. 基本観察：
   - 任意の共通の「割る」文字列 x の長さは、str1.length と str2.length
     の公約数でなければならない。なぜなら str1, str2 の長さが x
     の長さで割り切れる必要があるから。
2. まずは単純チェック：
   - もし文字列を連結したときの等式 str1 + str2 == str2 + str1
     が成り立たなければ、共通の繰り返しパターンは存在しない（空文字列）。これは、両者が同じ基底文字列の繰り返しでできている場合にのみ成立する重要な条件。
3. 長さの最大公約数（GCD）を利用：
   - str1 と str2 の長さの GCD を g とする。すると、長さ g
     の接頭辞が答えになる可能性がある。なぜなら最長の共通基底の長さは必ず g
     の約数だが、上の等式が成り立てば長さ g の接頭辞が最長解になる。
4. 実装手順（効率的）：
   1. もし str1+str2 != str2+str1 なら "" を返す。
   2. g = gcd(len(str1), len(str2))
   3. 答えは str1[:g]（あるいは str2[:g]。等しいはず）
5. 補足（検証）：
   - str1[:g]
     が本当に両方を割り切るか気になる場合は、両方について繰り返して比較するチェックを入れても良い。ただし
     str1+str2 == str2+str1 を利用すれば不要。

問題の性質（簡潔）：

- 文字列の周期性と長さの GCD
  を利用する問題。文字列連結の可換性チェック（str1+str2 == str2+str1）と長さの
  GCD から直接解を得られるため、実装は単純で高速。

## キーワード

- 文字列の周期
- 最大公約数（GCD）
- 文字列連結チェック（str1+str2 == str2+str1）
- 接頭辞
- 文字列比較

---

以下に Go 言語での具体的な実装とテストコードを示します。

ファイル: main.go

```go
package main

import (
    "fmt"
    "io"
    "os"
)

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func gcdOfStrings(str1 string, str2 string) string {
    // もし連結順序が異なるときは共通の基底は存在しない
    if str1+str2 != str2+str1 {
        return ""
    }
    g := gcd(len(str1), len(str2))
    return str1[:g]
}

func main() {
    // 簡単な標準入力/出力のサンプル。実行時には引数を与えるか、stdin から読み取る。
    // ここでは固定サンプルを出力。
    fmt.Println(gcdOfStrings("ABCABC", "ABC"))   // "ABC"
    fmt.Println(gcdOfStrings("ABABAB", "ABAB"))  // "AB"
    fmt.Println(gcdOfStrings("LEET", "CODE"))    // ""
    // または stdin から読みたい場合は以下のようにする（必要に応じて有効化）
    // buf := make([]byte, 4096)
    // n, _ := os.Stdin.Read(buf)
    // input := string(buf[:n])
    // io.WriteString(os.Stdout, gcdOfStrings(...))
}
```

ファイル: main_test.go

```go
package main

import "testing"

func TestGcdOfStrings(t *testing.T) {
    tests := []struct {
        a, b string
        want string
    }{
        {"ABCABC", "ABC", "ABC"},
        {"ABABAB", "ABAB", "AB"},
        {"LEET", "CODE", ""},
        {"AAAA", "AA", "AA"},
        {"ABCABCABC", "ABCABC", "ABCABC"},
        {"XYZ", "XYZXYZ", "XYZ"},
        {"ABAB", "AB", "AB"},
        {"ABAB", "ABA", ""}, // 長さの観点で合わない例
        {"", "", ""},       // 仕様上長さ>=1だが堅牢性のためのテスト
    }

    for _, tt := range tests {
        got := gcdOfStrings(tt.a, tt.b)
        if got != tt.want {
            t.Fatalf("gcdOfStrings(%q, %q) = %q; want %q", tt.a, tt.b, got, tt.want)
        }
    }
}
```

注意：

- 問題制約では文字列長は少なくとも1だが、上のテストには空文字列を入れて堅牢性を確認している（実運用では省いて良い）。空文字列に対する動作はこの実装だと
  str1+str2 == str2+str1 が true になり、gcd(0, n) の扱いに依存するため注意。
- 上の実装は O(n + m) の文字列連結比較（実際は新しい文字列を作るコストがある）と
  GCD 計算 O(log(min(n,m)))
  に依存する。巨大入力でより最適化するなら連結比較を明示的ループで行う実装に変更可能。
