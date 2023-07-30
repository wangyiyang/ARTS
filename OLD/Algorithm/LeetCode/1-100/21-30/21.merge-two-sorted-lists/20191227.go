package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
}

func main() {
	l1 := ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				4,

				nil,
			},
		},
	}
	l2 := ListNode{
		1,
		&ListNode{
			3,
			&ListNode{
				4,
				nil,
			},
		},
	}
	result := mergeTwoLists(&l1, &l2)
	for {
		fmt.Println(result.Val)
		if result.Next != nil {
			result = result.Next

		} else {
			break
		}
	}
}
