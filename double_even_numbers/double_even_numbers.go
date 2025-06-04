package double_even_numbers

import "fmt"

func DoubleEvenNumbers(x []int) {
	fmt.Println("🔧 Выполнение: Double Even Numbers\n")
	fmt.Printf("Input Array: %d \n", x)

	var newSlice []int
	for _, number := range x {

		if number%2 == 0 {
			newSlice = append(newSlice, number*2)
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
