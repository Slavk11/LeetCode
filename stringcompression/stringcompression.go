package stringcompression

import (
	"fmt"
)

func CompressString(inputString string) {
	if len(inputString) == 0 {
		return
	}

	fmt.Println("Выполнение: String Compression")

	counter := 0
	for i := 0; i < len(inputString); i++ {
		counter++

		if i == len(inputString)-1 || inputString[i] != inputString[i+1] {
			fmt.Printf("%c", inputString[i])
			for j := 0; j < counter; j++ {
				fmt.Print("#")
			}
			counter = 0
		}
	}

	fmt.Printf("\nThe input string is: %s\n", inputString)
	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(1)")
	fmt.Println("==============================")
}
