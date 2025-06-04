package abundant_number

import "fmt"

func IsAbundantNumber(x int) {
	fmt.Println("🔧 Выполнение: Abundant_Number\n")
	fmt.Print("(divisors: ")
	var sum int

	for i := 1; i < x; i++ {
		if x%i == 0 {
			sum += i
			fmt.Printf("%d, ", i)
		}
	}

	fmt.Printf("sum: %d)\n", sum)

	if x < sum {
		fmt.Printf("The number %d is Abundant", x)
	} else {
		fmt.Printf("The number %d is not Abundant", x)
	}
	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(1)")
	fmt.Println("==============================")

}
