package main

import "fmt"

func main() {
	var mySlice []int
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 4)

	mySlice2 := mySlice[:3]
	fmt.Printf("%d ", mySlice2)
}
