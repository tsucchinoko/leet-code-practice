# delete-the-middle-node-of-a-linked-list

## 問題の性質

単方向連結リストのノードを1つ削除する操作。リスト長 n に対して 0-index で
floor(n/2)
番目のノード（中央ノード）を削除して新しい先頭を返す。線形走査で解けるシンプルなリスト操作問題。アルゴリズムと実装の正しさ・境界条件（n=1,
n=2 など）がポイント。

## 問題要約

与えられた単方向連結リストの中央ノード（0始まりで
⌊n/2⌋）を削除して、削除後のリストの先頭を返す。リスト長は 1
以上。削除対象はリストの途中・末尾・先頭になる場合があり得る（ただし n≥1
なので先頭のみが残ることはありうる）。

例:

- [1,3,4,7,1,2,6] → 中央は index 3 の 7 → 結果 [1,3,4,1,2,6]
- [1,2,3,4] → 中央は index 2 の 3 → 結果 [1,2,4]
- [2,1] → 中央は index 1 の 1 → 結果 [2]

## 制約

- ノード数 n は 1 ≤ n ≤ 10^5
- 各ノードの値は 1 ≤ val ≤ 10^5
- 時間計算量は O(n) が実用的（単一走査または定数回の走査）
- 空間計算量は O(1)（追加メモリは定数オーダー）で十分

## 考え方

要点を簡潔にまとめると以下の2通りの方針がある。

1. 2回走査（簡潔で実装が容易）
   - 1回目でリスト長 n を数える。
   - 削除すべきインデックス m = ⌊n/2⌋ を計算する。
   - 2回目で m 番目のノードを探し、それを削除する（前のノードの next を変更）。
   - 特殊ケース: n=1 のときは空リスト（nil）を返す。

   長所: 実装が分かりやすくバグが起きにくい。短所: リストを2回走査する（ただし
   O(n)）。

2. 1回走査（双ポインタ法：遅いポインタと速いポインタ）
   - fast ポインタを2ステップ、slow
     を1ステップ進めることでリスト中央を検出する。
   - ただし削除には「slow の1つ前（prev）」が必要なので prev
     ポインタを保持する。
   - 初期化例: prev = nil, slow = head, fast = head
   - ループ: fast と fast.next が存在する間、prev = slow; slow = slow.next; fast
     = fast.next.next
   - 終了時、slow が削除すべき中央ノード。prev が nil なら（n=1） head
     を削除して nil を返す。そうでなければ prev.next = slow.next。
   - 長所: 1回走査で O(n)。短所: 実装で prev の扱いに注意（オフバイワン）。

問題の性質は単純なリスト操作（中央検出と削除）で、典型的には双ポインタ法で1回走査か、シンプルに長さを数えて2回走査する実装が多い。

### 具体的なイメージ（簡単な例）

具体的なイメージ（簡単な例）

- リスト: A -> B -> C -> D
  - head は A を指す。
  - A.Next = B, B.Next = C, C.Next = D, D.Next = nil
- 中央ノードを削除するために prev = B（A->B->C の prev が B, slow = C）
- prev.Next = slow.Next を実行すると、B.Next = D になる。
- 結果のリストは A -> B -> D（head は引き続き A を指す）

なぜ head が正しく見えるのか

- head は A のアドレスを保持しているだけ。list の「次に何が来るか」は A.Next
  の値に依存します。
- A.Next は削除操作で変わっていない（A.Next = B）。しかし B.Next（B の
  Next）を変更することで A を起点にたどった場合に見えるノード列が変わる（C
  が消えて D に繋がる）。
- つまり、head 自体を再割り当てする必要はなく、内部の Next
  ポインタを書き換えれば先頭から見える構造が更新される。

より一般的に

- 単方向連結リストにおける「削除」は、削除対象 node
  を指している変数を変えるのではなく、その直前のノード prev の Next を node.Next
  に変更する操作です。
- これによって prev
  の後に続くノード（削除対象）は探索経路から外れ、ガベージコレクタに回る（Go
  の場合）か破棄されます。
- head が指す先頭ノードに変更が無ければ、呼び出し元に返す head
  はそのままで良く、たどった結果が削除後のリストになります。

特殊ケースの補足

- 削除対象が先頭（head）自身の場合（本問題では n=1 のときなど）、prev
  が存在しない（nil）なので prev.Next を書けません。このため事前に
  head.Next==nil をチェックして head を nil にする、あるいは別途 head =
  head.Next（先頭を置き換える）といった処理が必要です。
- 本問題の実装では n>=2 のケースで prev が必ず non-nil
  になるようにしているため、prev.Next = slow.Next が安全に使えます。

短くまとめると

- head は先頭ノードを指すだけで、そのノード内の Next フィールドを変えれば「head
  から見たリストの見え方」が変わる。
- prev.Next = slow.Next は「prev
  の次に来るノード」を差し替える操作であり、それだけで head
  を先頭とした連結関係が更新されるため、返す head
  を別に変更する必要がない、ということです。

## キーワード

- 単方向連結リスト（Singly linked list）
- 中央ノード（middle node）
- 双ポインタ（fast/slow pointer）
- 1回走査 / 2回走査
- 前ノード（prev）
- O(n) 時間、O(1) 空間

---

以下に Go
言語での具体的実装（main.go）とテストコード（main_test.go）を示します。実装は双ポインタ法（1回走査）で書きつつ、境界条件に対応しています。

main.go

```go
package main

import "fmt"

// ListNode defines a singly-linked list node.
type ListNode struct {
    Val  int
    Next *ListNode
}

// deleteMiddle deletes the middle node (0-indexed floor(n/2)) and returns the head of the modified list.
func deleteMiddle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    // If only one node, deleting it yields empty list.
    if head.Next == nil {
        return nil
    }

    var prev *ListNode
    slow := head
    fast := head

    // Move fast by 2 and slow by 1, keeping prev one step behind slow.
    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }

    // slow is the middle node to delete, prev is its previous node (guaranteed non-nil because n>=2)
    prev.Next = slow.Next
    return head
}

// helper: build list from slice
func buildList(vals []int) *ListNode {
    if len(vals) == 0 {
        return nil
    }
    head := &ListNode{Val: vals[0]}
    cur := head
    for _, v := range vals[1:] {
        cur.Next = &ListNode{Val: v}
        cur = cur.Next
    }
    return head
}

// helper: convert list to slice
func toSlice(head *ListNode) []int {
    res := []int{}
    for cur := head; cur != nil; cur = cur.Next {
        res = append(res, cur.Val)
    }
    return res
}

func main() {
    cases := [][]int{
        {1, 3, 4, 7, 1, 2, 6},
        {1, 2, 3, 4},
        {2, 1},
        {1},
    }
    for _, c := range cases {
        head := buildList(c)
        newHead := deleteMiddle(head)
        fmt.Println(toSlice(newHead))
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

func TestDeleteMiddle(t *testing.T) {
    tests := []struct {
        in   []int
        want []int
    }{
        {[]int{1, 3, 4, 7, 1, 2, 6}, []int{1, 3, 4, 1, 2, 6}},
        {[]int{1, 2, 3, 4}, []int{1, 2, 4}},
        {[]int{2, 1}, []int{2}},
        {[]int{1}, []int{}},
        {[]int{1, 2, 3}, []int{1, 3}}, // n=3 middle index=1
        {[]int{1, 2}, []int{1}},       // n=2 middle index=1
    }

    for _, tt := range tests {
        head := buildList(tt.in)
        gotHead := deleteMiddle(head)
        got := toSlice(gotHead)
        if !reflect.DeepEqual(got, tt.want) {
            t.Fatalf("deleteMiddle(%v) = %v; want %v", tt.in, got, tt.want)
        }
    }
}
```

これでテスト `go test`
を実行すると、与えたケースで正しく中央ノードが削除されることを検証できます。
