// Đề Bài: Two Sum (LeetCode #1)
// Bài toán: Cho một mảng số nguyên nums và một số nguyên target, hãy tìm hai số trong mảng sao cho tổng của chúng bằng target và trả về chỉ số (index)
// của hai số đó.
// Bạn có thể giả định rằng mỗi đầu vào sẽ có một lời giải duy nhất, và bạn không được phép sử dụng cùng một phần tử hai lần.

package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := numMap[complement]; ok {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := TwoSum(nums, target)
	fmt.Print(result)
}
