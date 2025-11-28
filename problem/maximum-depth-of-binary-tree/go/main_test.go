package main

import "testing"

func buildTree(vals ...interface{}) *TreeNode {
	// ヘルパー: レベル順の配列（nil を表す nil インタフェース）からツリーを構築
	// ここでは簡易実装：配列が空なら nil、インデックス i の左右は 2*i+1, 2*i+2
	n := len(vals)
	if n == 0 {
		return nil
	}
	nodes := make([]*TreeNode, n)
	for i, v := range vals {
		if v == nil {
			nodes[i] = nil
		} else {
			nodes[i] = &TreeNode{Val: v.(int)}
		}
	}
	for i := 0; i < n; i++ {
		if nodes[i] == nil {
			continue
		}
		leftIdx := 2*i + 1
		rightIdx := 2*i + 2
		if leftIdx < n {
			nodes[i].Left = nodes[leftIdx]
		}
		if rightIdx < n {
			nodes[i].Right = nodes[rightIdx]
		}
	}
	return nodes[0]
}

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{"empty", nil, 0},
		{"single", &TreeNode{Val: 1}, 1},
		{"example1", buildTree(3, 9, 20, nil, nil, 15, 7), 3},
		{"example2", buildTree(1, nil, 2), 2},
		{"left-heavy", &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, 3},
		{"right-heavy", &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.root); got != tt.want {
				t.Fatalf("maxDepth() = %d, want %d", got, tt.want)
			}
		})
	}
}
