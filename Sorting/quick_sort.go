/*
Ý tưởng:
	B1: chọn khóa
	B2: Phân bố lại mảng theo khóa
	B3: Chia ra
*/

package main

import "fmt"

func QuickSort(a []int, L int, R int) {
	if L >= R {
		return
	}

	// chọn khóa
	key := a[(L+R)/2] // ví dụ : 0->7, (0+7)/2 = 3

	// phân bố lại mảng theo quá
	pivot := partition(a, L, R, key)

	// chia đôi mảng và lặp lại
	QuickSort(a, L, pivot-1)
	QuickSort(a, pivot, R)
}

// phân chia mảng theo khóa xong trả về 1 giá trị "chốt"
func partition(a []int, L int, R int, key int) int {

	iL := L
	iR := R

	for iL <= iR {
		// Với iL, đi tìm phần tử >= key để đổi chỗ
		for a[iL] < key {
			iL++
		}
		// Với iR, đi tìm phần từ < key để đổi chỗ
		for a[iR] >= key {
			iR--
		}
		// đổi chỗ 2 phần tử đang đứng không đúng vị trí
		if iL <= iR {
			temp := a[iL]
			a[iL] = a[iR]
			a[iR] = temp
			iL++
			iR--
		}
	}

	return iL
}

func main() {
	a := []int{6, 7, 8, 5, 4, 1, 2, 3}
	QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
