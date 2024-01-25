package main

// cách này là tối ưu đệ quy bằng cách lưu các kết quả đã tính vào trong 1 mảng tên là F ...

import "fmt"

var F [1000]int

func Fibonacci(n int) int {
	if F[n] != 0 {
		return F[n]
	}

	if n <= 2 {
		F[1] = 1
		F[2] = 1
		return 1
	}

	F[n] = Fibonacci(n-2) + Fibonacci(n-1)
	return F[n]
}

func main() {
	fmt.Println(Fibonacci(10))
}
