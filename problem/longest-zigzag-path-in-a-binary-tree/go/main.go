package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	var ans int
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		// ここで子の戻り値を必ず左・右の順で受け取る
		lLeft, lRight := dfs(node.Left)
		rLeft, rRight := dfs(node.Right)

		// leftLen: 直前が右だった場合（次は左へ行く）の最大長
		leftLen := 1 + lRight
		// rightLen: 直前が左だった場合（次は右へ行く）の最大長
		rightLen := 1 + rLeft

		if leftLen > ans {
			ans = leftLen
		}
		if rightLen > ans {
			ans = rightLen
		}
		// rRight はこの実装では使わないが、上のように受け取っているため未使用変数エラーは起きない
		_ = rRight
		_ = lLeft

		return leftLen, rightLen
	}

	dfs(root)
	if ans < 0 {
		return 0
	}
	return ans - 1
}

func main() {
	// ---- sample 1: known zigzag length 3 (right->left->right)
	a := &TreeNode{Val: 1}
	a.Right = &TreeNode{Val: 1}
	a.Right.Left = &TreeNode{Val: 1}
	a.Right.Right = &TreeNode{Val: 1}
	a.Right.Left.Left = &TreeNode{Val: 1}
	a.Right.Left.Right = &TreeNode{Val: 1}
	a.Right.Left.Right.Right = &TreeNode{Val: 1}
	fmt.Println(longestZigZag(a)) // expected 3

	// ---- sample 2: explicit zigzag chain of 5 nodes -> edges = 4
	b := &TreeNode{Val: 1}
	b.Right = &TreeNode{Val: 2}
	b.Right.Left = &TreeNode{Val: 3}
	b.Right.Left.Right = &TreeNode{Val: 4}
	b.Right.Left.Right.Left = &TreeNode{Val: 5}
	fmt.Println(longestZigZag(b)) // expected 4

	// ---- sample 3: single node
	s := &TreeNode{Val: 1}
	fmt.Println(longestZigZag(s)) // expected 0
}
