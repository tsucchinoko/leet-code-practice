package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

	// p and q find in a different tree
	if left != nil && right == nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

// helper to build tree from slice (nil represented by sentinel with ok flag)
func buildTree(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	idx := 1
	for idx < len(vals) && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if idx < len(vals) && vals[idx] != nil {
			node.Left = &TreeNode{Val: vals[idx].(int)}
			queue = append(queue, node.Left)
		}
		idx++
		if idx < len(vals) && vals[idx] != nil {
			node.Right = &TreeNode{Val: vals[idx].(int)}
			queue = append(queue, node.Right)
		}
		idx++
	}
	return root
}

func main() {
	// Example usage
	root := buildTree([]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4})
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
