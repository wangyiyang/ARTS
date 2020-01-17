package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	var visited []*ListNode
	node := head
	for node != nil {
		for i := 0; i < len(visited); i++ {
			if node == visited[i] {
				return node
			}
		}
		visited = append(visited, node)
		node = node.Next
	}
	return nil
}

func main() {
	fmt.Println("vim-go")
}
