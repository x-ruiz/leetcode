package easy

import (
	"sort"
)

// Doubting that this modifies original nums array, might have to not use beginning or end and combined slice
func removeElement(nums []int, val int) int {
	count := len(nums)
	for idx := 0; idx < len(nums); idx++ {
		if val == nums[idx] {
			// slice up to the idx
			nums[idx] = -1
			count -= 1
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return count
}

// func main() {
// 	nums := []int{3, 2, 2, 3}
//
// 	fmt.Println(removeElement(nums, 3))
// }
