/*
Nguyên lý:
	- chạy từ đầu đến cuối mảng
	- Nếu phần tử đứng trước mà lớn hơn phần tử đứng sau -> thì đổi chỗ
	- Sau mỗi vòng lặp thì phần tử lớn nhất sẽ trôi xuống dưới
*/

package main

import "fmt"

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
	fmt.Println("Mảng sau khi sắp xếp: ", arr)
}
