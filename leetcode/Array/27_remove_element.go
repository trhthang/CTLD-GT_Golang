package main

import "fmt"

// [3,2,2,3]
func removeElement(nums []int, val int) int {
	k := 0

	for _, num := range nums {
		if num != val {
			nums[k] = num
			k++
		}
	}

	return k
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	result := removeElement(nums, 3)
	fmt.Println(result)
}
