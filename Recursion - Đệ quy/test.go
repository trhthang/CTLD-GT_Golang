package main

import "fmt"

func PrintF() {
	fmt.Println("F")
	PrintF()
}

func main() {
	PrintF()
}
