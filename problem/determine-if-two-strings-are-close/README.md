# determine-if-two-strings-are-close

## 問題の性質

文字列の変換可否を判定する問題。文字の位置入れ替え（任意の2文字のスワップ）と、ある既存文字を別の既存文字と全置換して入れ替える操作を何度でも使えるとき、2つの文字列が互いに到達可能（"close"）かを判定する決定問題。文字の出現集合と各文字の出現頻度の分布（マルチセット）に依存する性質を持つ。

## 問題要約

与えられた2つの小文字のみからなる文字列 word1, word2
が、以下の操作を任意回数適用して互いに変換可能かを判定する。

- Operation 1: 任意の2つの位置にある文字を入れ替える（任意の順列へ変更可能）。
- Operation 2: 文字 a と文字
  b（両方とも現在文字列に存在すること）を選び、すべての a を b に、すべての b を
  a に入れ替える（文字ラベルの交換）。

変換可能なら true、そうでなければ false を返す。

## 制約

- 1 <= word1.length, word2.length <= 10^5
- word1, word2 は小文字 'a'–'z' のみを含む
- 計算量・メモリは上記長さに対して現実的な範囲（例えば O(n)
  時間・O(1)（アルファベット定数）追加メモリが望ましい）

## 考え方

要点は Operation 1 と Operation 2 が許す変換の本質的な影響を理解すること。

- Operation
  1（任意スワップ）は文字列中の文字の順序を任意の順列にできるため、文字の「頻度分布（どの文字が何回現れるか）」以外の順序情報は自由に変えられる。したがって、最終的に重要なのは各文字の出現回数の集合（頻度のマルチセット）と、どの文字セットが使われているか、である。
- Operation
  2（ラベル交換）は、文字同士のラベルを入れ替える操作であり、文字ごとの頻度を別のラベルに再割り当てすることを可能にする。ただし、交換できるのは既に文字列に存在する文字同士のみ（存在しない文字を新規に導入はできない）。つまり、出現している文字の集合自体は変えられない（新しい文字を追加・削除は不可）
  — 正確には「どの文字が存在するか」の集合は変えられない。

これらから導かれる必要十分条件は次の2点：

1. word1 と word2 が使っている文字の集合が同じであること（例えば word1
   にしか現れない文字があれば、word2 側でその文字を作り出せない）。
2. 各文字の出現回数を集めた「頻度のマルチセット」が等しいこと。つまり、word1
   の全文字頻度の集合をソートしたものと、word2 のそれが一致すること。なぜなら
   Operation 2
   によって頻度同士を入れ替えられるので、どの文字がどの頻度を持つかではなく、頻度の分布自体が一致していればラベル交換で対応付け可能だからである。

この 2 条件が満たされれば、Operation 1
による並べ替えを併用して実際の文字列を一致させられるため close である。

計算手順（効率良く）：

- 長さが異なる場合は即 false（長さ保存）。
- 26文字固定なので配列 size=26 の整数カウントをそれぞれ作成して頻度を得る。
- 両方のカウント配列で値が 0 と 0
  以外の位置を比較し、存在集合が一致しているか確認する（任意の文字について片方は
  0 でもう片方は非0 なら false）。
- 非ゼロの頻度だけを取り出して昇順ソートし、両方が一致するかを確認する（計算量は
  O(26 log 26) ≒ 定数）。
- 上記を満たせば true、そうでなければ false。

時間計算量: O(n)（文字列長のスキャン）＋定数ソート。 空間計算量:
O(1)（26要素配列など、文字数に依存する定数）。

## キーワード

- 頻度カウント
- マルチセット（頻度の分布）
- ラベル交換（文字の全置換）
- 定数アルファベット（26）
- O(n) 時間

---

以下に Go 言語での実装を、メイン実行ファイル main.go とテストファイル
main_test.go として示します。

main.go

```go
package main

import (
    "fmt"
    "os"
    "sort"
)

func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }

    var cnt1 [26]int
    var cnt2 [26]int
    for i := 0; i < len(word1); i++ {
        cnt1[word1[i]-'a']++
    }
    for i := 0; i < len(word2); i++ {
        cnt2[word2[i]-'a']++
    }

    // 出現文字集合が同じか確認
    for i := 0; i < 26; i++ {
        if (cnt1[i] == 0) != (cnt2[i] == 0) {
            return false
        }
    }

    // 出現頻度のマルチセットを比較（非ゼロのみ）
    freqs1 := make([]int, 0, 26)
    freqs2 := make([]int, 0, 26)
    for i := 0; i < 26; i++ {
        if cnt1[i] > 0 {
            freqs1 = append(freqs1, cnt1[i])
        }
        if cnt2[i] > 0 {
            freqs2 = append(freqs2, cnt2[i])
        }
    }
    sort.Ints(freqs1)
    sort.Ints(freqs2)
    if len(freqs1) != len(freqs2) {
        return false
    }
    for i := range freqs1 {
        if freqs1[i] != freqs2[i] {
            return false
        }
    }
    return true
}

func main() {
    // 簡単なコマンドライン実行用の例
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run main.go <word1> <word2>")
        return
    }
    w1 := os.Args[1]
    w2 := os.Args[2]
    fmt.Println(closeStrings(w1, w2))
}
```

main_test.go

```go
package main

import "testing"

func TestCloseStrings(t *testing.T) {
    tests := []struct {
        w1   string
        w2   string
        want bool
    }{
        {"abc", "bca", true},
        {"a", "aa", false},
        {"cabbba", "abbccc", true},
        {"abbzzca", "babzzcz", false}, // 異なる出現集合
        {"aaabbb", "ababab", true},   // 同じ頻度分布・集合
        {"abc", "def", false},        // 使われる文字集合が違う
        {"", "", true},               // 空文字（制約では長さ>=1だが保険）
    }

    for _, tt := range tests {
        got := closeStrings(tt.w1, tt.w2)
        if got != tt.want {
            t.Errorf("closeStrings(%q, %q) = %v; want %v", tt.w1, tt.w2, got, tt.want)
        }
    }
}
```
