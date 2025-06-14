package removesduplicate

import "fmt"

func RemoveDuplicate(inputString string) string {
	if len(inputString) == 0 {
		return ""
	}

	result := []byte{inputString[0]}

	for i := 1; i < len(inputString); i++ {
		if inputString[i] != inputString[i-1] {
			result = append(result, inputString[i])
		}
	}

	fmt.Printf("Input string is %s, Output: %s", inputString, string(result))
	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(n)")
	fmt.Println("==============================")

	return string(result)
}
