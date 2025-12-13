package main

import "testing"

func levelOrderVals(root *TreeNode) []any {
	if root == nil {
		return nil
	}
	res := []any{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if n != nil {
			res = append(res, n.Val)
			q = append(q, n.Left, n.Right)
		} else {
			res = append(res, nil)
		}
	}
	// trim trailing nils
	i := len(res) - 1
	for i >= 0 && res[i] == nil {
		i--
	}
	return res[:i+1]
}

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
	root := buildTree([]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4})
	p := findByVal(root, 5)
	q := findByVal(root, 1)
	t.Logf("tree: %v", levelOrderVals(root))
	t.Logf("p: val=%d ptr=%p, q: val=%d ptr=%p", p.Val, p, q.Val, q)
	got := lowestCommonAncestor(root, p, q)
	if got == nil || got.Val != 5 {
		t.Fatalf("expected LCA val 5, got %v", got)
	}
}

func TestLowestCommonAncestor_Example2(t *testing.T) {
	root := buildTree([]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4})
	p := findByVal(root, 5)
	q := findByVal(root, 4)
	t.Logf("tree: %v", levelOrderVals(root))
	t.Logf("p: val=%d ptr=%p, q: val=%d ptr=%p", p.Val, p, q.Val, q)
	got := lowestCommonAncestor(root, p, q)
	if got == nil || got.Val != 3 {
		t.Fatalf("expected LCA val 3, got %v", got)
	}
}

func TestLowestCommonAncestor_Example3(t *testing.T) {
	root := buildTree([]any{1, 2})
	p := findByVal(root, 1)
	q := findByVal(root, 2)
	t.Logf("tree: %v", levelOrderVals(root))
	t.Logf("p: val=%d ptr=%p, q: val=%d ptr=%p", p.Val, p, q.Val, q)
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
	t.Logf("tree: %v", levelOrderVals(root))
	t.Logf("p: val=%d ptr=%p, q: val=%d ptr=%p", p.Val, p, q.Val, q)
	got := lowestCommonAncestor(root, p, q)
	if got == nil || got.Val != 1 {
		t.Fatalf("expected LCA val 1, got %v", got)
	}
}
