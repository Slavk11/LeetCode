package stringreversalandduplication

import "fmt"

func ReverseAndDouble(inputString string) string {
	fmt.Println("Выполнение: String Reversal And Duplication")

	const charRepeat = 2

	strLen := len(inputString)
	outputString := make([]byte, charRepeat*strLen)

	fmt.Printf("Input String: %s\n", inputString)

	for i := range strLen {
		char := inputString[strLen-1-i]
		outputString[2*i] = char
		outputString[2*i+1] = char
	}

	fmt.Printf("The result is '%s", outputString)

	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(n)")
	fmt.Println("==============================")

	return string(outputString)
}
