package main

func pivotIndex(nums []int) int {
	var sum = 0
	var leftSum = 0
	for _, element := range nums {
		sum = sum + element
	}

	for i := 0; i < len(nums); i++ {

		rightSum := sum - leftSum - nums[i]

		if leftSum == rightSum {
			return i
		}

		leftSum = leftSum + nums[i]
	}

	return -1
}

func main() {

}
