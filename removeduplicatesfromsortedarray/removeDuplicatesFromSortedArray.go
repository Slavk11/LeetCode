package removeduplicatesfromsortedarray

import "fmt"

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[result] = nums[i]
			result++
		}
	}

	fmt.Print(result)

	return result
}
