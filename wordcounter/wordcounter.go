package wordcounter

import "fmt"

func CountWords(sentence string) int {
	const space byte = 32

	if len(sentence) == 0 {
		return 0
	}

	var counter int

	for i := range len(sentence) {
		if sentence[i] == space {
			counter++
		}
	}

	if sentence[0] == space {
		counter--
	}

	if sentence[len(sentence)-1] == space {
		counter--
	}

	fmt.Printf("В предложении: %s, %d: слова ", sentence, counter+1)
	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(1)")
	fmt.Println("==============================")

	return counter + 1
}
