package main

import "fmt"

var curMax int

func FindMax(arr []int, i int) {
	if i < len(arr) {
		if arr[i] > curMax {
			curMax = arr[i]
		}
		FindMax(arr, i+1)
	}

}

func main() {
	arr := []int{1, 2, 3, 4, 6, 7, 9, 0}
	FindMax(arr, 0)
	fmt.Println(curMax)
}
