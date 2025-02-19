package utils

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkedList(numbers []int) *ListNode {
	fmt.Println(numbers)

	var head *ListNode
	var current *ListNode

	// create head
	head = &ListNode{Val: numbers[0], Next: nil}
	current = head
	// create linked list by iterating through input
	for _, num := range numbers {
		fmt.Printf("[INFO] Adding %d to linked list \n", num)
		newNode := &ListNode{Val: num, Next: nil}
		current.Next = newNode // update previous next to current node
		current = newNode      // update created node to previous for next iteration
	}
	return head
}
