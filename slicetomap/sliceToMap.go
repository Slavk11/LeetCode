package slicetomap

import "fmt"

func SliceToMap(arr []string) map[string]int {
	fmt.Println("Выполнение: Slice To Map")

	mapFromSlice := make(map[string]int, len(arr))
	for _, str := range arr {
		mapFromSlice[str] = len(str)
	}

	fmt.Printf("Input Array: %s\n", arr)
	fmt.Println("Result Map:", mapFromSlice)

	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(n)")
	fmt.Println("==============================")

	return mapFromSlice
}
