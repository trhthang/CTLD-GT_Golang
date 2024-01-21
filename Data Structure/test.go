package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for _, ai := range nums2 {
		insertNewElementToArray(ai, &nums1, m)
		m++
	}
}

func insertNewElementToArray(x int, nums1 *[]int, m int) {
	timDuocK := false
	for k := 0; k < m; k++ {
		if (*nums1)[k] > x {
			timDuocK = true

			// Dịch chuyển các phần tử trong mảng nums1
			for i := m - 1; i >= k; i-- {
				(*nums1)[i+1] = (*nums1)[i]
			}

			(*nums1)[k] = x
			break
		}
	}

	if !timDuocK {
		(*nums1)[m] = x
	}
}

func main() {
	n1 := []int{2, 3, 4, 5, 0, 0, 0}
	n2 := []int{2, 3, 6}

	merge(n1, 4, n2, 3)

	for _, ai := range n1 {
		fmt.Println(ai)
	}
}
