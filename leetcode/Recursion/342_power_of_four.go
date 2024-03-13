package main

import "fmt"

func isPowerOfFour(n int) bool {
	// Điều kiện cơ bản: nếu n bằng 1, n là lũy thừa của 4
	if n == 1 {
		return true
	}
	// Điều kiện dừng: nếu n nhỏ hơn 0 hoặc n không chia hết cho 4
	if n < 1 || n%4 != 0 {
		return false
	}
	// Bước đệ quy: gọi lại hàm với n được chia cho 4
	return isPowerOfFour(n / 4)
}

func main() {
	fmt.Println(isPowerOfFour(16)) // true
	fmt.Println(isPowerOfFour(5))  // false
	fmt.Println(isPowerOfFour(1))  // true
}
