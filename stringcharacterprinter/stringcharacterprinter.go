package stringcharacterprinter

import (
	"fmt"
)

func PrintCharacters(inputString string) {
	if len(inputString) == 0 {
		return
	}

	fmt.Println("Выполнение: Character Occurrence Counter")
	fmt.Printf("Input String: %s\n", inputString)

	for i := 1; i <= len(inputString); i++ {
		fmt.Println(inputString[:i])
	}

	for i := len(inputString) - 1; i > 0; i-- {
		fmt.Println(inputString[:i])
	}

	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(1)")
	fmt.Println("==============================")
}
