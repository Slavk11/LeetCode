package main

import "fmt"

func main() {
	sumAndDivide(4, 10)
}

func sumAndDivide(fistNum int, secondNum int) float64 {
	var sum int
	var result float64

	for i := fistNum; i <= secondNum; i++ {
		sum += i
	}

	result = float64(sum) / (float64(fistNum) + float64(secondNum))
	fmt.Print("The result is ", result)
	return result
}
