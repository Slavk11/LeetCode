package slice_to_map

import "fmt"

func SliceToMap(arr []string) map[string]int {
	fmt.Println("🔧 Выполнение: Slice To Map\n")
	testMap := make(map[string]int)
	for _, str := range arr {
		testMap[str] = len(str)
	}
	fmt.Printf("Inpit Array: %s\n", arr)
	fmt.Println("Result Map:")
	fmt.Print(testMap)
	return testMap

}
