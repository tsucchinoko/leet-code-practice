# path-sum-3

## 問題の性質

木構造（二分木）を扱う典型的な探索・累積和の問題です。パスは「親から子へ下に進む連続したノード列」であり、開始ノードと終了ノードは任意（ルートや葉である必要はない）です。パスの値はノード値の和で、目標和
targetSum と一致するパスの個数を数えます。

## 問題要約

与えられた二分木の中で、親→子の方向に連続して進む任意のパス（開始は任意、終了は任意）で、ノード値の和が
targetSum に等しくなるものの総数を返せ。

例:

- root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8 → 出力 3

## 制約

- ノード数は 0〜1000 の範囲
- 各ノードの値は -10^9 〜 10^9
- targetSum は -1000 〜 1000
- 計算量／メモリについては、ノード数 n に対して O(n) 〜 O(n log n)
  程度を目標にできる（最悪 O(n^2) のアルゴリズムは n=1000
  で許容範囲だが改善可能）

## 考え方

問題は「任意の開始点から下方向に連続した部分経路の和」を数えることなので、単純に全てのノードを開始点として
DFS で下に走査して和を調べる方法（各ノードからの DFS を行う）だと最悪 O(n^2)
になる可能性があります。ただし n ≤ 1000
なので最悪法でも通る場合が多いですが、より効率的に解くためには累積和（prefix
sum）を用います。

考え方（累積和を用いる方法）:

- ルートから現在のノードまでの累積和を currentSum とする。
- ある過去の祖先（もしくは仮想的に長さ0の前の位置）の累積和 prevSum
  が存在して、currentSum - prevSum = targetSum
  なら、その祖先の直後から現在までのパス和が targetSum に等しい。
- したがって、走査中に出会った累積和の出現回数をハッシュマップで保存しておき、現在の
  currentSum に対して currentSum - targetSum
  が何回出現しているかを結果に加える。
- DFS（再帰）で子に進む際に、currentSum
  の出現回数をインクリメントし、帰る際にデクリメントする（バックトラック）ことで正しいカウントを保つ。
- 初期状態として累積和 0 を 1 回持っておく（これによりルートからの経路そのものが
  targetSum を作るケースをカバー）。

時間計算量: O(n)（各ノードを一度ずつ訪問し、ハッシュ操作は平均 O(1)）
空間計算量: O(n)（ハッシュマップと再帰の深さ）

## キーワード

- 二分木（Binary Tree）
- 深さ優先探索（DFS）
- 累積和（Prefix Sum）
- ハッシュマップ（map）
- バックトラック

---

以下に Go 言語での具体的な実装を示します。ファイル構成は指定どおり main.go と
main_test.go です。

main.go

```go
package main

import (
    "fmt"
)

// TreeNode は二分木のノード定義
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// pathSum は targetSum に等しい下向きパスの数を返す
func pathSum(root *TreeNode, targetSum int) int {
    prefix := make(map[int]int)
    prefix[0] = 1
    var dfs func(node *TreeNode, curr int) int
    dfs = func(node *TreeNode, curr int) int {
        if node == nil {
            return 0
        }
        curr += node.Val
        // currentSum - prevSum == targetSum -> prevSum == currentSum - targetSum
        count := prefix[curr-targetSum]
        prefix[curr]++
        count += dfs(node.Left, curr)
        count += dfs(node.Right, curr)
        prefix[curr]--
        return count
    }
    return dfs(root, 0)
}

func main() {
    // 簡単な動作確認 (例1)
    root := &TreeNode{Val: 10}
    root.Left = &TreeNode{Val: 5}
    root.Right = &TreeNode{Val: -3}
    root.Left.Left = &TreeNode{Val: 3}
    root.Left.Right = &TreeNode{Val: 2}
    root.Right.Right = &TreeNode{Val: 11}
    root.Left.Left.Left = &TreeNode{Val: 3}
    root.Left.Left.Right = &TreeNode{Val: -2}
    root.Left.Right.Right = &TreeNode{Val: 1}

    fmt.Println(pathSum(root, 8)) // 期待出力: 3
}
```

main_test.go

```go
package main

import "testing"

func buildTreeFromSlice(vals []interface{}) *TreeNode {
    // 局所的にレベル順（nil を表す場合は nil を入れる）からツリーを構築する簡易関数
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

func TestPathSumExamples(t *testing.T) {
    // Example 1
    root1 := buildTreeFromSlice([]interface{}{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1})
    got1 := pathSum(root1, 8)
    if got1 != 3 {
        t.Fatalf("example1: got %d, want 3", got1)
    }

    // Example 2
    root2 := buildTreeFromSlice([]interface{}{5,4,8,11,nil,13,4,7,2,nil,nil,5,1})
    got2 := pathSum(root2, 22)
    if got2 != 3 {
        t.Fatalf("example2: got %d, want 3", got2)
    }
}

func TestEdgeCases(t *testing.T) {
    // 空の木
    if pathSum(nil, 0) != 0 {
        t.Fatalf("empty tree should return 0")
    }
    // 単一ノードがターゲットに等しい
    single := &TreeNode{Val: 5}
    if pathSum(single, 5) != 1 {
        t.Fatalf("single node equal target should be 1")
    }
    // 負の値を含むケース
    root := buildTreeFromSlice([]interface{}{1, -2, -3, 1, 3, -2, nil, -1})
    if pathSum(root, -1) != 4 {
        t.Fatalf("negative values case: got %d", pathSum(root, -1))
    }
}
```

注意:

- buildTreeFromSlice はテスト用の簡易的なレベル順構築関数です。テスト配列は nil
  を使って欠損ノードを示します。
- main.go の main 関数は簡単な実行例を示します。ユニットテストは go test
  で実行してください。
