package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// deleteMiddle deletes the middle node (0-indexed floor(n/2)) and returns the head of the modified list.
func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var prev *ListNode
	slow, fast := head, head

	// Move fast by 2 and slow by 1, keeping prev one step behind slow.
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// slow is the middle node to delete, prev is its previous node (guaranteed non-nil because n>=2)
	prev.Next = slow.Next
	return head
}

func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for _, val := range vals[1:] {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}
	return head
}

func toSlice(head *ListNode) []int {
	res := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		res = append(res, cur.Val)
	}

	return res
}

func main() {
	cases := [][]int{
		{1, 3, 4, 7, 1, 2, 6},
		{1, 2, 3, 4},
		{2, 1},
		{1},
	}
	for _, c := range cases {
		head := buildList(c)
		newHead := deleteMiddle(head)
		fmt.Println(toSlice(newHead))
	}
}
