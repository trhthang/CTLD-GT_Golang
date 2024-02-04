package main

import (
	"fmt"
)

type Student struct {
	age  int
	name string
}

func NewStudent(age int, name string) Student {
	return Student{
		age:  age,
		name: name,
	}
}

func (student Student) PrintStudent() {
	fmt.Print("Hoc sinh co ten: ", student.name, " co tuoi la: ", student.age)
}

func main() {
	// #1. Sort an integer array
	// intArray := []int{3, 1, 2, 0}
	// sort.Ints(intArray)
	// fmt.Println(intArray)

	// #2. Sort an String Array
	// stringArray := []string{"banana", "apple", "cherry"}
	// sort.Strings(stringArray)
	// fmt.Print(stringArray)

	// #3. Sort an Float Array
	// floatArray := [3]float64{3.14, 5.33, 9.01}
	// sort.Float64s(floatArray[:])
	// fmt.Print(floatArray)

	student := NewStudent(20, "Thang")

	student.PrintStudent()
}
