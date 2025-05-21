package even_and_divisible_numbers

import "fmt"

func DivisibleNumbers(x int) {
	fmt.Println("🔧 Выполнение: DivisibleNumbersn\n")

	var sum int
	var numbers []int

	for num := 2; num < x; num++ {
		if num%2 == 0 || num%6 == 0 {
			numbers = append(numbers, num)
			sum += num
		}
	}
	fmt.Println(numbers, sum)

	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("==============================")

}
