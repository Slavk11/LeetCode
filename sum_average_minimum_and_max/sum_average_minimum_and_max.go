package sum_average_minimum_and_max

import "fmt"

func Calculate(x []float64) {
	fmt.Println("🔧 Выполнение: Sum Average Minimum and Max\n")

	fmt.Printf("Here is a Slice with elements: %g\n", x)

	var sum float64
	var average float64
	numbers := len(x)
	minimumValue := 9999999.0
	maximumValue := -9999999.0
	for _, number := range x {
		sum += number
		average = sum / float64(numbers)

		if number < minimumValue {
			minimumValue = number
		}

		if number > maximumValue {
			maximumValue = number
		}
	}

	fmt.Printf(
		"Elements sum = %g,"+
			" the average number is %g,"+
			" minumum value is %g,"+
			" maximum value is %g",
		sum, average, maximumValue, maximumValue)

	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(1)")
	fmt.Println("==============================")
}
