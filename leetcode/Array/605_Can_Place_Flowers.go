package main

import "fmt"

// canPlaceFlowers kiểm tra xem có thể trồng thêm n bông hoa mới vào flowerbed mà không vi phạm quy tắc không có hoa liền kề nhau
func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0 // số lượng hoa có thể trồng được

	// Duyệt qua flowerbed
	for i := 0; i < len(flowerbed) && count < n; i++ {
		// Kiểm tra nếu ô hiện tại trống và là (là ô đầu tiên hoặc ô trước đó trống)
		// và (là ô cuối cùng hoặc ô tiếp theo trống)
		if flowerbed[i] == 0 &&
			(i == 0 || flowerbed[i-1] == 0) &&
			(i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1 // trồng hoa vào ô đó
			count++          // Tăng số lượng hoa đã trồng
			if count >= n {  // Nếu đã trồng đủ số lượng hoa, trả về true
				return true
			}
		}
	}
	return count >= n // kiểm tra nếu đã trồng đủ số lượng
}

func main() {
	// các trường hợp kiểm thử
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1)) // Đầu ra: true
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2)) // Đầu ra: false
}
