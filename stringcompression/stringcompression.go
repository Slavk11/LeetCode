package stringcompression

import (
	"fmt"
	"strings"
)

func CompressString(inputString string) string {
	if len(inputString) == 0 {
		return ""
	}

	fmt.Println("Выполнение: String Compression")

	var symbol strings.Builder

	result := make([]string, 0, len(inputString))

	for i := range len(inputString) {
		symbol.WriteByte('#')

		if i == len(inputString)-1 || inputString[i] != inputString[i+1] {
			result = append(result, string(inputString[i]), symbol.String())
			symbol.Reset()
		}
	}

	fmt.Printf("The input string is: %s, result: %s", inputString, strings.Join(result, ""))
	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(1)")
	fmt.Println("==============================")

	return strings.Join(result, "")
}
