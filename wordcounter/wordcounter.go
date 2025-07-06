package wordcounter

import "fmt"

func CountWords(sentence string) int {

	if len(sentence) == 0 {
		return 0
	}

	var counter int

	for i := 0; i < len(sentence); i++ {
		if sentence[i] != ' ' && (i == 0 || sentence[i-1] == ' ') {
			counter++
		}
	}

	fmt.Printf("В предложении: %s, %d: слова ", sentence, counter)
	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(1)")
	fmt.Println("==============================")

	return counter
}
