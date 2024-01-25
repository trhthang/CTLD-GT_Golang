package main

import (
	"fmt"
)

func firstUniqChar(s string) int {

	// tạo 1 map với : key là kiểu "rune" (tương đương với char trong java) , value là kiểu int -> để đếm count của mỗi kí tự

	charCount := make(map[rune]int)

	//charCount[ch]++ là lệnh tăng giá trị của key tương ứng trong map lên 1.
	//  Nếu ch chưa có trong map, Go sẽ tự động khởi tạo giá trị là 0 và sau đó tăng lên 1.
	// Nếu ch đã có trong map, thì giá trị hiện tại của nó sẽ được tăng lên.
	for _, ch := range s {
		charCount[ch]++
	}

	//duyệt mảng string từ đầu đến cuối
	for i, ch := range s {
		if charCount[ch] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	s := "aabb"
	result := firstUniqChar(s)
	fmt.Println(result)
}
