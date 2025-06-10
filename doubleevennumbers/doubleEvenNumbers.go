package doubleevennumbers

import "fmt"

func DoubleEvenNumbers(x []int) {
	fmt.Println("Выполнение: Double Even Numbers")
	fmt.Printf("Input Array: %d \n", x)

	var newSlice []int

	const doublingFactor = 2

	for _, number := range x {
		if number%2 == 0 {
			newSlice = append(newSlice, number*doublingFactor)
		} else {
			newSlice = append(newSlice, number)
		}
	}

	fmt.Printf("Output Slice: %d", newSlice)

	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(n)")
	fmt.Println("==============================")
}
