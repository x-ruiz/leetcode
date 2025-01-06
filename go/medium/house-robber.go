package main

import (
	"fmt"
	"math"
)

func rob(nums []int) int {
	// base case
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return int(math.Max(float64(nums[0]), float64(nums[1])))
	}
	fmt.Printf("NUMS: %v \n ", nums)
	firstVal := nums[0]
	yesSlice := nums[2:]
	noSlice := nums[1:]

	yes := rob(yesSlice) + firstVal
	no := rob(noSlice)

	if yes > no {
		return yes
	}
	return no
}

func main() {
	fmt.Println("HELL")
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
}
