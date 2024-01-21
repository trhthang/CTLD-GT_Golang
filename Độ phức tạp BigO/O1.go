/*
Độ phức tạp O(1), còn được gọi là độ phức tạp thời gian hằng số,
là khi thời gian chạy của một thuật toán không phụ thuộc vào kích thước của đầu vào.
Một ví dụ đơn giản về độ phức tạp O(1) có thể là một hàm trả về phần tử thứ n của một mảng đã biết


Việc truy cập một phần tử trong mảng là một thao tác có độ phức tạp O(1) vì nó không phụ thuộc vào kích thước của mảng.
Thao tác này luôn mất một lượng thời gian cố định, bất kể mảng có bao nhiêu phần tử.
*/

package main

func getElement(arr []int, index int) int {
	// trả về 1 phần từ trong mảng là một thao tác với độ phức tạp O(1)
	return arr[index]
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	index := 2

	result := getElement(arr, index)
	println("Phần từ tại vị trí", index, "là: ", result)
}
