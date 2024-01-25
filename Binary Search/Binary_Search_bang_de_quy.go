package main

import "fmt"

func BinarySearch(arr []int, key int, iLeft int, iRight int) int {
	if iLeft > iRight {
		return -1
	}

	iMid := (iLeft + iRight) / 2
	if arr[iMid] == key {
		return iMid
	} else if arr[iMid] > key {
		return BinarySearch(arr, key, iLeft, iMid-1)
	} else {
		return BinarySearch(arr, key, iMid+1, iRight)
	}
}

func main() {
	arr := []int{1, 3, 9, 12, 19, 15, 31, 47, 50, 57, 72}
	fmt.Println(BinarySearch(arr, 31, 0, len(arr)))
}
