# decode-string

## 問題の性質

スタックまたは再帰を用いる典型的な文字列処理問題。入れ子になった繰り返し表現
k[... ]
を展開する。入力は常に正しい形式で、数字は繰り返し回数のみを表す（文字列本体に数字は含まれない）。

## 問題要約

エンコードされた文字列 s が与えられる。エンコード規則は k[encoded_string]
で、角括弧内の encoded_string をちょうど k 回繰り返す。入れ子（例えば
3[a2[c]]）も可能。展開した結果の文字列を返す。

例:

- s = "3[a]2[bc]" -> "aaabcbc"
- s = "3[a2[c]]" -> "accaccacc"
- s = "2[abc]3[cd]ef" -> "abcabccdcdcdef"

## 制約

- 1 <= s.length <= 30
- s は小文字アルファベット、数字、角括弧 '[]' からなる
- s は常に正しい形式（空白や未整合な括弧はない）
- 文字列本体に数字は含まれない（数字は繰り返し回数のみ）
- すべての整数は 1 以上 300 以下
- 出力長は最大 10^5 を超えない（テストケース生成

の保証）

## 考え方

主なアイデアはスタックまたは再帰で括弧構造を扱うこと。

方法1（スタック）:

- 文字を左からスキャンする。
- 数字を見つけたらその数値（k）を構築して数値スタックに入れる。
- '['
  のときは現在構築中の文字列（部分文字列）を文字列スタックに入れて、新しい部分文字列を開始する。
- 文字が来たら現在の部分文字列に追加する。
- ']' のときは、数値スタックから k を取り出し、文字列スタックから直前の文字列
  prev を取り出す。現在の部分文字列 cur を k 回繰り返し prev
  に連結してこれを新しい cur とする。
- 最終的に cur が答えとなる。

方法2（再帰）:

- ポインタを持って再帰的に parse を行い、'[' を見たら内部を再帰で処理して ']'
  までの部分文字列を得る。数字と組み合わせて

繰り返す。

計算量:

- 各文字は定数回処理されるため O(n + output_length)。出力サイズを M とすると時間
  O(M)（展開に比例）。
- 追加の空間はスタックや結果保持で O(M)。

問題の性質としては「文字列の構文解析」と「スタック／再帰の基本的利用」に分類される。実装は比較的単純でバグの出やすい箇所は数字の複数桁処理、入れ子と現在の部分文字列の管理。

## キーワード

- スタック
- 再帰
- 文字列展開
- 逐次パース
- 入れ子構造

---

以下に Go 言語での実装例を示す。ファイル構成:

- main.go — 実装
- main_test.go — テスト（標準の go test で実行）

main.go

```go
package main

import (
    "fmt"
    "unicode"
)

func decodeString(s string) string {
    // スタックを使った実装
    var strStack []string
    var countStack []int
    cur := ""
    num := 0

    for _, ch := range s {
        if unicode.IsDigit(ch) {
            // 数字は複数桁対応
            num = num*10 + int(ch-'0')
        } else if ch == '[' {
            // 現在の数と文字列をスタックに保存して新しい cur を開始
            countStack = append(countStack, num)
            strStack = append(strStack, cur)
            num = 0
            cur = ""
        } else if ch == ']' {
            // スタックから取り出して展開
            if len(countStack) == 0 || len(strStack) == 0 {
                return "" // 入力保証があるため通常は到達しない
            }
            k := countStack[len(countStack)-1]
            countStack = countStack[:len(countStack)-1]
            prev := strStack[len(strStack)-1]
            strStack = strStack[:len(strStack)-1]

            // cur を k 回繰り返す
            repeated := ""
            // 効率を多少改善するためにビルドアップ
            for i := 0; i < k; i++ {
                repeated += cur
            }
            cur = prev + repeated
        } else {
            // 英小文字
            cur += string(ch)
        }
    }
    return cur
}

func main() {
    // 簡単なデモ
    examples := []string{"3[a]2[bc]", "3[a2[c]]", "2[abc]3[cd]ef"}
    for _, e := range examples {
        fmt.Println(e, "->", decodeString(e))
    }
}
```

main_test.go

```go
package main

import (
    "testing"
)

func TestDecodeString(t *testing.T) {
    cases := []struct {
        in  string
        out string
    }{
        {"3[a]2[bc]", "aaabcbc"},
        {"3[a2[c]]", "accaccacc"},
        {"2[abc]3[cd]ef", "abcabccdcdcdef"},
        {"10[a]", "aaaaaaaaaa"},
        {"1[abc]", "abc"},
        {"2[3[a]b]", "aaabaaab"},
        {"3[z2[y]]", "zyyzyyzyy"},
    }

    for _, c := range cases {
        got := decodeString(c.in)
        if got != c.out {
            t.Fatalf("decodeString(%q) == %q, want %q", c.in, got, c.out)
        }
    }
}
```

補足:

- 本実装は単純に文字列連結（repeated +=
  cur）を使っているため、大きな出力に対しては効率面で改善の余地がある（strings.Builder
  を使う、あるいはバイトスライスでビルドする等）。
- しかしこの問題の制約（最終出力長 <= 10^5）ではこの実装で十分実用的。
