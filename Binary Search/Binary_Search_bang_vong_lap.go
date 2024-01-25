package main

import "fmt"

func BinarySearch(arr []int, key int) int {
	n := len(arr)
	iLeft := 0
	iRight := n - 1

	for iLeft <= iRight {
		iMid := (iLeft + iRight) / 2

		if arr[iMid] == key {
			return iMid
		} else if arr[iMid] > key {
			iRight = iMid - 1
		} else {
			iLeft = iMid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 9, 12, 19, 15, 31, 47, 50, 57, 72}
	fmt.Println(BinarySearch(arr, 31))
}
