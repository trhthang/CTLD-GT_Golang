package main

import "fmt"

func Fibonacci(n int) int {

	// điều kiện dừng
	if n <= 2 {
		return 1
	}
	// Công thức quy nạp
	return Fibonacci(n-2) + Fibonacci(n-1)

}

func main() {
	fmt.Println(Fibonacci(9))
}
