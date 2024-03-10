package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	mySet := make(map[int]bool)

	for _, num := range nums {
		mySet[num] = true
	}

	var result []int
	for i := 1; i <= len(nums); i++ {
		if mySet[i] != true {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	result := findDisappearedNumbers(nums)
	fmt.Println(result)
}
