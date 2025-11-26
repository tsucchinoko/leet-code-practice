# odd-even-linked-list

## 問題の性質

単方向連結リストのノードを「位置（インデックス）が奇数のノード群」と「偶数のノード群」に分け、奇数群を先、偶数群を後に連結して新しい順序のリストを返す問題。リスト内のノード順序は各群内で入力順を保つ必要がある。追加のメモリを定数オーダー（O(1)）に抑え、時間計算量は一次（O(n)）で解く必要がある。

## 問題要約

入力：連結リスト head（先頭ノードがインデックス1として数える）\
出力：インデックスが奇数のノードを先に並べ、次にインデックスが偶数のノードを並べた同じノードを再利用したリスト。ノードの相対順序は各グループ内で保持する。

例：

- [1,2,3,4,5] -> [1,3,5,2,4]
- [2,1,3,5,6,4,7] -> [2,3,6,7,1,5,4]

## 制約

- ノード数 n は 0 <= n <= 10,000
- 各ノードの値は -10^6 <= val <= 10^6
- 時間計算量は O(n)
- 追加空間は O(1)（再帰による追加深さも含めて許容されない）
- リストは単方向（next ポインタのみ）

## 考え方

1. 目的：元のリストを走査して、奇数位置ノードだけで作るリスト（odd）と偶数位置ノードだけで作るリスト（even）をその場で分割し、最後に
   odd の末尾に even
   を繋げる。新しいノードを作らず、ポインタ操作だけで実現するため追加メモリは定数。
2. 初期化：
   - odd := head
   - even := head.next (ある場合)
   - evenHead := even（後で even リストの先頭を保持しておく）
3. 反復処理（while even != nil && even.next != nil）:
   - odd.next = even.next // 次の奇数ノードを繋ぐ
   - odd = odd.next // odd を進める
   - even.next = odd.next // 次の偶数ノードを繋ぐ
   - even = even.next // even を進める これにより in-place
     で奇数群と偶数群がそれぞれ連結されていく。
4. 終了後、odd.next = evenHead で偶数群を奇数群の末尾に繋げる。
5. これで head を返す（空リストや要素数が1,2 の場合にも安全に処理される）。

正しさの直感：

- ループ不変量：現在の odd はこれまで見た奇数位置の末尾、even
  はこれまで見た偶数位置の末尾。ループは次の偶数・奇数の存在を確認してから繋ぎ替えるため順序が保たれる。
- 各ノードは一定回数だけ next を更新するため O(n)
  時間。追加メモリはポインタ変数のみ。

エッジケース：

- head == nil -> nil を返す
- head.next == nil -> 1要素のみ -> そのまま返す
- 2要素 -> odd は head、evenHead は head.next、ループは走らず odd.next は
  evenHead のまま -> そのまま返る

## キーワード

- 連結リスト（Singly linked list）
- ポインタ操作（in-place）
- 奇数・偶数インデックス分割
- O(1) 追加空間、O(n) 時間
- two pointers（odd / even）

---

以下に Go 言語での実装を示します。ファイル名は指定どおり main.go、テストは
main_test.go です。

main.go

```go
package main

import (
    "fmt"
)

// ListNode singly-linked list node
type ListNode struct {
    Val  int
    Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    odd := head
    even := head.Next
    evenHead := even

    for even != nil && even.Next != nil {
        odd.Next = even.Next
        odd = odd.Next

        even.Next = odd.Next
        even = even.Next
    }

    odd.Next = evenHead
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
    for head != nil {
        res = append(res, head.Val)
        head = head.Next
    }
    return res
}

func main() {
    examples := [][]int{
        {1, 2, 3, 4, 5},
        {2, 1, 3, 5, 6, 4, 7},
        {},
        {1},
        {1, 2},
    }
    for _, ex := range examples {
        l := buildList(ex)
        out := oddEvenList(l)
        fmt.Println(toSlice(out))
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

func TestOddEvenList(t *testing.T) {
    cases := []struct {
        in  []int
        out []int
    }{
        {in: []int{1, 2, 3, 4, 5}, out: []int{1, 3, 5, 2, 4}},
        {in: []int{2, 1, 3, 5, 6, 4, 7}, out: []int{2, 3, 6, 7, 1, 5, 4}},
        {in: []int{}, out: []int{}},
        {in: []int{1}, out: []int{1}},
        {in: []int{1, 2}, out: []int{1, 2}},
        {in: []int{1, 2, 3}, out: []int{1, 3, 2}},
    }

    for _, c := range cases {
        got := oddEvenList(buildList(c.in))
        gotSlice := toSlice(got)
        if !reflect.DeepEqual(gotSlice, c.out) {
            t.Fatalf("oddEvenList(%v) = %v; want %v", c.in, gotSlice, c.out)
        }
    }
}
```

説明：

- buildList / toSlice はテストやデモのためのユーティリティです。
- oddEvenList は in-place でポインタをつなぎ替え、O(1) の追加空間、O(n)
  の時間で動作します。
