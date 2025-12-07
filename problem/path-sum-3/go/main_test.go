package main

import "testing"

func buildTreeFromSlice(vals []interface{}) *TreeNode {
	// 局所的にレベル順（nil を表す場合は nil を入れる）からツリーを構築する簡易関数
	if len(vals) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		if v == nil {
			nodes[i] = nil
		} else {
			nodes[i] = &TreeNode{Val: v.(int)}
		}
	}
	for i := 0; i < len(vals); i++ {
		if nodes[i] == nil {
			continue
		}
		leftIdx := 2*i + 1
		rightIdx := 2*i + 2
		if leftIdx < len(vals) {
			nodes[i].Left = nodes[leftIdx]
		}
		if rightIdx < len(vals) {
			nodes[i].Right = nodes[rightIdx]
		}
	}
	return nodes[0]
}

func TestPathSumExamples(t *testing.T) {
	// Example 1
	root1 := buildTreeFromSlice([]interface{}{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1})
	got1 := pathSum(root1, 8)
	if got1 != 3 {
		t.Fatalf("example1: got %d, want 3", got1)
	}

	// Example 2
	root2 := buildTreeFromSlice([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1})
	got2 := pathSum(root2, 22)
	if got2 != 3 {
		t.Fatalf("example2: got %d, want 3", got2)
	}
}

func TestEdgeCases(t *testing.T) {
	// 空の木
	if pathSum(nil, 0) != 0 {
		t.Fatalf("empty tree should return 0")
	}
	// 単一ノードがターゲットに等しい
	single := &TreeNode{Val: 5}
	if pathSum(single, 5) != 1 {
		t.Fatalf("single node equal target should be 1")
	}
	// 負の値を含むケース
	root := buildTreeFromSlice([]interface{}{1, -2, -3, 1, 3, -2, nil, -1})
	if pathSum(root, -1) != 4 {
		t.Fatalf("negative values case: got %d", pathSum(root, -1))
	}
}
