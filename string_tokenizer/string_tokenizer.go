package string_tokenizer

import "fmt"

func InsertSpace(counter int, str string) string {
	fmt.Println("🔧 Выполнение: String Tokenizer\n")
	fmt.Printf("You want to use space after %d chars\n", counter)
	fmt.Printf("Input string: %s\n", str)
	var result string
	for i := 0; i < len(str); i++ {
		if i != 0 && i%counter == 0 {
			result += " "
		}
		result += string(str[i])
	}
	fmt.Printf("Output words: %s\n", result)

	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(n²)")
	fmt.Println("==============================")

	return result
}
