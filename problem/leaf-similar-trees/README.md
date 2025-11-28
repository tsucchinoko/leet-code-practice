# leaf-similar-trees

## 問題の性質

二分木の葉ノード（左右子がともに存在しないノード）の値を左から右へ順に並べた列（leaf
value
sequence）が等しいかどうかを判定する問題。木の深さ順や構造自体は問わず、葉の値の順序のみ比較する決定問題（判定問題）で、木の全走査が必要なため線形時間で解ける典型的な探索問題。

## 問題要約

2本の二分木 root1, root2
が与えられる。各木について左から右へ葉ノードの値を列挙したシーケンスが一致すれば
true を返す。そうでなければ false を返す。

例:

- root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 =
  [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8] → true
- root1 = [1,2,3], root2 = [1,3,2] → false

## 制約

- 各木のノード数は 1 〜 200 の範囲。
- 各ノードの値は 0 〜 200 の範囲。
- メモリ・時間ともに木のノード数に対して線形（O(n)）で十分。

## 考え方

- 各木について深さ優先探索（DFS）または幅優先探索（BFS）で全ノードを走査し、葉ノードに遭遇した順にその値をスライス（配列）へ追加する。
- 2つの葉列が長さ・各要素ともに一致するか比較する。等しければ
  true、そうでなければ false。
- DFS を用いる場合、再帰で左→右の順に訪問すれば左から右の葉順を得られる。
- 再帰が心配ならスタックを用いた反復 DFS
  でも同様に左→右を保つように右子を先に積むなどの工夫をする。
- 時間計算量: O(n1 + n2)（各木を一度ずつ走査）
- 空間計算量: O(L1 + L2)（葉の数分の配列、最悪ノード数に比例）

問題の性質: 木の走査・列挙・列比較（簡潔な線形判定問題）

## キーワード

- 二分木（Binary Tree）
- 葉ノード（Leaf Node）
- 深さ優先探索（DFS）
- シーケンス比較
- 線形時間

---

以下に Go 言語での実装を示します。ファイル構成:

- main.go — 実装
- main_test.go — テストコード（例題といくつかの追加ケース）

注意: 標準の LeetCode 風の Tree 表現をテストでは手作りで構築しています。

main.go

```go
package main

import "fmt"

// TreeNode is binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// leafSimilar returns true if two trees have the same leaf value sequence.
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    a := leaves(root1)
    b := leaves(root2)
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

// leaves returns leaf value sequence (left-to-right) for a tree.
func leaves(root *TreeNode) []int {
    res := make([]int, 0)
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        if node.Left == nil && node.Right == nil {
            res = append(res, node.Val)
            return
        }
        dfs(node.Left)
        dfs(node.Right)
    }
    dfs(root)
    return res
}

func main() {
    // 簡単なデモ
    // root1 = [1,2,3]
    r1 := &TreeNode{Val: 1}
    r1.Left = &TreeNode{Val: 2}
    r1.Right = &TreeNode{Val: 3}

    // root2 = [1,3,2]
    r2 := &TreeNode{Val: 1}
    r2.Left = &TreeNode{Val: 3}
    r2.Right = &TreeNode{Val: 2}

    fmt.Println(leafSimilar(r1, r2)) // false
}
```

main_test.go

```go
package main

import "testing"

func TestLeafSimilar_Example1(t *testing.T) {
    // root1 = [3,5,1,6,2,9,8,null,null,7,4]
    root1 := &TreeNode{Val: 3}
    root1.Left = &TreeNode{Val: 5}
    root1.Right = &TreeNode{Val: 1}
    root1.Left.Left = &TreeNode{Val: 6}
    root1.Left.Right = &TreeNode{Val: 2}
    root1.Left.Right.Left = &TreeNode{Val: 7}
    root1.Left.Right.Right = &TreeNode{Val: 4}
    root1.Right.Left = &TreeNode{Val: 9}
    root1.Right.Right = &TreeNode{Val: 8}

    // root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
    root2 := &TreeNode{Val: 3}
    root2.Left = &TreeNode{Val: 5}
    root2.Right = &TreeNode{Val: 1}
    root2.Left.Left = &TreeNode{Val: 6}
    root2.Left.Right = &TreeNode{Val: 7}
    root2.Right.Left = &TreeNode{Val: 4}
    root2.Right.Right = &TreeNode{Val: 2}
    root2.Right.Right.Left = &TreeNode{Val: 9}
    root2.Right.Right.Right = &TreeNode{Val: 8}

    if !leafSimilar(root1, root2) {
        t.Fatalf("expected true for example1")
    }
}

func TestLeafSimilar_Example2(t *testing.T) {
    r1 := &TreeNode{Val: 1}
    r1.Left = &TreeNode{Val: 2}
    r1.Right = &TreeNode{Val: 3}

    r2 := &TreeNode{Val: 1}
    r2.Left = &TreeNode{Val: 3}
    r2.Right = &TreeNode{Val: 2}

    if leafSimilar(r1, r2) {
        t.Fatalf("expected false for example2")
    }
}

func TestLeafSimilar_SingleNode(t *testing.T) {
    a := &TreeNode{Val: 5}
    b := &TreeNode{Val: 5}
    if !leafSimilar(a, b) {
        t.Fatalf("single-node equal failed")
    }
    c := &TreeNode{Val: 6}
    if leafSimilar(a, c) {
        t.Fatalf("single-node different failed")
    }
}

func TestLeafSimilar_DifferentLengths(t *testing.T) {
    // left tree leaves: [1,2]
    a := &TreeNode{Val: 0}
    a.Left = &TreeNode{Val: 1}
    a.Right = &TreeNode{Val: 2}

    // right tree leaves: [1]
    b := &TreeNode{Val: 1}

    if leafSimilar(a, b) {
        t.Fatalf("different leaf counts should be false")
    }
}
```

この実装は DFS
を用いて葉を左から右へ集め、得られた2つの配列を比較しています。ノード数が最大200
と小さいので再帰深さも問題になりにくいですが、深い木が予想される場合は反復実装に変更してもよいです。
