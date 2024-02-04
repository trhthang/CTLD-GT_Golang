// Khởi tạo một mảng chứa các số từ 1 đến 20.
// Sử dụng slicing và loops để tạo evenNumbers và oddNumbers.
// Đảo ngược evenNumbers (bạn có thể viết hàm hoặc sử dụng loops).
// Tạo firstFiveOdds từ oddNumbers.
// In các slice này ra màn hình để kiểm tra kết quả

package main

import "fmt"

func main() {
	// Bước 1: Khởi tạo mảng chứa số từ 1 đến 20
	var numbers [20]int
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i + 1
	}

	// Bước 2: Tạo Slice cho số chẵn
	var evenNumbers, oddNumbers []int

	for _, ai := range numbers {

		if ai%2 == 0 {
			evenNumbers = append(evenNumbers, ai)
		}

		if ai%2 != 0 {
			oddNumbers = append(oddNumbers, ai)
		}

	}

	fmt.Println("Even Numbers:", evenNumbers)

	fmt.Println()

	// Bước 3: Đảo ngược Slice evenNumbers
	var i = 0
	var j = len(evenNumbers) - 1

	daoNguoc(evenNumbers, i, j)

	fmt.Println("Even Numbers:", evenNumbers)

	// Bước 4: Tạo slide firstFiveOdds từ oddNumbers
	firstFiveOdds := oddNumbers[2:5]

	// 	for _, ai := range numbers {
	// 		println(ai)
	// 	}

	fmt.Println("Odd Numbers:", oddNumbers)
	fmt.Println("First Five Odd Numbers:", firstFiveOdds)
}

func daoNguoc(arr []int, i int, j int) {
	for i >= j {
		return
	}

	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp

	daoNguoc(arr, i+1, j-1)
}
