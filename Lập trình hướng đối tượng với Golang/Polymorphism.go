package main

import (
	"fmt"
	"math"
)

// Shape là một interface với phương thức Area và Perimeter
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle triển khai interface Shape
type Circle struct {
	Radius float64
}

// Area của Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter (chu vi) của Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle triển khai interface Shape
type Rectangle struct {
	Width, Height float64
}

// code bên dưới là : kiểu dữ liệu Rectangle triển khai 2 phương thức trừu tượng của interface Shape : Area() và Perimeter() --> Rectangle là 1 Shape

// Area của Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter (chu vi) của Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Hàm tính tổng diện tích và chu vi của nhiều hình dạng
// kí hiệu ...Shape --> có nghĩa là hàm này có thể nhận một số lượng bất kì các đối tượng Shape
// Bất kì kiểu dữ liệu nào (như Struct) triển khai 2 phương thức : Area() và Perimeter() --> được gọi là 1 Shape
func TotalAreaAndPerimeter(shapes ...Shape) {
	for _, s := range shapes {
		fmt.Printf("Total Area: %f, Total Perimeter: %f\n", s.Area(), s.Perimeter())
	}
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 3, Height: 4}

	TotalAreaAndPerimeter(c, r)

}
