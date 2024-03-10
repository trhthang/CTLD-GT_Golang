package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	var cleanedString []rune
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			cleanedString = append(cleanedString, unicode.ToLower(char))
		}
	}

	length := len(cleanedString)
	for i := 0; i < length/2; i++ {
		if cleanedString[i] != cleanedString[length-i-1] {
			return false
		}
	}

	return true
}

func main() {
	examples := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		"",
	}

	for _, example := range examples {
		fmt.Printf("Đầu vào: %s - Kết quả: %t\n", example, isPalindrome(example))
	}
}
