package main

func pivotIndex(nums []int) int {
	var sum int
	var resultIndex int = -1

	for _, element := range nums {
		sum = sum + element
	}

	for i := 0; i < len(nums); i++ {
		resultIndex = i
		// tại i tính tổng bên trái và bên phải --> nếu bằng nhau thì return i
		var leftSum int = 0
		var rightSum int

		for j := 0; j < i; j++ {
			leftSum = leftSum + nums[j]
		}
		rightSum = sum - leftSum - nums[i]

		if leftSum == rightSum {
			return resultIndex
		}

		if resultIndex == 0 || resultIndex == len(nums)-1 {
			return 0
		}
	}

	return resultIndex
}

func main() {

}
