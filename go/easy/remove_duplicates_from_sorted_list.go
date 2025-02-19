package easy

import (
	"fmt"
	"xavierruiz/utils"
)

func RemoveDuplicatesFromSortedList() {
	// init linked list
	numbers := []int{1, 1, 2, 3, 3}
	fmt.Println("Hello World")
	head := utils.CreateLinkedList(numbers)

	// BEGIN PROBLEM
	// create a map to store values as we traverse the list
	// if value already exists in map, remove the node by setting previous node.Next to current.Next
	// no need to update anything else as current is untouched
	head = deleteDuplicates(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func deleteDuplicates(head *utils.ListNode) *utils.ListNode {
	current := head
	previous := &utils.ListNode{Val: 0, Next: nil}
	valueMap := make(map[int]int)
	for current != nil {
		_, ok := valueMap[current.Val]
		if ok {
			// duplicate found remove nodes
			fmt.Printf("Removing node with value: %d \n", current.Val)
			previous.Next = current.Next
		} else {
			valueMap[current.Val] = 0
			previous = current
			current = current.Next

		}

	}
	return current
}
