package substringextractor

import "fmt"

func ExtractSubstring(word string, firstIndex, lastIndex int) string {
	var result string

	if firstIndex == -1 || lastIndex == -1 {
		return result
	}

	fmt.Print("🔧 Выполнение: Substring Extractor\n")
	fmt.Printf("Input String: %s\nFirst Index: %d\nLast Index: %d\n", word, firstIndex, lastIndex)

	result = word[firstIndex:lastIndex]

	fmt.Printf("Output string %s\n", result)
	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(n)")
	fmt.Println("==============================")

	return result
}
