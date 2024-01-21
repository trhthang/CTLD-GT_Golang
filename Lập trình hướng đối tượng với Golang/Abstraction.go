/*
Trong Golang, tính trừu tượng không được thực hiện thông qua các lớp trừu tượng như trong các ngôn ngữ OOP cổ điển khác,
mà thông qua việc sử dụng interfaces. Một interface định nghĩa một tập hợp các phương thức nhưng không cung cấp triển khai
cụ thể cho bất kỳ phương thức nào. Các kiểu dữ liệu khác nhau có thể triển khai các interfaces này, cung cấp triển khai cụ thể
cho các phương thức của chúng.*/

package main

import (
	"fmt"
	"math"
)

// Shape là 1 interface định nghĩa hành vi trừu tượng của các hình dạng hình học
type Shape interface {
	Area() float64
}

// Hinh tron triển khai interface Shape
type Circle struct {
	Radius float64
}

// Area cung cấp một triển khai cụ thể cho Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Hình vuông triển khai interface Shape
type Square struct {
	Side float64
}

// Area cung cấp 1 triển khai cụ thể cho Square
func (s Square) Area() float64 {
	return s.Side * s.Side
}

// PrintArea nhận vào một Shape và in ra diện tích của nó
// Đây là một hàm trừu tượng vì nó có thể làm việc với bất kỳ
// kiểu Shape nào
func PrintArea(shape Shape) {
	fmt.Println("Area of the shape is: ", shape.Area())
}

func main() {
	circle := Circle{Radius: 3}
	square := Square{Side: 4}

	// Dùng hàm trừu tượng PrintArea để in ra diện tích của cả hình tròn và hình vuông
	PrintArea(circle)
	PrintArea(square)
}
