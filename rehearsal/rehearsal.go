package rehearsal

import "fmt"

func SumAndDivide(x int, y int) float64 {
	fmt.Println("🔧 Выполнение: Rehearsal\n")
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
