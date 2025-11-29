# count-good-nodes-in-binary-tree

## 問題の性質

二分木の各ノードについて、「根からそのノードまでのパス上にそのノードより大きい値を持つノードが存在しない」ならばそのノードは
good
とする。根から走査して経路上の最大値を保持し、各ノードがその最大値以上かどうかを判定する典型的な木の深さ優先探索（DFS）問題であり、各ノードを一度だけ訪問すれば良いので線形時間で解ける。

## 問題要約

与えられた二分木の根 root に対して、根から各ノード X までのパスに X
より大きい値を持つノードが存在しなければ X は good である。全ての good
ノードの数を返す。

入力例:

- root = [3,1,4,3,null,1,5] -> 出力 4
- root = [3,3,null,4,2] -> 出力 3
- root = [1] -> 出力 1

## 制約

- ノード数 n は 1 ≤ n ≤ 10^5
- 各ノードの値は -10^4 ≤ val ≤ 10^4
- 再帰深さは最悪で木が偏ると n までになるため、言語の再帰制限に注意（Go
  の場合は通常の再帰で通るが、非常に深い木ではスタックに注意）
- 時間計算量の目標: O(n)
- 空間計算量の目標: O(h) （h は木の高さ、再帰または明示的スタックの深さ）

## 考え方

1. 根から出発し、現在までの経路における最大値 currentMax を保持する。
2. 各ノードに到達したら、そのノードの値 val が currentMax 以上であればカウントを
   1 増やす（そのノードは good）。
3. 次の子へ進むときには currentMax を max(currentMax, val) に更新して伝播する。
4. 木全体を
   DFS（深さ優先探索）で訪問する。再帰または明示的スタックのどちらでもよい。各ノードは一度だけ確認するので
   O(n)。
5. 制約上、値の範囲は小さいのでオーバーフローの心配は不要。

要点:

- root 自身は常に good（経路上に自分より大きい値がない）。
- パス上の最大値を子孫に伝播し続けるのが肝。

## キーワード

- 二分木（Binary Tree）
- 深さ優先探索（DFS）
- 再帰（Recursion）
- パス上の最大値伝播
- 線形時間 O(n)

---

以下に Go 言語での実装を示します。ファイル名は指定どおり main.go、テストコードは
main_test.go としてあります。

main.go

```go
package main

import (
    "fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// countGoodNodes returns the number of good nodes in the binary tree.
func countGoodNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    var dfs func(node *TreeNode, currentMax int) int
    dfs = func(node *TreeNode, currentMax int) int {
        if node == nil {
            return 0
        }
        cnt := 0
        if node.Val >= currentMax {
            cnt = 1
            currentMax = node.Val
        }
        cnt += dfs(node.Left, currentMax)
        cnt += dfs(node.Right, currentMax)
        return cnt
    }
    return dfs(root, root.Val)
}

func main() {
    // 簡単な手動テスト
    // root = [3,1,4,3,null,1,5]
    root := &TreeNode{Val: 3}
    root.Left = &TreeNode{Val: 1}
    root.Right = &TreeNode{Val: 4}
    root.Left.Left = &TreeNode{Val: 3}
    root.Right.Left = &TreeNode{Val: 1}
    root.Right.Right = &TreeNode{Val: 5}

    fmt.Println(countGoodNodes(root)) // expect 4
}
```

main_test.go

```go
package main

import "testing"

func buildTreeFromSlice(vals []interface{}) *TreeNode {
    // レベル順配列から二分木を構築するユーティリティ
    // nil は空ノードを示す
    if len(vals) == 0 {
        return nil
    }
    nodes := make([]*TreeNode, len(vals))
    for i, v := range vals {
        if v == nil {
            nodes[i] = nil
        } else {
            nodes[i] = &TreeNode{Val: v.(int)}
        }
    }
    for i := 0; i < len(vals); i++ {
        if nodes[i] == nil {
            continue
        }
        l := 2*i + 1
        r := 2*i + 2
        if l < len(vals) {
            nodes[i].Left = nodes[l]
        }
        if r < len(vals) {
            nodes[i].Right = nodes[r]
        }
    }
    return nodes[0]
}

func TestCountGoodNodes(t *testing.T) {
    tests := []struct {
        vals []interface{}
        want int
    }{
        {[]interface{}{3, 1, 4, 3, nil, 1, 5}, 4},
        {[]interface{}{3, 3, nil, 4, 2}, 3},
        {[]interface{}{1}, 1},
        {[]interface{}{-1, -2, -3, nil, -2, nil, -1}, 3}, // マイナス値の例
    }

    for _, tc := range tests {
        root := buildTreeFromSlice(tc.vals)
        got := countGoodNodes(root)
        if got != tc.want {
            t.Fatalf("for %v: want %d, got %d", tc.vals, tc.want, got)
        }
    }
}
```

補足:

- テストのための buildTreeFromSlice はレベル順配列（nil
  を空ノード）から木を構築する簡易ユーティリティです。LeetCode
  の入力フォーマットに合わせています。
- 非再帰実装が必要なら、スタックで (node, currentMax)
  の組を管理する方法でも同等に実装できます。
