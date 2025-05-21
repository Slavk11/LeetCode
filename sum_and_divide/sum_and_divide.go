package sum_and_divide

import "fmt"

func SumAndDivide(fistNum int, secondNum int) float64 {
	var sum int
	var result float64

	for i := fistNum; i <= secondNum; i++ {
		sum += i
	}

	result = float64(sum) / (float64(fistNum) + float64(secondNum))
	fmt.Println("🔧 Выполнение: Sum & Divide и Generate Pattern\n")

	fmt.Println("The result is ", result)

	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("==============================")
	return result
}
