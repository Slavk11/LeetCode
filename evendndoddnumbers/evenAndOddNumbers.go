package evendndoddnumbers

import "fmt"

func PerformOperations(arr []int) int {
	fmt.Println("Выполнение: Even & Odd Numbers")

	var even int

	odd := 1

	for _, number := range arr {
		if number%2 == 0 {
			even += number
		} else {
			odd *= number
		}
	}

	result := even + odd

	fmt.Printf("Input array: %d\n", arr)
	fmt.Printf("Result is %d\n", result)

	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(1)")
	fmt.Println("==============================")

	return result
}
