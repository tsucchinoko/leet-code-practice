# lowest-common-ancestor-of-a-binary-tree

## 問題の性質

- 木（Binary Tree）上の探索・再帰問題。
- Lowest Common Ancestor（LCA）を求める典型問題。
- 与えられるノードは木中に必ず存在し、ノードの値は全て一意。

## 問題要約

二分木と、木中の2つのノード p, q
が与えられる。LCA（最も低い共通祖先）を返す。ノードは自分自身の子孫とみなせるため、片方のノードがもう片方の祖先であればその祖先が答えになる。

入力例（概念）：

- root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1 → 出力 3
- root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4 → 出力 5

## 制約

- ノード数 n は 2 ≤ n ≤ 10^5 の範囲。
- ノードの値は一意で、-10^9 ≤ Node.val ≤ 10^9。
- p != q。
- p, q は木に必ず存在する。

時間・空間の目安：

- 再帰的解法で O(n) 時間、再帰深さは木の高さ（最悪 O(n)）。
- 追加メモリをほとんど使わない（再帰のスタックを除き O(1)）解が可能。

## 考え方

考え方（簡潔に）：

1. 根から再帰的に探索する。
2. 各ノードで左右の部分木から p または q が見つかるかを調べる。
3. 現在のノードが p または q
   のどちらかであれば「そのノードを見つけた」と報告する。
4. 左右両方の部分木から報告が返ってきたら、現在のノードが LCA（左右それぞれに p
   と q が存在するため）。
5. 片側のみ報告があれば、その報告をそのまま上に返す（LCA はその側に存在する）。
6. どちらからも報告が無ければ nil を返す。

ポイント：

- 再帰終了条件：現在のノードが nil のとき nil を返す。現在のノードが p または q
  のときそのノードを返す。
- 二分探索や親ポインタを事前に作る工夫（例えば parent
  map）も可能だが、再帰1回で解けるシンプルな方法が一般的に最速で直感的。

直感的証明：

- ノード x を訪問したとき、左右から得られる結果は「その部分木に p または q
  が含まれるか」を表す。左右から両方返ってくれば x
  は最も低い（最深）共通地点。片側のみなら LCA はその側にある。

計算量：

- 時間 O(n)：各ノードを最大1回訪問。
- 空間 O(h)：再帰スタックの深さ h（木の高さ）。

## キーワード

- Lowest Common Ancestor (LCA)
- Binary Tree
- 深さ優先探索（DFS）
- 再帰
- Divide and Conquer

---

以下に Go
言語での具体実装（main.go）とテストコード（main_test.go）を示します。実装は標準的な再帰解法です。

main.go

```go
package main

import "fmt"

// TreeNode definition
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// lowest common ancestor
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    // If root is either p or q, root is part of the path to p or q
    if root == p || root == q {
        return root
    }
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)

    if left != nil && right != nil {
        // p and q found in different subtrees
        return root
    }
    if left != nil {
        return left
    }
    return right
}

// helper to build tree from slice (nil represented by sentinel with ok flag)
func buildTree(vals []interface{}) *TreeNode {
    if len(vals) == 0 || vals[0] == nil {
        return nil
    }
    nodes := make([]*TreeNode, len(vals))
    for i, v := range vals {
        if v != nil {
            nodes[i] = &TreeNode{Val: v.(int)}
        } else {
            nodes[i] = nil
        }
    }
    for i := 0; i < len(vals); i++ {
        if nodes[i] == nil {
            continue
        }
        leftIdx := 2*i + 1
        rightIdx := 2*i + 2
        if leftIdx < len(vals) {
            nodes[i].Left = nodes[leftIdx]
        }
        if rightIdx < len(vals) {
            nodes[i].Right = nodes[rightIdx]
        }
    }
    return nodes[0]
}

func main() {
    // Example usage
    root := buildTree([]interface{}{3,5,1,6,2,0,8,nil,nil,7,4})
    // locate nodes p and q by value (since values are unique)
    var p, q *TreeNode
    var dfs func(*TreeNode)
    dfs = func(n *TreeNode) {
        if n == nil {
            return
        }
        if n.Val == 5 {
            p = n
        }
        if n.Val == 1 {
            q = n
        }
        dfs(n.Left)
        dfs(n.Right)
    }
    dfs(root)
    lca := lowestCommonAncestor(root, p, q)
    if lca != nil {
        fmt.Println("LCA:", lca.Val)
    } else {
        fmt.Println("LCA not found")
    }
}
```

main_test.go

```go
package main

import "testing"

func findByVal(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Val == val {
        return root
    }
    if left := findByVal(root.Left, val); left != nil {
        return left
    }
    return findByVal(root.Right, val)
}

func TestLowestCommonAncestor_Example1(t *testing.T) {
    root := buildTree([]interface{}{3,5,1,6,2,0,8,nil,nil,7,4})
    p := findByVal(root, 5)
    q := findByVal(root, 1)
    got := lowestCommonAncestor(root, p, q)
    if got == nil || got.Val != 3 {
        t.Fatalf("expected LCA val 3, got %v", got)
    }
}

func TestLowestCommonAncestor_Example2(t *testing.T) {
    root := buildTree([]interface{}{3,5,1,6,2,0,8,nil,nil,7,4})
    p := findByVal(root, 5)
    q := findByVal(root, 4)
    got := lowestCommonAncestor(root, p, q)
    if got == nil || got.Val != 5 {
        t.Fatalf("expected LCA val 5, got %v", got)
    }
}

func TestLowestCommonAncestor_Example3(t *testing.T) {
    root := buildTree([]interface{}{1,2})
    p := findByVal(root, 1)
    q := findByVal(root, 2)
    got := lowestCommonAncestor(root, p, q)
    if got == nil || got.Val != 1 {
        t.Fatalf("expected LCA val 1, got %v", got)
    }
}

// additional tests for edge cases
func TestLowestCommonAncestor_Chain(t *testing.T) {
    // chain 1-2-3-4 (as left children)
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 2}
    root.Left.Left = &TreeNode{Val: 3}
    root.Left.Left.Left = &TreeNode{Val: 4}
    p := findByVal(root, 3)
    q := findByVal(root, 4)
    got := lowestCommonAncestor(root, p, q)
    if got == nil || got.Val != 3 {
        t.Fatalf("expected LCA val 3, got %v", got)
    }
}
```

注意：

- buildTree はレベル順配列（nil
  表示で空ノード）から簡易的に木を構築します。テストや例示用途向けです。
- 提示した実装は値が一意である前提で findByVal を使って p, q
  を取得していますが、実際の問題ではノードポインタが直接与えられます。
