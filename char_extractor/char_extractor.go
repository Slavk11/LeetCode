package char_extractor

import "fmt"

func ExtractChar(word string, startChar string, lastChar string) string {
	fmt.Print("🔧 Выполнение: Сhar Extractor\n")
	fmt.Printf("Input String: %s\nFirst Char: %s\nLast Char: %s\n", word, startChar, lastChar)
	var result string
	var first int
	var last int

	for i, r := range word {
		if string(r) == startChar {
			first = i
		}
		if string(r) == lastChar {
			last = i
		}
	}

	for i := first; i <= last; i++ {
		result += string(word[i])
	}

	fmt.Printf("Output string is: %s", result)
	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(1)")
	fmt.Println("==============================")
	return result
}
