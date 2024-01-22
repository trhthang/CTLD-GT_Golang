package main

import "fmt"

// Tính n!, return n!
func GiaiThua(n int) int {

	// bài toán cơ sở
	if n == 0 {
		return 1
	}
	// Công thức quy nạp
	return n * GiaiThua(n-1)
}

func main() {
	fmt.Println(GiaiThua(7))
}
