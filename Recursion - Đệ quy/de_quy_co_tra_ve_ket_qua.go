package main

import (
	"fmt"
	"math"
)

func FindMax(arr []int, i int, previousMax int) int {

	if i < len(arr) {
		curMax := math.Max(float64(previousMax), float64(arr[i]))
		return FindMax(arr, i+1, int(curMax))
	}

	return previousMax
}

func main() {
	arr := []int{1, 2, 3, 4, 6, 7, 9, 0}
	fmt.Println(FindMax(arr, 0, -1))
}
