package generate_pattern

import "fmt"

func GeneratePattern(n int) {
	fmt.Println("🔧 Выполнение: Generate Pattern и Generate Pattern\n")

	if n > 10 {
		fmt.Print("The number should be less than 10")
	}
	for counter := n; counter > 0; counter-- {
		for i := 1; i < counter; i++ {
			fmt.Print(i)
		}
		fmt.Println(" ")
	}

	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("==============================")
}
