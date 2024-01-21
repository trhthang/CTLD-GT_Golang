package main

import "fmt"

// Định nghĩa struct cơ bản Vehicle
type Vehicle struct {
	MadeBy  string
	Version string
}

// Phương thức của Vehicle
func (v Vehicle) DisplayDetails() {
	fmt.Printf("Made by: %s, Version: %s\n", v.MadeBy, v.Version)
}

// Định nghĩa struct Car, "kế thừa" từ Vehicle
type Car struct {
	Vehicle    //Nhúng struct Vehicle vào Car
	IsElectric bool
}

// Định nghĩa riêng của Car
func (c Car) DisplayCarDetails() {
	c.DisplayDetails() // Gọi phương thức từ Vehicle
	fmt.Printf("Is Electric: %t\n", c.IsElectric)
}

func main() {
	myCar := Car{
		Vehicle:    Vehicle{MadeBy: "Tesla", Version: "Model 3"},
		IsElectric: true,
	}

	myCar.DisplayCarDetails()
}

// So Sánh:
// Pointer Receiver (c *Car): Thay đổi trạng thái của đối tượng gốc. Sử dụng cho các hàm cần sửa
// đổi đối tượng hoặc để tối ưu hiệu suất với các struct lớn (do tránh sao chép dữ liệu).

// Value Receiver (c Car): Làm việc với bản sao của đối tượng, không ảnh hưởng đến đối tượng gốc.
// Phù hợp cho các phương thức đọc hoặc các hàm không cần sửa đổi đối tượng.
