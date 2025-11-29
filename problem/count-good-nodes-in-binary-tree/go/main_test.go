package main

import "testing"

func printLevelOrder(root *TreeNode) []any {
	if root == nil {
		return nil
	}
	res := []any{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if n == nil {
			res = append(res, nil)
			continue
		}
		res = append(res, n.Val)
		q = append(q, n.Left, n.Right)
	}
	// 末尾の nil をトリムすると見やすくなります
	i := len(res) - 1
	for i >= 0 && res[i] == nil {
		i--
	}
	return res[:i+1]
}

// buildTreeFromSlice builds a binary tree from a level-order slice where nil indicates absent nodes.
// This uses a queue to properly assign children in level order (compatible with LeetCode-style arrays).
func buildTreeFromSlice(vals []any) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	if vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	idx := 1
	for idx < len(vals) && len(queue) > 0 {
		parent := queue[0]
		queue = queue[1:]

		// left child
		if idx < len(vals) {
			if vals[idx] != nil {
				left := &TreeNode{Val: vals[idx].(int)}
				parent.Left = left
				queue = append(queue, left)
			}
			idx++
		}

		// right child
		if idx < len(vals) {
			if vals[idx] != nil {
				right := &TreeNode{Val: vals[idx].(int)}
				parent.Right = right
				queue = append(queue, right)
			}
			idx++
		}
	}
	return root
}

func TestCountGoodNodes(t *testing.T) {
	tests := []struct {
		vals []any
		want int
	}{
		{[]any{3, 1, 4, 3, nil, 1, 5}, 4},
		{[]any{3, 3, nil, 4, 2}, 3},
		{[]any{1}, 1},
		{[]any{-1, -2, -3, nil, -2, nil, -1}, 2}, // マイナス値の例
	}

	for _, tc := range tests {
		root := buildTreeFromSlice(tc.vals)
		t.Logf("built tree: %v", printLevelOrder(root))
		got := goodNodes(root)
		if got != tc.want {
			t.Fatalf("for %v: want %d, got %d", tc.vals, tc.want, got)
		}
	}
}
