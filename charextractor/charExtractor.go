package charextractor

import "fmt"

func ExtractChar(word, startChar, lastChar string) string {
	fmt.Println("Выполнение: Char Extractor")
	fmt.Printf("Input String: %s\nFirst Char: %s\nLast Char: %s\n", word, startChar, lastChar)

	var result string

	first := -1
	last := -1

	for i, char := range word {
		if string(char) == startChar && first == -1 {
			first = i
		}

		if string(char) == lastChar && last == -1 {
			last = i
		}

		if first > -1 && last > -1 {
			break
		}
	}

	result = word[first : last+1]

	fmt.Printf("Output string is: %s", result)
	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(1)")
	fmt.Println("==============================")

	return result
}
