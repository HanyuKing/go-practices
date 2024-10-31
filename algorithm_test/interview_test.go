package algorithm

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	a := 2
	b := 3

	// a, b = b, a

	a = a + b
	b = a - b
	a = a - b

	fmt.Println(a)
	fmt.Println(b)
}

/**

李心宇对所有人说 (10:53)
type Node struct {
	Left  *Node
	Right *Node
	Val   int
}

// binary search tree to double linked list
// for double linked list, use "Left" as prev pointer
// for double linked list, use "Right" as next pointer
// param: root of a binary search tree
// ret: head of a sorted double linked list

      2
    1    3

1<->2<->3

*/

// 1 2 4 5
// 3->2
// 6->5
// 0->-1

func TestFindTarget(t *testing.T) {
	array := []int{1, 2, 4, 5}
	fmt.Println(findTarget(array, 3))
	fmt.Println(findTarget(array, 6))
	fmt.Println(findTarget(array, 0))
	fmt.Println(findTarget(array, 2))
}

// epoll select poll
// tcp ip 高效编程
func findTarget(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low - 1
}

//
//func BST2DoubleLinkedList(root *Node) *Node {
//	if root == nil {
//		return nil
//	}
//
//	left := BST2DoubleLinkedList(root.Left)
//
//	root.Left = left
//	//fmt.Println(root.Val)
//
//	right := BST2DoubleLinkedList(root.Right)
//	right
//
//	return nil
//}
//
//type Node struct {
//	Left  *Node
//	Right *Node
//	Val   int
//}
