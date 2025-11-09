# reverse-vowels-of-a-string

## 問題の性質

文字列操作、双方向ポインタ（two
pointers）を使った簡単なアルゴリズム。インプレースで母音のみを入れ替える問題で、線形時間・定数追加空間で解ける典型的な「双方向スキャン」問題。

## 問題要約

与えられた文字列 s の中で、母音（a, e, i, o, u
と大文字小文字両方）だけを逆順に並べ替えた新しい文字列を返す。その他の文字の順序や位置は変えない。

例:

- 入力: "IceCreAm" → 出力: "AceCreIm"
- 入力: "leetcode" → 出力: "leotcede"

## 制約

- 1 <= s.length <= 3 * 10^5
- s は可視 ASCII 文字列（printable ASCII）
- 時間計算量は O(n) を目指す（n = s.length）
- 追加の空間は O(1)（文字配列に変換するための定数オーバーヘッドは許容）

## 考え方

1. 母音集合を定義する（小文字・大文字両方を含む、もしくは比較時に小文字化して判定）。
2. 文字列を可変な配列（runes または bytes）に変換する。ASCII 前提なら byte
   配列で十分だが、日本語環境に合わせても入力は printable ASCII のため byte
   を使える。
3. 左ポインタ i を先頭、右ポインタ j を末尾に置く。
4. i が指す文字が母音でない限り右へ進める。j が母音でない限り左へ進める。
5. 両方母音を指したら交換し、i を進め j を戻す。
6. i >= j になったら終了。
7. 交換後の配列を文字列に戻して返す。

時間計算量: O(n)、追加空間: O(n)（Go の文字列は不変なので内部で []byte
を作るが、実質追加は定数以外に依存しない実装も可能）。

注意点:

- ASCII 前提のためバイト単位での処理で問題ない。
- 大文字小文字の判定は母音集合に双方を登録するか、比較時に小文字に変換して行う。

## キーワード

- 双方向ポインタ (two pointers)
- 文字列操作
- 母音判定
- O(n) 時間, O(1) 追加空間（実装上は []byte 使用）

---

以下に Go 言語での具体的な実装を示します。main.go に解法、main_test.go
にテストコードを置いています。

```main.go
    package main

    import (
        "fmt"
    )

    func reverseVowels(s string) string {
        if len(s) <= 1 {
            return s
        }

        // 母音集合（ASCII のため byte 比較で OK）
        vowels := map[byte]bool{
            'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
            'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
        }

        b := []byte(s)
        i, j := 0, len(b)-1

        for i < j {
            // 左を母音に移動
            for i < j && !vowels[b[i]] {
                i++
            }
            // 右を母音に移動
            for i < j && !vowels[b[j]] {
                j--
            }
            if i < j {
                b[i], b[j] = b[j], b[i]
                i++
                j--
            }
        }

        return string(b)
    }

    func main() {
        // 簡単なデモ
        fmt.Println(reverseVowels("IceCreAm")) // AceCreIm
        fmt.Println(reverseVowels("leetcode")) // leotcede
    }
```

```main_test.go
package main

import "testing"

func TestReverseVowels(t *testing.T) {
    tests := []struct {
        in  string
        out string
    }{
        {"IceCreAm", "AceCreIm"},
        {"leetcode", "leotcede"},
        {"aA", "Aa"},
        {"hello", "holle"},
        {"", ""},
        {"bcd", "bcd"}, // 母音なし
        {"AEIOU", "UOIEA"},
        {"race car", "rece cara"}, // スペースを含む
        {"AaEeIiOoUu", "uUoOiIeEaA"},
    }

    for _, tt := range tests {
        got := reverseVowels(tt.in)
        if got != tt.out {
            t.Errorf("reverseVowels(%q) = %q; want %q", tt.in, got, tt.out)
        }
    }
}
```
