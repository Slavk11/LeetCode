package checkanumber

import "fmt"

func CheckNumber(x int) {
	fmt.Println("🔧 Выполнение: Checking Number")

	var sum int

	for i := 1; i < x; i++ {
		if x%i == 0 {
			sum += i
		}
	}

	switch {
	case sum > x:
		fmt.Printf("The number %d is Abundant", x)
	case sum == x:
		fmt.Printf("The number %d is Perfect", x)
	case sum < x:
		fmt.Printf("The number %d is Deficient", x)
	}

	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(1)")
	fmt.Println("==============================")
}
