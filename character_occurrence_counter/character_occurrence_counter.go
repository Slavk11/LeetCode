package character_occurrence_counter

import (
	"fmt"
	"strings"
)

func CountOccurrences(inputStr string, checkStr string) int {
	fmt.Print("🔧 Выполнение: Character Occurrence Counter\n")
	var counter int
	inputLen := len(inputStr)
	checkLen := len(checkStr)

	inputStr = strings.ToLower(inputStr)
	checkStr = strings.ToLower(checkStr)

	for i := 0; i <= (inputLen - checkLen); i++ {
		if inputStr[i:i+checkLen] == checkStr {
			counter++
		}
	}
	fmt.Printf("In string '%s' <%s>: %d times\n", inputStr, checkStr, counter)
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(1)")
	fmt.Println("==============================")
	return counter
}
