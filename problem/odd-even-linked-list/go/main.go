package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next

		even.Next = even.Next.Next
		even = even.Next
	}

	odd.Next = evenHead
	return head
}

// helper: build list from slice
func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	cur := head
	for _, v := range vals[1:] {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return head
}

// helper: convert list to slice
func toSlice(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func main() {
	examples := [][]int{
		{1, 2, 3, 4, 5},
		{2, 1, 3, 5, 6, 4, 7},
		{},
		{1},
		{1, 2},
	}
	for _, ex := range examples {
		l := buildList(ex)
		out := oddEvenList(l)
		fmt.Println(toSlice(out))
	}
}
