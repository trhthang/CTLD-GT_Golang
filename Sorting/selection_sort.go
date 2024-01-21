/*
Nguyên lý:
	- chạy từ đầu đến cuối mảng:
	- Tại vòng lặp thứ i, tìm phần tử nhỏ nhất từ [i+1, n-1], nếu nhỏ hơn a[i]
	thì đổi chỗ cho a[i]
	- Sau vòng lặp thứ i, thì dãy [0,i] đã được sắp xếp
*/

package main

import (
	"fmt"
)

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// tìm phần tử nhỏ nhất trong khoảng từ [i+1, n-1]
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		// sau khi tìm được index của phần tử nhỏ nhất --> hoán đổi vị trí:
		arr[i], arr[minIndex] = arr[minIndex], arr[i]

	}
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Mảng trước khi sắp xếp: ", arr)

	selectionSort(arr)
	fmt.Println("Mảng sau khi sắp xếp: ", arr)
}
