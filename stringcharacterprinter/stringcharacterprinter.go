package stringcharacterprinter

import (
	"fmt"
	"strings"
)

func PrintCharacters(inputString string) {
	if len(inputString) == 0 {
		return
	}

	fmt.Println("Выполнение: Character Occurrence Counter")
	fmt.Printf("Input String: %s\n", inputString)

	var builder strings.Builder

	for _, ch := range inputString {
		builder.WriteRune(ch)
		fmt.Println(builder.String())
	}

	result := builder.String()

	for i := len(result) - 1; i > 0; i-- {
		fmt.Println(result[:i])
	}

	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(n)")
	fmt.Println("==============================")
}
