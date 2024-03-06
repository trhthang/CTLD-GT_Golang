package main

import (
	"fmt"
)

// prices = [7,1,5,3,6,4]

func maxProfit(prices []int) int {
	var minPrice = prices[0]
	var maxProfit = 0

	for i := 0; i < len(prices); i++ {
		profit := prices[i] - minPrice
		maxProfit = max(maxProfit, profit)
		minPrice = min(minPrice, prices[i])
	}
	return maxProfit
}

func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit(arr)

	fmt.Println(result)
}
