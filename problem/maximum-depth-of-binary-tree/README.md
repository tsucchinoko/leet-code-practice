# maximum-depth-of-binary-tree

## 問題の性質

二分木の深さ（根から最も遠い葉までのノード数）を求める根本的で標準的な再帰／反復問題。木構造のトラバース、深さの計算であり、入力は二分木の根ノード。

## 問題要約

与えられた二分木の根 root
に対して、根から最も遠い葉までのノード数（最大深さ）を返す。空の木なら深さは 0。

例:

- root = [3,9,20,null,null,15,7] → 出力 3
- root = [1,null,2] → 出力 2

## 制約

- ノード数は 0〜10^4
- 各ノードの値は -100〜100
- 空の木が与えられる可能性あり
- スタック深さ（再帰）に関してはノード数上限が 10^4
  なので、最悪の連結リスト状の木では再帰深度が 10^4 となり、Go
  のデフォルトスタックが十分であれば問題ないが、深さが深すぎる環境では注意が必要（反復実装を選べる）。

## 考え方

- 再帰（DFS）:
  1. ベース: ノードが nil なら深さ 0 を返す。
  2. 再帰: 左部分木の深さと右部分木の深さをそれぞれ求め、より大きい方に 1
     を足して返す。
  3. 計算量 O(n)、空間計算量（再帰呼び出し）最悪 O(n)（偏った木）、平均 O(log
     n)（平衡木）。

- 反復（BFS / レベル順）:
  1. キューでレベルごとにノードを走査し、レベル数をカウントする。
  2. 各レベルでキューの要素数を取り、その数だけノードを取り出して子を追加する。
  3. 計算量 O(n)、追加の空間 O(n)（キュー）。

どちらも簡単で安全。ここでは分かりやすい再帰実装を採用する。

## キーワード

- 二分木
- 深さ（高さ）
- DFS（深さ優先探索）
- BFS（幅優先探索）
- 再帰
- レベル順走査

---

以下は
Go（main.go）での実装と単体テスト（main_test.go）です。テストはいくつかのケース（空木、単一ノード、左右偏り、標準例）を含みます。

main.go

```go
package main

import "fmt"

// TreeNode は二分木のノード定義（LeetCode 準拠）
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// maxDepth は再帰を使って二分木の最大深さを返す
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftDepth := maxDepth(root.Left)
    rightDepth := maxDepth(root.Right)
    if leftDepth > rightDepth {
        return leftDepth + 1
    }
    return rightDepth + 1
}

func main() {
    // 簡単な実行例
    root := &TreeNode{Val: 3}
    root.Left = &TreeNode{Val: 9}
    root.Right = &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}
    fmt.Println(maxDepth(root)) // => 3
}
```

main_test.go

```go
package main

import "testing"

func buildTree(vals ...interface{}) *TreeNode {
    // ヘルパー: レベル順の配列（nil を表す nil インタフェース）からツリーを構築
    // ここでは簡易実装：配列が空なら nil、インデックス i の左右は 2*i+1, 2*i+2
    n := len(vals)
    if n == 0 {
        return nil
    }
    nodes := make([]*TreeNode, n)
    for i, v := range vals {
        if v == nil {
            nodes[i] = nil
        } else {
            nodes[i] = &TreeNode{Val: v.(int)}
        }
    }
    for i := 0; i < n; i++ {
        if nodes[i] == nil {
            continue
        }
        leftIdx := 2*i + 1
        rightIdx := 2*i + 2
        if leftIdx < n {
            nodes[i].Left = nodes[leftIdx]
        }
        if rightIdx < n {
            nodes[i].Right = nodes[rightIdx]
        }
    }
    return nodes[0]
}

func TestMaxDepth(t *testing.T) {
    tests := []struct {
        name string
        root *TreeNode
        want int
    }{
        {"empty", nil, 0},
        {"single", &TreeNode{Val: 1}, 1},
        {"example1", buildTree(3, 9, 20, nil, nil, 15, 7), 3},
        {"example2", buildTree(1, nil, 2), 2},
        {"left-heavy", &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, 3},
        {"right-heavy", &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}, 3},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := maxDepth(tt.root); got != tt.want {
                t.Fatalf("maxDepth() = %d, want %d", got, tt.want)
            }
        })
    }
}
```
