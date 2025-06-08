package numberrotation

import "fmt"

func RotateNumbers(firstValue, secondValue int) {
	fmt.Println("Выполнение: Number Rotation")

	for a := firstValue; a >= 0; a -= secondValue {
		fmt.Println(a)
	}

	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("==============================")
}
