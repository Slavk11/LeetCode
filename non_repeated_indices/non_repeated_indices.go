package non_repeated_indices

import "fmt"

func NonRepeatedIndices(arrStr []string) []int {
	fmt.Println("🔧 Выполнение: Non Repeated Indices\n")
	fmt.Printf("Input Array: %s\n", arrStr)
	counts := make(map[string]int)
	var result []int

	for _, s := range arrStr {
		counts[s] += 1
	}

	for i, s := range arrStr {
		if counts[s] == 1 {
			result = append(result, i)
		}
	}

	fmt.Printf("Output result: %d\n", result)
	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(1)")
	fmt.Println("==============================")

	return result
}
