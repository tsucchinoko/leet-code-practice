package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	prefix := make(map[int]int)
	// ルートからの経路そのものがtargetSumになる場合を考慮
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
