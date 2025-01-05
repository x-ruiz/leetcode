package main

func productExceptSelf(nums []int) []int {
	product := 1
	var output []int
	zeroCounter := 0

	for _, value := range nums {
		if value == 0 && zeroCounter == 0 {
			zeroCounter += 1
			continue
		}
		product = product * value
	}

	for _, value := range nums {
		if value != 0 {
			newProduct := product
			if zeroCounter > 0 {
				newProduct = product * 0
			}
			output = append(output, newProduct/value)
		} else {
			// replace 0 with a 1 to calculate product / value for that slot only
			output = append(output, product)
		}

	}

	return output
}

// func main() {
// 	nums := []int{1, 0, 0, 4}
// 	fmt.Println(productExceptSelf(nums))
// }
