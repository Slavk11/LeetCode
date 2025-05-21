package number_rotation

import "fmt"

func RotateNumbers(firstValue int, secondValue int) {
	fmt.Println("🔧 Выполнение: Number Rotation\n")
	for a := firstValue; a >= 0; a -= secondValue {
		fmt.Println(a)
	}
	fmt.Println("\n==============================")
	fmt.Println("   ✅ Программа завершена    ")
	fmt.Println("==============================")
}
