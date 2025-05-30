package substring_extractor

import "fmt"

func ExtractSubstring(word string, firstIndex int, lastIndex int) string {
	var result string
	if firstIndex == -1 || lastIndex == -1 {
		return result
	}
	fmt.Print("🔧 Выполнение: Substring Extractor\n")
	fmt.Printf("Input String: %s\nFirst Index: %d\nLast Index: %d\n", word, firstIndex, lastIndex)

	for i := firstIndex; i <= lastIndex; i++ {
		result += string(word[i])
	}
	fmt.Printf("Output string %s\n", result)
	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(1)")
	fmt.Println("==============================")
	return result
}
