package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1, root2 *TreeNode) bool {
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

func leafSimilarbyStack(root1, root2 *TreeNode) bool {
	a := leavesStack(root1)
	b := leavesStack(root2)
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

func leaves(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			res = append(res, node.Val)
			return
		}
		// process left child first
		dfs(node.Left)
		dfs(node.Right)

	}
	dfs(root)
	return res
}

func leavesStack(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node == nil {
			continue
		}
		if node.Left == nil && node.Right == nil {
			res = append(res, node.Val)
			continue
		}
		// LIFOのため、先に右側をスタックに積む
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
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

	fmt.Println(leafSimilar(r1, r2))        // false
	fmt.Println(leafSimilarbyStack(r1, r2)) // false
}
