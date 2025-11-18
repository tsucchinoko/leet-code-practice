# max-number-of-ksum-pairs

## 問題の性質

配列から和が k
になる「ペア」を選んで取り除く操作を何回できるかを最大化する問題。各要素は整数で一度使うと消える（重複利用不可）。典型的なハッシュテーブルによる頻度カウントで解ける貪欲問題。要素の順序や配置は重要でなく、組の数を数える組合せ問題に帰着する。

## 問題要約

整数配列 nums と整数 k が与えられる。操作は「配列から和が k になる 2
個の要素を選んで削除する」こと。可能な限り多くの操作を行ったときの操作回数（ペア数）を返す。

例:

- nums = [1,2,3,4], k = 5 → 出力 2 （1+4, 2+3）
- nums = [3,1,3,4,3], k = 6 → 出力 1 （3+3 を 1 回だけ）

## 制約

- 1 <= nums.length <= 10^5
- 1 <= nums[i] <= 10^9
- 1 <= k <= 10^9

入力サイズは最大 100,000 程度なので、O(n)〜O(n log n)
程度のアルゴリズムが望ましい。メモリは数値の種類に依存するが、ハッシュマップで頻度を保持するのが現実的。

## 考え方

考え方の要点を簡潔に述べる。

1. 頻度カウント:
   - nums の各値 x に対して、その出現回数 count[x] を数える（ハッシュマップ）。
2. ペア形成ルール:
   - 2 つの異なる値 a, b（a + b = k, a < b）については、ペア数は min(count[a],
     count[b])。
   - a == b の場合（つまり 2a = k）には、ペア数は floor(count[a] /
     2)（同じ値同士で組を作る）。
3. 実装上の注意:
   - ハッシュマップを走査して各 a について b = k - a
     を見つける。二重計算を避けるために a < b の場合のみ min を加え、a == b
     の場合は floor を加える。
   - 走査中にカウントを減らしていく（双方を 0
     にする）方法でもよい。最終的に合計ペア数を返す。
4. 計算量:
   - 時間計算量 O(n)（頻度カウントとハッシュ走査）、追加のログは不要。
   - 空間計算量 O(m)（m は互いに異なる nums の要素数、最大 n）。

問題の性質は「グリーディ +
ハッシュマップによる頻度集計」で、競技プログラミングの基本的なテクニックで解ける典型問題です。

## キーワード

- ハッシュマップ（map）
- 頻度カウント
- 貪欲（greedy）
- 組合せ（pairing）
- O(n) 時間

---

以下に Go 言語での実装を示します。ファイル構成は指定通り main.go と main_test.go
です。

main.go

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func maxOperations(nums []int, k int) int {
    freq := make(map[int]int)
    for _, x := range nums {
        freq[x]++
    }

    ops := 0
    for value, count := range freq {
        if count == 0 {
            continue
        }

        complement := k - value
        countComplement, ok := freq[complement]
        if !ok || countComplement == 0 {
            continue
        }
        if value < complement {
            // 異なる値同士 (value + complement == k)
            pairs := count
            if countComplement < pairs {
                pairs = countComplement
            }
            ops += pairs
            freq[value] -= pairs
            freq[complement] -= pairs
        } else if value == complement {
            // 同じ値同士 (2 * value == k)
            pairs := count / 2
            ops += pairs
            freq[value] -= pairs * 2
        }
        // value > complement の場合は既に complement 側で処理済みなのでスキップ
    }

    return ops
}

func main() {
    // 入力の簡単なデモ: 標準入力から n, k, nums を読んで結果を出力
    // フォーマット（例）:
    // n k
    // a1 a2 ... an
    in := bufio.NewReader(os.Stdin)
    var n int
    var k int
    if _, err := fmt.Fscan(in, &n, &k); err != nil {
        return
    }
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(in, &nums[i])
    }
    fmt.Println(maxOperations(nums, k))
}
```

main_test.go

```go
package main

import "testing"

func TestMaxOperations(t *testing.T) {
    tests := []struct {
        nums []int
        k    int
        want int
    }{
        {[]int{1, 2, 3, 4}, 5, 2},
        {[]int{3, 1, 3, 4, 3}, 6, 1},
        {[]int{1, 1, 1, 1}, 2, 2},
        {[]int{2, 2, 2, 3, 3, 3}, 5, 3},
        {[]int{1}, 2, 0},
        {[]int{1, 4, 1, 4, 2, 3}, 5, 3},
        {[]int{1000000000, 1}, 1000000001, 1},
    }

    for _, tt := range tests {
        got := maxOperations(tt.nums, tt.k)
        if got != tt.want {
            t.Fatalf("maxOperations(%v, %d) = %d; want %d", tt.nums, tt.k, got, tt.want)
        }
    }
}
```

実行方法:

- go test でテストを実行できます。
- 標準入力から試す場合は、main をビルドしてから入力を渡してください。
