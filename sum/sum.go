package sum

import "fmt"

func CalculateSum(x []int) int {
	fmt.Println("🔧 Выполнение: Sum\n")

	checkEvenOdd := len(x)
	var result int
	if checkEvenOdd%2 == 0 {
		result = -1
	} else {
		middle := x[len(x)/2]
		result = x[0] + middle + x[len(x)-1]
	}
	fmt.Printf("First number is %d, middle number is %d and last number is %d\n", x[0], x[len(x)/2], x[len(x)-1])
	fmt.Printf("The result is %d", result)

	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(1)")
	fmt.Println("==============================")
	return result
}
