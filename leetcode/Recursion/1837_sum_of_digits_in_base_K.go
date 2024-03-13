package main

import "fmt"

// logic sẽ là : lấy phần dư phép tính hiện tại cộng với phần dư phép tính tiếp theo

func sumBase(n int, k int) int {
	// Trường  hợp cơ sở: nếu n bằng 0, trả về tổng là 0 (điều kiện kết thúc đệ quy)
	if n == 0 {
		return 0
	}

	// Trả về phần dư của n chia cho k (chữ số cuối cùng trong hệ cơ số k)
	// cộng với phần dư của phép tính chia tiếp theo
	return n%k + sumBase(n/k, k)
}

func main() {
	// Kiểm tra hàm sumBase với một số ví dụ

	fmt.Println(sumBase(34, 6))
}
