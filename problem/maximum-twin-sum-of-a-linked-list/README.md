# maximum-twin-sum-of-a-linked-list

## 問題の性質

連結リストの前半と後半を対（ツイン）にして、それぞれの和（ツイン和）を計算し、その最大値を求める問題。リスト長は必ず偶数。

## 問題要約

与えられた偶数長の単方向連結リスト head に対して、i
番目ノード（0-indexed）のツインは (n-1-i)
番目ノードである。各ツインペアのノード値の和を計算し、その最大値を返す。

例:

- head = [5,4,2,1] → ツイン和は 5+1=6、4+2=6 → 最大 6
- head = [4,2,2,3] → ツイン和は 4+3=7、2+2=4 → 最大 7

## 制約

- ノード数 n は偶数で、2 ≤ n ≤ 10^5
- 各ノードの値は 1 ≤ Node.val ≤ 10^5
- 時間・空間の制約を考慮すると、O(n) 時間程度、追加の空間はできれば O(1)〜O(n)
  程度（問題によっては O(1) を目指す）

## 考え方

解法は主に2種類（どちらも O(n) 時間）：

1. 配列に展開する方法（簡潔・実装容易）
   - 連結リストを走査してノード値を配列に格納する（O(n) 時間・O(n) 追加空間）。
   - 配列の左右から二つずつ取り出して和を計算し最大値を更新する（i と n-1-i）。
   - 実装が簡単でバグが少ない。

2. リスト操作（O(1) 追加空間） — 推奨される面もある
   - 連結リストの中間点を見つける（slow/fast ポインタ）。
   - 後半部分を反転する。
   - 先頭と反転した後半を同時に走査し、対応する和の最大値を求める。
   - 最後に後半を元に戻す（必要なら）。
   - メモリ使用量は O(1)（ポインタだけ）、時間は O(n)。
   - 実装は配列法よりやや複雑（反転と中間点検出の注意）。

問題の性質は「配列化すれば簡単」かつ「リストを操作すれば追加空間を抑えられる」二択で、制約（n
≤ 1e5）からどちらでも十分高速。競技プログラミングや面接では O(1)
空間解（リスト反転）を好むことが多い。

## キーワード

- 連結リスト（singly linked list）
- two pointers（slow/fast）
- リスト反転（reverse linked list）
- 配列化（copy to array）
- 対称インデックス（i と n-1-i）
- O(n) 時間、O(1)/O(n) 空間

---

以下に Go 言語での具体的実装を示します。main.go は解法の実装、main_test.go
はテストコードです。ここでは配列を使う簡潔な解法を採用します（読みやすく安全）。必要ならリスト反転による
O(1) 空間版も提示可能です。

ファイル: main.go

```go
package main

import "fmt"

// ListNode は単方向リンクリストのノード定義
type ListNode struct {
    Val  int
    Next *ListNode
}

// pairSumReverse はリスト反転を使った O(1) 追加空間の実装
func pairSumReverse(head *ListNode) int {
    if head == nil || head.Next == nil {
        return 0
    }

    // 1. 中間点を見つける（slow が前半終了、slow.Next が後半先頭）
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    // リスト長は偶数なので、slow は n/2 番目（0-indexed だと n/2）
    // 後半の先頭は slow
    second := reverseList(slow)

    // 2. head と second を同時走査して最大ツイン和を求める
    p1, p2 := head, second
    maxSum := 0
    for p2 != nil {
        s := p1.Val + p2.Val
        if s > maxSum {
            maxSum = s
        }
        p1 = p1.Next
        p2 = p2.Next
    }

    // 3. 元に戻す（オプションだが良い慣習）
    reverseList(second)

    return maxSum
}

// reverseList は単方向リストの反転を行う（頭を返す）
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    cur := head
    for cur != nil {
        next := cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }
    return prev
}

// ヘルパー: スライスから連結リストを作る（テスト用）
func buildList(vals []int) *ListNode {
    if len(vals) == 0 {
        return nil
    }
    head := &ListNode{Val: vals[0]}
    cur := head
    for i := 1; i < len(vals); i++ {
        cur.Next = &ListNode{Val: vals[i]}
        cur = cur.Next
    }
    return head
}

func main() {
    head := buildList([]int{5, 4, 2, 1})
    fmt.Println(pairSumReverse(head)) // 6
}
```

ファイル: main_test.go

```go
package main

import "testing"

func TestPairSumExamples(t *testing.T) {
    tests := []struct {
        vals []int
        want int
    }{
        {[]int{5, 4, 2, 1}, 6},
        {[]int{4, 2, 2, 3}, 7},
        {[]int{1, 100000}, 100001},
        {[]int{1, 2, 3, 4, 5, 6}, 7}, // pairs: (1+6)=7,(2+5)=7,(3+4)=7
        {[]int{10, 20, 30, 40}, 50},  // (10+40)=50,(20+30)=50
    }

    for _, tt := range tests {
        head := buildList(tt.vals)
        got := pairSum(head)
        if got != tt.want {
            t.Fatalf("vals=%v: got %d, want %d", tt.vals, got, tt.want)
        }
    }
}
```

必要であれば、追加で O(1)
追加空間のリスト反転バージョンの実装も提供できます。どちらを優先しますか？
