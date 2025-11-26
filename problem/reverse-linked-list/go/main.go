package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func reverseRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// helper: sliceToList はスライスから連結リストを作成する。
func sliceToList(nums []int) *ListNode {
	var head, tail *ListNode
	for _, v := range nums {
		node := &ListNode{Val: v}
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}
	return head
}

// helper: listToSlice は連結リストをスライスに変換する。
func listToSlice(head *ListNode) []int {
	res := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		res = append(res, cur.Val)
	}
	return res
}

func main() {
	examples := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2},
		{},
	}
	for _, ex := range examples {
		l := sliceToList(ex)
		r := reverseList(l)
		fmt.Printf("iterative %v -> %v\n", ex, listToSlice(r))

		l2 := sliceToList(ex)
		r2 := reverseRecursive(l2)
		fmt.Printf("recursive %v -> %v\n", ex, listToSlice(r2))
	}
}
