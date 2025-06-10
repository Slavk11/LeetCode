package diamondpattern

import "fmt"

func PrintDiamondPattern(x int) {
	fmt.Println("Выполнение: Diamond Pattern")

	for i := 1; i <= x; i++ {
		for range x - i {
			fmt.Print(" ")
		}

		for k := 1; k <= i; k++ {
			fmt.Printf("%d ", k)
		}

		fmt.Println()
	}

	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("==============================")
}
