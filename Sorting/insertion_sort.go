/*
Nguyên lý:
	- chạy từ đầu đến cuối mảng
	- Tại vòng lặp i, coi như dãy từ [0. i-1] đã được sắp xếp, chèn phần tử a[i] vào
	vị trí thích hợp
	- Sau vòng lặp thứ i thì dãy [0,i] đã được sắp xếp
*/

package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// chèn a[i] vào dãy 0 -> i - 1

		ai := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > ai {
			arr[j+1] = arr[j]
			j = j - 1
		}

		arr[j+1] = ai

	}

}

func main() {
	arr := []int{12, 11, 13, 5, 6}
	fmt.Println("Mảng trước khi sắp xếp:", arr)

	insertionSort(arr)
	fmt.Println("Mảng sau khi sắp xếp:", arr)
}
