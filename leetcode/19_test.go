package leetcode

import (
	"fmt"
	"testing"
)

func TestAAA(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	fmt.Println("before")
	headTemp := head
	for headTemp != nil {
		fmt.Print(fmt.Sprintf("%d ", headTemp.Val))
		headTemp = headTemp.Next
	}
	fmt.Println()

	head = removeNthFromEnd(head, 2)

	fmt.Println("after")
	headTemp = head
	for headTemp != nil {
		fmt.Print(fmt.Sprintf("%d ", headTemp.Val))
		headTemp = headTemp.Next
	}
	fmt.Println()
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	faster := head
	slower := head
	for n > 0 {
		faster = faster.Next
		if faster == nil {
			if n == 1 {
				return head.Next
			} else {
				return head
			}
		}
		n--
	}
	for faster.Next != nil {
		slower = slower.Next
		faster = faster.Next
	}
	target := slower.Next
	slower.Next = target.Next
	target.Next = nil
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
