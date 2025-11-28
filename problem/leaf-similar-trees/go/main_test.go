package main

import "testing"

func TestLeafSimilar_Example1(t *testing.T) {
	// root1 = [3,5,1,6,2,9,8,null,null,7,4]
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 5}
	root1.Right = &TreeNode{Val: 1}
	root1.Left.Left = &TreeNode{Val: 6}
	root1.Left.Right = &TreeNode{Val: 2}
	root1.Left.Right.Left = &TreeNode{Val: 7}
	root1.Left.Right.Right = &TreeNode{Val: 4}
	root1.Right.Left = &TreeNode{Val: 9}
	root1.Right.Right = &TreeNode{Val: 8}

	// root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
	root2 := &TreeNode{Val: 3}
	root2.Left = &TreeNode{Val: 5}
	root2.Right = &TreeNode{Val: 1}
	root2.Left.Left = &TreeNode{Val: 6}
	root2.Left.Right = &TreeNode{Val: 7}
	root2.Right.Left = &TreeNode{Val: 4}
	root2.Right.Right = &TreeNode{Val: 2}
	root2.Right.Right.Left = &TreeNode{Val: 9}
	root2.Right.Right.Right = &TreeNode{Val: 8}

	if !leafSimilar(root1, root2) {
		t.Fatalf("expected true for example1")
	}
}

func TestLeafSimilar_Example2(t *testing.T) {
	r1 := &TreeNode{Val: 1}
	r1.Left = &TreeNode{Val: 2}
	r1.Right = &TreeNode{Val: 3}

	r2 := &TreeNode{Val: 1}
	r2.Left = &TreeNode{Val: 3}
	r2.Right = &TreeNode{Val: 2}

	if leafSimilar(r1, r2) {
		t.Fatalf("expected false for example2")
	}
}

func TestLeafSimilar_SingleNode(t *testing.T) {
	a := &TreeNode{Val: 5}
	b := &TreeNode{Val: 5}
	if !leafSimilar(a, b) {
		t.Fatalf("single-node equal failed")
	}
	c := &TreeNode{Val: 6}
	if leafSimilar(a, c) {
		t.Fatalf("single-node different failed")
	}
}

func TestLeafSimilar_DifferentLengths(t *testing.T) {
	// left tree leaves: [1,2]
	a := &TreeNode{Val: 0}
	a.Left = &TreeNode{Val: 1}
	a.Right = &TreeNode{Val: 2}

	// right tree leaves: [1]
	b := &TreeNode{Val: 1}

	if leafSimilar(a, b) {
		t.Fatalf("different leaf counts should be false")
	}
}
