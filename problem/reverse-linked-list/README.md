# reverse-linked-list

## 問題の性質

単方向（singly）連結リストの要素順序を反転する問題。リストのポインタ（next参照）を操作して先頭と末尾を入れ替える操作を行うため、ポインタ操作と境界条件（空リスト、要素数1）の扱いが重要。

## 問題要約

与えられた単方向連結リストの先頭ノード（head）を受け取り、リストを反転して新しい先頭ノードを返す。入力例：

- head = [1,2,3,4,5] → 出力 [5,4,3,2,1]
- head = [1,2] → 出力 [2,1]
- head = [] → 出力 []

反転は反復（イテレーティブ）または再帰（リカーシブ）で実装可能。両方とも実装する。

## 制約

- ノード数は 0 以上 5000 以下。
- 各ノードの値 Node.val は -5000 以上 5000 以下。
- 時間・空間の要件（暗黙）：
  - イテレーティブ解法は O(n) 時間、O(1) 追加空間。
  - 再帰解法は O(n) 時間、再帰呼び出し分の深さで O(n) 空間（コールスタック）。

## 考え方

- イテレーティブ（反復）:
  1. prev = nil, curr = head とする。
  2. curr を走査しながら次 node を保存 next = curr.Next。
  3. curr.Next を prev に向ける（反転）。
  4. prev = curr、curr = next に進める。
  5. 最後に prev が新しい先頭を指すので返す。
  - 境界条件: head が nil（空リスト）や要素数1のときも正しく動作する。

- 再帰（リカーシブ）:
  1. ベースケース: head が nil または head.Next が nil のとき、その head
     を返す（新しい先頭）。
  2. 再帰的に rest := reverse(head.Next) を得る。rest は反転済みリストの先頭。
  3. head.Next.Next = head として現在の head を末尾側に繋ぎ、head.Next = nil
     にして終端にする。
  4. rest を返す。
  - 注意:
    再帰深さはリスト長に比例するため長いリストではスタックオーバーフローの懸念あり（ただし制約内は最大5000で実用可能だが注意が必要）。

## キーワード

- 単方向連結リスト
- ポインタ反転
- イテレーション（反復）
- 再帰（リカーシブ）
- O(n) 時間、O(1) 追加空間（イテレーティブ）
- コールスタック（再帰）

---

以下に Go 言語での具体的な実装を示します。ファイル構成：

- main.go —
  連結リスト定義、反復版と再帰版の実装、簡単なヘルパ関数（スライス⇄リスト変換）と
  main（例）。
- main_test.go — テストコード（table-driven tests）で両実装を検証。

main.go

```go
package main

import (
    "fmt"
)

// ListNode は単方向連結リストのノードを表す。
type ListNode struct {
    Val  int
    Next *ListNode
}

// reverseIterative は反復的に連結リストを反転して新しい先頭を返す。
func reverseIterative(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}

// reverseRecursive は再帰的に連結リストを反転して新しい先頭を返す。
func reverseRecursive(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    rest := reverseRecursive(head.Next)
    head.Next.Next = head
    head.Next = nil
    return rest
}

// helper: sliceToList はスライスから連結リストを作成する。
func sliceToList(nums []int) *ListNode {
    var head, tail *ListNode
    for _, v := range nums {
        node := &ListNode{Val: v}
        if head == nil {
            head = node
            tail = node
        } else {
            tail.Next = node
            tail = node
        }
    }
    return head
}

// helper: listToSlice は連結リストをスライスに変換する。
func listToSlice(head *ListNode) []int {
    res := []int{}
    for cur := head; cur != nil; cur = cur.Next {
        res = append(res, cur.Val)
    }
    return res
}

func main() {
    examples := [][]int{
        {1, 2, 3, 4, 5},
        {1, 2},
        {},
    }
    for _, ex := range examples {
        l := sliceToList(ex)
        r := reverseIterative(l)
        fmt.Printf("iterative %v -> %v\n", ex, listToSlice(r))

        l2 := sliceToList(ex)
        r2 := reverseRecursive(l2)
        fmt.Printf("recursive %v -> %v\n", ex, listToSlice(r2))
    }
}
```

main_test.go

```go
package main

import (
    "reflect"
    "testing"
)

func TestReverseIterative(t *testing.T) {
    tests := []struct {
        name string
        in   []int
        want []int
    }{
        {"empty", []int{}, []int{}},
        {"one", []int{1}, []int{1}},
        {"two", []int{1, 2}, []int{2, 1}},
        {"five", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
        {"negatives", []int{-1, 0, 3}, []int{3, 0, -1}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            head := sliceToList(tt.in)
            got := reverseIterative(head)
            if !reflect.DeepEqual(listToSlice(got), tt.want) {
                t.Fatalf("reverseIterative(%v) = %v, want %v", tt.in, listToSlice(got), tt.want)
            }
        })
    }
}

func TestReverseRecursive(t *testing.T) {
    tests := []struct {
        name string
        in   []int
        want []int
    }{
        {"empty", []int{}, []int{}},
        {"one", []int{1}, []int{1}},
        {"two", []int{1, 2}, []int{2, 1}},
        {"five", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
        {"negatives", []int{-1, 0, 3}, []int{3, 0, -1}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            head := sliceToList(tt.in)
            got := reverseRecursive(head)
            if !reflect.DeepEqual(listToSlice(got), tt.want) {
                t.Fatalf("reverseRecursive(%v) = %v, want %v", tt.in, listToSlice(got), tt.want)
            }
        })
    }
}
```

これでイテレーティブ版と再帰版の両方が確認できるテスト付きの Go
実装が揃います。コンパイル・テスト実行は通常 go test ./...
をプロジェクトルートで実行してください。
