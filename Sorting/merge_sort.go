// Xắp xếp trộn
// leetcode 912
// Ý tưởng: chia ra và merge vào
package main

import "fmt"

func MyMergeSort(arr []int, L int, R int) []int {
	// THBD DKD
	if L > R {
		return []int{}
	}

	if L == R {
		singleElement := [1]int{arr[L]}
		return singleElement[:]
	}
	// Chia ra
	k := (L + R) / 2

	arr1 := MyMergeSort(arr, L, k)
	arr2 := MyMergeSort(arr, k+1, R)

	// Tron vao:
	n := len(arr1) + len(arr2)
	result := make([]int, n)
	i := 0
	i1 := 0
	i2 := 0

	for i < n {
		if i1 < len(arr1) && i2 < len(arr2) { // trường hợp: cả arr1 và arr2 đều khác rỗng
			if arr1[i1] <= arr2[i2] {
				result[i] = arr1[i1]
				i++
				i1++
			} else {
				result[i] = arr2[i2]
				i++
				i2++
			}
		} else { // trường hợp: arr1 hoặc arr2 rỗng
			if i1 < len(arr1) { // arr1 khác rỗng
				result[i] = arr1[i1]
				i++
				i1++
			} else { // arr2 khác rỗng
				result[i] = arr2[i2]
				i++
				i2++
			}
		}
	}

	return result
}

func sortArray(arr []int) []int {

	return MyMergeSort(arr, 0, len(arr)-1)
}

func main() {

	arr := []int{1, 5, 3, 2, 8, 7, 6, 4}

	fmt.Println("Mảng trước khi sort: ")
	for _, ai := range arr {
		fmt.Print(ai)
	}

	result := sortArray(arr)

	fmt.Println("Mảng sau khi sort: ")
	for _, ai := range result {
		fmt.Print(ai)
	}

}
