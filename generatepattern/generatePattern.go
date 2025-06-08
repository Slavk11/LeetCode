package generatepattern

import "fmt"

func GeneratePattern(n int) {
	fmt.Println("Выполнение: Generate Pattern")

	const maxAllowed = 10

	if n > maxAllowed {
		fmt.Print("The number should be less than 10")
	} else {
		for counter := n; counter > 0; counter-- {
			for i := 1; i <= counter; i++ {
				fmt.Print(i)
			}

			fmt.Println()
		}
	}

	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("==============================")
}
