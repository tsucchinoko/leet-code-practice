package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	if head == nil || head.Next == nil {
		return 0
	}
	// 1. 中間点を見つける（slow が前半終了、slow.Next が後半先頭）
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// リスト長は偶数なので、slow は n/2 番目（0-indexed だと n/2）
	// 後半の先頭は slow
	lastHalfHead := reverseList(slow)

	// 2. head と lastHalfHead を同時走査して最大ツイン和を求める
	p1, p2 := head, lastHalfHead
	maxSum := 0
	for p1 != nil && p2 != nil {
		s := p1.Val + p2.Val
		if s > maxSum {
			maxSum = s
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 後半を元に戻す（テスト用）
	reverseList(lastHalfHead)

	return maxSum
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// ヘルパー: スライスから連結リストを作る（テスト用）
func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	cur := head
	for i := 1; i < len(vals); i++ {
		cur.Next = &ListNode{Val: vals[i]}
		cur = cur.Next
	}
	return head
}

func main() {
	// 簡単なデモ
	head := buildList([]int{5, 4, 2, 1})
	fmt.Println(pairSum(head)) // 6
}
