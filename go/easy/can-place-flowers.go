package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	// determine number of valid places
	for idx, val := range flowerbed {
		// if place is a 0 surrounded by 0s or an edge 0 with only 1 0 to left or right
		if len(flowerbed) == 1 {
			if val == 0 && n == 1 {
				return true
			}
		}
		if val == 0 && len(flowerbed) > 1 {
			// first val checks
			if idx == 0 && idx-1 == -1 {
				if flowerbed[idx+1] == 0 {
					n -= 1
					flowerbed[idx] = 1
					fmt.Printf("FIRST VALID: n - %v, flowerbed[idx] - %v \n", n, flowerbed[idx])
				}
			}

			// last val checks
			if idx == len(flowerbed)-1 {
				if flowerbed[idx-1] == 0 {
					n -= 1
					flowerbed[idx] = 1
					fmt.Printf("LAST VALID: n - %v, flowerbed[idx] - %v \n", n, flowerbed[idx])
				}
			}

			// middle val checks
			if idx-1 != -1 && idx != len(flowerbed)-1 {
				if flowerbed[idx-1] == 0 && flowerbed[idx+1] == 0 {
					n -= 1
					flowerbed[idx] = 1
					fmt.Printf("MIDDLE VALID: n - %v, flowerbed[idx] - %v \n", n, flowerbed[idx])

				}
			}
		}
	}
	return n <= 0
}

func main() {
	flowerbed := []int{0, 1, 0}
	fmt.Println(canPlaceFlowers(flowerbed, 1))
}

// 0, 0, 1, 1, 1 1
// 1, 0, 1, 0, 1 0
// 1, 1, 1, 0, 0 1
// 0, 0, 0, 0, 0 3
