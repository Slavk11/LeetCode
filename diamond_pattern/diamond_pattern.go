package diamond_pattern

import "fmt"

func PrintDiamondPattern(x int) {
	for i := 1; i <= x; i++ {

		for j := 0; j < x-i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= i; k++ {
			fmt.Printf("%d ", k)
		}

		fmt.Println()
	}
}
