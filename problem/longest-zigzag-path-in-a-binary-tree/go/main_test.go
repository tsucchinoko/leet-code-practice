package main

import "testing"

func TestLongestZigZag(t *testing.T) {
	t.Run("sample1_manual", func(t *testing.T) {
		a := &TreeNode{Val: 1}
		a.Right = &TreeNode{Val: 1}
		a.Right.Left = &TreeNode{Val: 1}
		a.Right.Right = &TreeNode{Val: 1}
		a.Right.Left.Left = &TreeNode{Val: 1}
		a.Right.Left.Right = &TreeNode{Val: 1}
		a.Right.Left.Right.Right = &TreeNode{Val: 1}

		got := longestZigZag(a)
		want := 3
		if got != want {
			t.Fatalf("sample1: got %d want %d", got, want)
		}
	})

	t.Run("sample2_zigzag_chain", func(t *testing.T) {
		b := &TreeNode{Val: 1}
		b.Right = &TreeNode{Val: 2}
		b.Right.Left = &TreeNode{Val: 3}
		b.Right.Left.Right = &TreeNode{Val: 4}
		b.Right.Left.Right.Left = &TreeNode{Val: 5}

		got := longestZigZag(b)
		want := 4
		if got != want {
			t.Fatalf("sample2: got %d want %d", got, want)
		}
	})

	t.Run("single_node", func(t *testing.T) {
		s := &TreeNode{Val: 1}
		got := longestZigZag(s)
		want := 0
		if got != want {
			t.Fatalf("single node: got %d want %d", got, want)
		}
	})

	t.Run("left_chain_no_zigzag", func(t *testing.T) {
		n := &TreeNode{Val: 1}
		n.Left = &TreeNode{Val: 2}
		n.Left.Left = &TreeNode{Val: 3}
		n.Left.Left.Left = &TreeNode{Val: 4}
		got := longestZigZag(n)
		want := 1
		if got != want {
			t.Fatalf("left chain: got %d want %d", got, want)
		}
	})

	t.Run("right_chain_no_zigzag", func(t *testing.T) {
		n := &TreeNode{Val: 1}
		n.Right = &TreeNode{Val: 2}
		n.Right.Right = &TreeNode{Val: 3}
		got := longestZigZag(n)
		want := 1
		if got != want {
			t.Fatalf("right chain: got %d want %d", got, want)
		}
	})

	t.Run("small_zigzag_two_edges", func(t *testing.T) {
		p := &TreeNode{Val: 1}
		p.Right = &TreeNode{Val: 2}
		p.Right.Left = &TreeNode{Val: 3}
		got := longestZigZag(p)
		want := 2
		if got != want {
			t.Fatalf("small zigzag: got %d want %d", got, want)
		}
	})
}
