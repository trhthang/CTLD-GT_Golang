package main

import "fmt"

// type NumArray struct {
// 	myArray []int
// }

// func Constructor(nums []int) NumArray {
// 	myArray := make([]int, len(nums))
// 	for i, num := range nums {
// 		myArray[i] = num
// 	}
// 	return NumArray{myArray}
// }

// func (this *NumArray) SumRange(left int, right int) int {
// 	var result int = 0
// 	for i := left; i <= right; i++ {
// 		result = result + this.myArray[i]
// 	}

// 	return result
// }
type NumArray struct {
	cumSum []int
}

func Constructor(nums []int) NumArray {
	cumSum := make([]int, len(nums)+1)
	fmt.Println(cumSum)
	for i, num := range nums {
		cumSum[i+1] = cumSum[i] + num
	}
	return NumArray{cumSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.cumSum[right+1] - this.cumSum[left]
}

func main() {
	// Initialize NumArray with a slice of integers.
	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})

	// Perform sum range queries.
	// fmt.Println(numArray.SumRange(0, 2)) // Outputs 1
	// fmt.Println(numArray.SumRange(2, 5)) // Outputs -1
	// fmt.Println(numArray.SumRange(0, 5)) // Outputs -3

	fmt.Println(numArray)
}
