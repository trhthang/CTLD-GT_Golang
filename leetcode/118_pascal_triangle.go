package main

import "fmt"

func generate(numRows int) [][]int {
	// create 2D array
	array := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		array[i] = make([]int, i+1) // Khởi tạo slice con với độ dài là i+1
	}

	// calculate:
	array[0] = make([]int, 1)
	array[0][0] = 1
	for i := 1; i < numRows; i++ {
		array[i] = make([]int, i+1)
		array[i][0] = 1
		array[i][i] = 1
		for j := 1; j < i; j++ {
			array[i][j] = array[i-1][j] + array[i-1][j-1]
		}
	}
	return array
}

func main() {
	result := generate(5)
	fmt.Println(result)
}
