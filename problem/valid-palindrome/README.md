# valid-palindrome

## 問題の性質

文字列処理、双方向走査（two
pointers）、文字フィルタリング（英数字のみを抽出）、ケース正規化。線形時間で解ける決定問題（真偽値を返す）。

## 問題要約

与えられた文字列 s
について、英大文字を小文字に変換し、英数字以外の文字をすべて取り除いた結果が前から読んでも後ろから読んでも同じなら
true を、そうでなければ false を返す。

例:

- "A man, a plan, a canal: Panama" → "amanaplanacanalpanama" → 回文 → true
- "race a car" → "raceacar" → 回文でない → false
- " " → "" → 空文字は回文とみなす → true

## 制約

- 1 <= s.length <= 2 * 10^5
- s は表示可能な ASCII 文字のみ
- 時間計算量は O(n) を目指す（n = s.length）
- 追加メモリは可能なら O(1)（入力のインプレース走査）を目指す

## 考え方

1. 単純な方法:
   - 文字列を1回走査して英数字だけを抽出し、英字は小文字に正規化して新しい配列（または文字列）を作る。
   - 作った配列の前後を比較して回文判定する。
   - 時間 O(n)、追加メモリ O(n)。

2. 改良（推奨） — two pointers（左右からのポインタ）:
   - 左ポインタ i を先頭、右ポインタ j を末尾に置く。
   - i が指す文字が英数字でなければ i++、同様に j が指す文字が英数字でなければ
     j--。
   - どちらも英数字なら、英字は小文字にしてから比較。等しければ i++、j--
     を続ける。異なれば false。
   - i >= j になれば true。
   - このやり方は1回の走査で済み、追加メモリを定数に抑えられる（O(1)）。
   - ASCII のみという制約により、英字判定・小文字変換は簡単（'A'–'Z' を 'a'–'z'
     に変換など）。

3. 注意点:
   - 数字はそのまま比較可能。
   - 空白や句読点などは無視。
   - 大文字小文字は区別しない（小文字化）。
   - マルチバイト文字は出ない（表示可能 ASCII のみ）。

## キーワード

- 双方向走査 (two pointers)
- 文字フィルタリング
- 小文字正規化
- O(n) 時間、O(1) 追加空間

---

以下に Go の実装を示します。ファイル名は指定どおり `main.go`、テストは
`main_test.go` です。

```main.go
    package main

    import (
        "bufio"
        "fmt"
        "os"
    )

    func isAlphanumeric(b byte) bool {
        // ASCII の範囲で判定
        if b >= '0' && b <= '9' {
            return true
        }
        if b >= 'A' && b <= 'Z' {
            return true
        }
        if b >= 'a' && b <= 'z' {
            return true
        }
        return false
    }

    func toLower(b byte) byte {
        if b >= 'A' && b <= 'Z' {
            return b + ('a' - 'A')
        }
        return b
    }

    func isPalindrome(s string) bool {
        i, j := 0, len(s)-1
        for i < j {
            // move i to next alphanumeric
            for i < j && !isAlphanumeric(s[i]) {
                i++
            }
            // move j to prev alphanumeric
            for i < j && !isAlphanumeric(s[j]) {
                j--
            }
            if i >= j {
                break
            }
            if toLower(s[i]) != toLower(s[j]) {
                return false
            }
            i++
            j--
        }
        return true
    }

    func main() {
        in := bufio.NewReader(os.Stdin)
        var s string
        // Read whole line (including spaces)
        line, err := in.ReadString('\n')
        if err != nil && len(line) == 0 {
            // no input
            return
        }
        // Trim trailing newline characters
        // Keep other characters intact
        s = line
        // Remove trailing '\n' or '\r\n'
        if len(s) > 0 && s[len(s)-1] == '\n' {
            s = s[:len(s)-1]
            if len(s) > 0 && s[len(s)-1] == '\r' {
                s = s[:len(s)-1]
            }
        }
        fmt.Println(isPalindrome(s))
    }
```

```main_test.go
    package main

    import "testing"

    func TestIsPalindromeExamples(t *testing.T) {
        tests := []struct {
            input string
            want  bool
        }{
            {"A man, a plan, a canal: Panama", true},
            {"race a car", false},
            {" ", true},
            {"", true},
            {"0P", false},
            {"Able was I, ere I saw Elba", true},
            {"No 'x' in Nixon", true},
            {"Madam, I'm Adam", true},
            {"12321", true},
            {"1231", false},
        }

        for _, tc := range tests {
            got := isPalindrome(tc.input)
            if got != tc.want {
                t.Errorf("isPalindrome(%q) = %v; want %v", tc.input, got, tc.want)
            }
        }
    }
```
