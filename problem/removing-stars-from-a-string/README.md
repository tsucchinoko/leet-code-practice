# removing-stars-from-a-string

## 問題の性質

スタック操作に相当する文字列処理問題。
左側にある直近の非'*'文字を削除する操作が繰り返されるため、逐次処理（スキャン）で解ける。結果は一意。

## 問題要約

文字列 s が与えられる。s には小文字英字と '*' が含まれる。操作は以下：

- 任意の '_' を選び、その左側にある最も近い非 '_'
  文字（存在することは保証される）とその '*' 自身を削除する。

すべての '*'
が取り除かれたあとの文字列を返す。操作順序は任意だが、最終結果は一意である。

例:

- s = "leet**cod*e" → "lecoe"
- s = "erase*****" → ""

## 制約

- 1 <= s.length <= 10^5
- s は小文字英字と '*' のみから成る
- 操作は常に可能（各 '*' に対して左側に削除対象が存在するように生成される）
- 結果は一意

計算量目安:

- 時間: O(n

)

- 追加空間: O(n)（出力やスタックに相当するメモリ）

## 考え方

この問題は「スタック」や「結果を構築する配列」を使って線形スキャンで解ける。

アルゴリズム:

- 結果を格納する可変配列（あるいはスライス）res を用意する。
- 文字列 s を左から右へ一文字ずつ見る：
  - 現在の文字が '*' なら、res の最後の要素を削除（pop）する。
  - そうでなければ、res にその文字を追加（push）する。
- スキャン終了後、res を連結して返す。

理由:

- '_' が現れた時に直近の左側の非 '_'
  を削除すればよく、スタックの末尾が常に「直近の未削除の文字」を表すため正しい。
- 各文字は最大1回 push / 1回 pop されるだけなので O(n)。

特記事項:

- Go の場合、バイトスライスで操作すると効率的（文字は ASCII の小文字 と '*'
  のみ）。
- 文字列の再作成を避けるため、res を []byte として使い、最後に string(res)
  とする。

## キーワード

-

スタック

- 文字列スキャン
- 線形時間 O(n)
- in-place 相当の構築（追加空間は出力分）

---

以下に Go 言語での具体的な実装を示します。ファイルは main.go と main_test.go
として構成しています。

main.go

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func removeStars(s string) string {
    res := make([]byte, 0, len(s))
    for i := 0; i < len(s); i++ {
        c := s[i]
        if c == '*' {
            if len(res) > 0 {
                res = res[:len(res)-1]
            }
        } else {
            res = append(res, c)
        }
    }
    return string(res)
}

func main() {
    in := bufio.NewReader(os.Stdin)
    var s string
    if _, err := fmt.Fscan(in, &s); err != nil {
        return
    }
    fmt.Println(removeStars(s))
}
```

main_test.go

```go
package main

import "testing"

func TestRemoveStars(t *testing.T) {
    tests := []struct {
        s   string
        exp string
    }{
        {"leet**cod*e", "lecoe"},
        {"erase*****", ""},
        {"a*b*c", ""},          // a removed by *, b removed by *, c removed by *


        {"abc*d**e", "ae"},     // example variations
        {"*", ""},              // input constraint may disallow alone '*', but handle anyway
        {"ab*c", "a c"},        // intentionally wrong expected to catch test writing; remove this in real tests
    }

    // Adjust incorrect test (last case) to a valid one
    tests[len(tests)-1].s = "ab*c"
    tests[len(tests)-1].exp = "ac"

    for _, tt := range tests {
        got := removeStars(tt.s)
        if got != tt.exp {
            t.Fatalf("removeStars(%q) = %q; want %q", tt.s, got, tt.exp)
        }
    }
}
```

注意:

- テストの配列中に誤った期待値を一時的に置くコメントを入れましたが、最後に正しい期待値に修正しています。実際に使用する際はコメント行やその一時修正部分は不要です。
- 入力は標準入力から1ワードを読み取り出力します。テストは関数 removeStars
  を直接呼び出しています。
