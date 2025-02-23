package easy

import "fmt"

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	fmt.Println("MERGE SORTED ARRAY")
	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	fmt.Printf("Values - nums1: %v - m: %d - nums2: %v - n: %d \n", nums1, m, nums2, n)
}
