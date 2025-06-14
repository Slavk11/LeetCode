package removeduplicatesfromsortedarray

import "fmt"

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	fmt.Println("Выполнение: Remove Duplicates from Sorted Array")

	result := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[result] = nums[i]
			result++
		}
	}

	fmt.Printf("Input array was: %v, Output is: %v \n", nums, nums[:result])
	fmt.Printf("Result is: %d", result)
	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(1)")
	fmt.Println("==============================")

	return result
}
