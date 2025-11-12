# string-compression

## 問題の性質

配列内で連続する同一文字の塊を走査して、その塊を文字と出現回数で置き換える「ランレングス圧縮」に相当します。ただし、圧縮した結果を新しい配列に返すのではなく、与えられた配列
`chars`
をその場で書き換え、一定の長さ（新しい長さ）を返すインプレース問題です。追加で使えるメモリは定数サイズに限られます。

## 問題要約

与えられた文字配列 `chars`
を次のルールで圧縮して入力配列に上書きし、圧縮後の有効長さを返す。

- 連続する同一文字のグループごとに処理する。
- グループ長が1なら文字のみを残す。
- グループ長が2以上なら、文字の後にグループ長の数値（桁ごとに文字）を追加する。
- 数字は10以上の場合も桁ごとに分けて格納する（例: 12 -> '1','2'）。
- 返すのは新しい長さだけで、配列のそれ以降の要素は無視される。

例:

- ["a","a","b","b","c","c","c"] -> ["a","2","b","2","c","3"], 長さ6
- ["a"] -> ["a"], 長さ1
- ["a","b", ... 12個のb ...] -> ["a","b","1","2"], 長さ4

## 制約

- 1 <= chars.length <= 2000
- chars[i] は英小文字・英大文字・数字・記号のいずれか
- 追加で使えるメモリは O(1)（定数）であること
- 出力配列は入力配列を書き換えること（インプレース）

## 考え方

アルゴリズムは二つのインデックスを用いる「読み取りポインタ」と「書き込みポインタ」を使った一回走査（O(n)）で実現できます。考え方は次の通りです。

1. 読み取りポインタ `i` を先頭から進め、現在のグループの先頭文字 `chars[i]`
   を取る。
2. 同一文字が連続する限り `i` を進めてカウント `count` を得る（右へスキャン）。
3. 書き込みポインタ `write` の位置にその文字を書き込む。`write` を1進める。
4. `count` が1より大きければ、`count` を文字列に変換し、その各桁を
   `chars[write]` に順に書き込み、`write` を進める。
5. 次のグループに移るまで手順1〜4を繰り返す。
6. 最終的に `write` が新しい長さであるためそれを返す。

重要点・注意点:

- 数を文字に分解するために一時的に文字列に変換して桁を順に書き込む（この変換は定数メモリの範囲内）。
- 走査は前から後ろへ一度だけ行うため時間計算量は O(n)。
- 追加で使うメモリは固定長の変数のみ（カウント、ポインタ、局所的な数値→文字列変換のための一時バッファ）で
  O(1)。

## キーワード

- ランレングス圧縮（Run-Length Encoding）
- インプレース書き換え
- 2ポインタ法（読み取りポインタ／書き込みポインタ）
- O(n) 時間、O(1) 追加空間
- 桁分解（数値 -> 文字列）

---

以下に Go 言語での実装を示します。ファイル構成は要求に従い `main.go`（解法）と
`main_test.go`（ユニットテスト）です。

main.go

```go
package main

import (
    "fmt"
    "strconv"
)

func compress(chars []byte) int {
    n := len(chars)
    write := 0
    i := 0

    for i < n {
        ch := chars[i]
        j := i
        for j < n && chars[j] == ch {
            j++
        }
        count := j - i

        // write character
        chars[write] = ch
        write++

        // if count > 1, write the digits of count
        if count > 1 {
            s := strconv.Itoa(count)
            for k := 0; k < len(s); k++ {
                chars[write] = s[k]
                write++
            }
        }

        // move to next group
        i = j
    }

    return write
}

func main() {
    // example run
    arr := []byte{'a','a','b','b','c','c','c'}
    newLen := compress(arr)
    fmt.Printf("newLen=%d, compressed=%q\n", newLen, arr[:newLen])
}
```

main_test.go

```go
package main

import (
    "reflect"
    "testing"
)

func toByteSlice(ss []string) []byte {
    b := make([]byte, len(ss))
    for i := range ss {
        b[i] = ss[i][0]
    }
    return b
}

func sliceEqualBytes(a []byte, b []byte) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func TestCompressExamples(t *testing.T) {
    tests := []struct {
        input    []byte
        want     []byte
        wantLen  int
    }{
        {[]byte{'a','a','b','b','c','c','c'}, []byte{'a','2','b','2','c','3'}, 6},
        {[]byte{'a'}, []byte{'a'}, 1},
        {[]byte{'a','b','b','b','b','b','b','b','b','b','b','b','b'}, []byte{'a','b','1','2'}, 4},
        {[]byte{'a','a','a','a','a','a','a','a','a','a'}, []byte{'a','1','0'}, 3}, // 10 a's
        {[]byte{'a','b','c'}, []byte{'a','b','c'}, 3},
    }

    for _, tt := range tests {
        // copy input to avoid mutating test case data across runs
        inp := make([]byte, len(tt.input))
        copy(inp, tt.input)
        gotLen := compress(inp)
        if gotLen != tt.wantLen {
            t.Fatalf("compress(%q) returned length %d, want %d", tt.input, gotLen, tt.wantLen)
        }
        if !sliceEqualBytes(inp[:gotLen], tt.want) {
            t.Fatalf("compress(%q) = %q (len %d), want %q (len %d)", tt.input, inp[:gotLen], gotLen, tt.want, tt.wantLen)
        }
    }
}

func TestCompressRandom(t *testing.T) {
    // basic additional sanity: repeating groups and singletons
    inp := []byte{'z','z','z','y','x','x','1','1','1','1'}
    want := []byte{'z','3','y','x','2','1','4'}
    gotLen := compress(inp)
    if gotLen != len(want) {
        t.Fatalf("length got %d want %d", gotLen, len(want))
    }
    if !reflect.DeepEqual(inp[:gotLen], want) {
        t.Fatalf("got %q want %q", inp[:gotLen], want)
    }
}
```

注意:

- `compress` 関数は与えられた `[]byte`
  スライスを直接書き換え、新しい有効長を返します。
- テストでは入力スライスのコピーを使って関数を呼び出し、期待結果と長さを検証しています。
