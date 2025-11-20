# maximum-number-of-vowels-in-a-substring-of-given-length

## 問題の性質

スライディングウィンドウ（連続部分列）を使った配列・文字列の区間最適化問題。部分列の長さが固定されたときに、区間内の特定の要素（ここでは母音）の個数を最大化する問題で、線形時間で解ける。

## 問題要約

文字列 s と整数 k が与えられる。長さ k
の連続する部分文字列（サブストリング）の中で、母音（a, e, i, o,
u）の数が最大となる値を返す。

例:

- s = "abciiidef", k = 3 → 出力: 3（"iii"）
- s = "aeiou", k = 2 → 出力: 2
- s = "leetcode", k = 3 → 出力: 2

## 制約

- 1 <= s.length <= 10^5
- s は小文字英字のみ
- 1 <= k <= s.length

これらの制約のもとで、計算量は O(n)（n = s.length）を目指す。追加の空間は定数
O(1) にできる。

## 考え方

1. 母音判定を高速に行うため、a, e, i, o, u
   をハッシュセットまたはビットマスクで保持する。
2. 長さ k のスライディングウィンドウを文字列上で左から右へ1文字ずつ移動させる。
   1. 最初のウィンドウ（先頭 k 文字）の母音数を数える（初期カウント）。
   2. ウィンドウを1文字右に移動するときは、左端の文字の影響（もし母音ならカウントを減らす）と、新しく右端に入る文字の影響（もし母音ならカウントを増やす）を反映させる。各ステップは
      O(1)。
3. 各ウィンドウでの母音数の最大値を追跡し、最後に返す。
4. 特別ケース: k == 0（ここでは k>=1 のため不要）や s.length == k
   の場合は初期カウントをそのまま返す。

性質の簡潔な説明:

- 連続部分列の固定長で局所的な寄与（加算・減算）だけを更新するため、線形時間で解ける典型的なスライディングウィンドウ問題。

## キーワード

- スライディングウィンドウ
- 母音判定
- O(n) 時間、O(1) 追加空間
- 逐次更新（差分更新）

---

以下に Go 言語での実装（main.go）とテストコード（main_test.go）を示します。

main.go

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func maxVowels(s string, k int) int {
    isVowel := func(c byte) bool {
        switch c {
        case 'a', 'e', 'i', 'o', 'u':
            return true
        }
        return false
    }

    n := len(s)
    if k > n {
        k = n
    }

    // 初期ウィンドウ
    count := 0
    for i := 0; i < k; i++ {
        if isVowel(s[i]) {
            count++
        }
    }
    maxCount := count

    // スライド
    for i := k; i < n; i++ {
        if isVowel(s[i-k]) {
            count--
        }
        if isVowel(s[i]) {
            count++
        }
        if count > maxCount {
            maxCount = count
        }
    }
    return maxCount
}

func main() {
    in := bufio.NewReader(os.Stdin)
    var s string
    var k int
    // 入力例（簡易パース）: 1行目に文字列、2行目に整数 k を想定
    if _, err := fmt.Fscanln(in, &s); err != nil {
        // handle if input contains spaces by reading whole line
        line, _ := in.ReadString('\n')
        s = strings.TrimSpace(line)
    }
    if _, err := fmt.Fscanln(in, &k); err != nil {
        // try read next line
        fmt.Fscan(in, &k)
    }
    fmt.Println(maxVowels(s, k))
}
```

main_test.go

```go
package main

import "testing"

func TestMaxVowels(t *testing.T) {
    tests := []struct {
        s    string
        k    int
        want int
    }{
        {"abciiidef", 3, 3},
        {"aeiou", 2, 2},
        {"leetcode", 3, 2},
        {"rhythms", 4, 0},    // 母音なし
        {"a", 1, 1},          // 最小ケース
        {"abcde", 5, 2},      // 全長ウィンドウ
        {"abecidofu", 3, 3},  // 複数母音の分布
    }

    for _, tt := range tests {
        got := maxVowels(tt.s, tt.k)
        if got != tt.want {
            t.Fatalf("maxVowels(%q, %d) = %d; want %d", tt.s, tt.k, got, tt.want)
        }
    }
}
```

注意:

- main.go
  の入力パースは簡易的で、1行目に文字列、2行目に整数を想定しています。用途に応じて入力形式を調整してください。
