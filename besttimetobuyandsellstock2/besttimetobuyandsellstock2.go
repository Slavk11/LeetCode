package besttimetobuyandsellstock2

import "fmt"

func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	profit := 0

	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			profit += prices[i+1] - prices[i]
		}
	}

	fmt.Printf("Input prices %d\n", prices)
	fmt.Printf("Result %d\n", profit)
	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("Алгоритмическая сложность по CPU - O(n), по памяти O(1)")
	fmt.Println("==============================")

	return profit
}
