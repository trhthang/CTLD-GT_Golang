package main

import (
	"fmt"
	"strconv"
)

func convertToBase7(num int) string {

	// Nếu num < 0 --> thêm "-" + giá trị tuyệt đối của num
	if num < 0 {
		return "-" + convertToBase7(-num)
	}

	// Nếu num
	if num < 7 && num >= 0 {
		return strconv.Itoa(num)
	}
	return convertToBase7(num/7) + strconv.Itoa(num%7)
}

func main() {
	fmt.Println(convertToBase7(100))
}
