package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var dfs func(node *TreeNode, currentMax int) int
	dfs = func(node *TreeNode, currentMax int) int {
		if node == nil {
			return 0
		}
		count := 0
		if node.Val >= currentMax {
			count = 1
			currentMax = node.Val
		}
		count += dfs(node.Left, currentMax)
		count += dfs(node.Right, currentMax)
		return count
	}
	return dfs(root, root.Val)
}

func main() {
	// root = [3,1,4,3,null,1,5]
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 5}

	fmt.Println(goodNodes(root)) // expect 4
}
