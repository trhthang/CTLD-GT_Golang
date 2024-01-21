/*
	Độ phức tạp O(n^2) ám chỉ rằng thời gian thực thi của thuật toán tăng theo bình phương của kích thước đầu vào.
*/

// Bubblesort sắp xếp một mảng sử dụng thuật toán sắp xếp nổi bọt
// Độ phức tạp cảu thuật toán này là O(N^2)

package main

import (
	"fmt"
)

// bubbleSort sắp xếp một mảng sử dụng thuật toán sắp xếp nổi bọt
// Độ phức tạp của thuật toán này là O(n^2)
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{5, 3, 8, 4, 2}
	bubbleSort(arr)
	fmt.Println("Mảng sau khi sắp xếp:", arr)
}
