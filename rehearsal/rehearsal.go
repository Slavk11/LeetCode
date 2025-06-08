package rehearsal

import "fmt"

func SumAndDivide(x, y int) float64 {
	fmt.Println("🔧 Выполнение: Rehearsal")

	var sum int

	var result float64

	for i := x; i <= y; i++ {
		sum += i
	}

	result = float64(sum) / (float64(x) + float64(y))
	fmt.Print(result)

	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("==============================")

	return result
}
