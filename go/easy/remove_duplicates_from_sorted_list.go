package easy

import (
	"fmt"
	"xavierruiz/utils"
)

func RemoveDuplicatesFromSortedList() {
	// init linked list
	numbers := []int{1, 1, 7, 8, 2, 3, 3}
	fmt.Println("Hello World")
	head := utils.CreateLinkedList(numbers)
	// BEGIN PROBLEM
	// create a map to store values as we traverse the list
	// if value already exists in map, remove the node by setting previous node.Next to current.Next
	// no need to update anything else as current is untouched

	head = deleteDuplicates(head)
	fmt.Println("------------OUTPUT-------------")

	for head != nil {
		fmt.Printf("Output value: %d \n", head.Val)
		head = head.Next
	}
}

func deleteDuplicates(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}
	fmt.Println("------START------")
	previous := head
	current := head.Next
	valueMap := make(map[int]int)
	valueMap[head.Val] = 1
	for current != nil {
		fmt.Println("------NODE VALUES-------")

		fmt.Printf("[INFO] Previous: %d \n", previous.Val)
		fmt.Printf("[INFO] Current: %d \n", current.Val)
		// always increment counter when encountering a value
		valueMap[current.Val] += 1

		// deletion logic
		if valueMap[current.Val] > 1 {
			fmt.Printf("[INFO] Deleting node value: %d \n", current.Val)
			previous.Next = current.Next
			// don't increment previous only current
		} else {
			// if no duplicate found, increment previous and current
			previous = current
		}
		current = current.Next

	}
	return head
}
