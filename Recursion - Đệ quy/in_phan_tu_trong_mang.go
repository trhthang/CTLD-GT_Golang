package main

import "fmt"

func PrintElement(arr []int, index int) {
	if index < 0 || index >= len(arr) {
		return
	}

	// đệ quy thực hiện trước --> in ngược mảng
	PrintElement(arr, index+1)
	fmt.Println(arr[index])

	// đệ quy thực hiện sau --> in xuôi mảng
	// fmt.Println(arr[index])
	// PrintElement(arr, index+1)

}

func main() {
	myarr := [5]int{10, 20, 12, 48, 65}

	PrintElement(myarr[:], 0)

}
