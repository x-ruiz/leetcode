package main

import (
	"fmt"
)

type Stack struct {
	items []rune
}

func (s *Stack) Push(data rune) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	if len(s.items) == 0 {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Top() rune {
	if len(s.items) == 0 {
		return -1
	}
	return s.items[len(s.items)-1]
}

func isValid(s string) bool {
	parenMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack Stack
	for _, c := range s {
		value, ok := parenMap[c]
		fmt.Println(stack)
		if ok {
			// if ok then character is a closer, pop last added opener
			fmt.Printf("STACK CHECKING Value: %v | Stack Top %v | Stack %v \n", value, stack.Top(), stack.items)
			if value == stack.Top() {
				fmt.Printf("Encountered character %v, popping last character of stack %v \n", string(c), stack)

				stack.Pop()
			} else {
				return false // encountered closer without corresponding opener, indicates out of order
			}
		} else {
			// character is an opener
			fmt.Printf("Pushing character %v", string(c))
			stack.Push(c)
		}
	}
	fmt.Printf("STACK: %v \n", stack.items)
	if len(stack.items) == 0 {
		return true
	} else {
		return false
	}
}

// func main() {
// 	fmt.Println(isValid("(]"))
// }
