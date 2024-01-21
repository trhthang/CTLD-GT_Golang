package main

import (
	"fmt"
)

func main() {

	var number int

	number = 10

	// Type Inference

	var email = "trhthang0401@gmail.com"

	// Tạo nhiều biến cùng lúc  cùng kiểu dữ liệu:

	var a1, a2 int
	a1 = 1
	a2 = 2

	fmt.Println(a1)
	fmt.Println(a2)

	// Tạo nhiều biến nhưng khác kiểu dữ liệu:
	var (
		name    string
		address string
		age     int
	)

	name = "Thang"
	address = "Viet Nam"
	age = 24

	fmt.Println(name)
	fmt.Println(address)
	fmt.Println(age)

	// hoặc :
	var (
		name1    string = "thang1"
		address1 string = "vietnam"
		age1     int    = 25
	)
	fmt.Println(name1)
	fmt.Println(address1)
	fmt.Println(age1)

	fmt.Println(number)
	fmt.Println(email)

	// short hand declaration

	language := "golang"

	fmt.Println(language)
}
