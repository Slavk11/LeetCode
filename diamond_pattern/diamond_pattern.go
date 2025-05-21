package diamond_pattern

import "fmt"

func PrintDiamondPattern(x int) {
	fmt.Println("🔧 Выполнение: Diamond Pattern\n")
	for i := 1; i <= x; i++ {

		for j := 0; j < x-i; j++ {
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
