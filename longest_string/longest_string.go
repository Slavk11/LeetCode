package longest_string

import (
	"fmt"
	"unicode/utf8"
)

func LongestString(x []string) {
	fmt.Println("🔧 Выполнение: Longest String\n")

	var result int
	var longestString string

	fmt.Println("Input Words:")

	for i := 0; i < len(x); i++ {

		fmt.Println(x[i])
		counter := utf8.RuneCountInString(x[i])

		if counter > result {
			result = counter
			longestString = x[i]
		}
	}
	fmt.Printf("The longest word in array is '%s' %d chars", longestString, result)

	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(1)")
	fmt.Println("==============================")
}
