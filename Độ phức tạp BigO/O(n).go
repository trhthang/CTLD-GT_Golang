/*
Độ phức tạp O(n) trong ngôn ngữ lập trình Go (hay Golang)
ám chỉ rằng thời gian thực thi của một thuật toán tăng tuyến tính theo kích thước của đầu vào.
 Một ví dụ phổ biến của độ phức tạp O(n) là một hàm tính tổng tất cả các phần tử trong một mảng.
*/

package main

import (
	"fmt"
)

// sumArray tính tổng tất cả các phần tử trong mảng
// Độ phức tạp của hàm này là O(n) vì nó duyệt qua từng phần tử của mảng
func sumArray(arr []int) int {
	total := 0
	for _, value := range arr {
		total += value
	}

	return total
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	sum := sumArray(arr)

	fmt.Println("Tổng các phần tử trong mảng là: ", sum)
}
