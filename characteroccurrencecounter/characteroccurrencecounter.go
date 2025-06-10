package characteroccurrencecounter

import (
	"fmt"
	"strings"
)

func CharacterOccurrenceCounter(inputString, checkString string) int {
	fmt.Println("Character Occurrence Counter")

	var counter int

	inputString = strings.ToLower(inputString)
	checkString = strings.ToLower(checkString)

	for i := 0; i <= len(inputString)-len(checkString); i++ {
		if inputString[i:i+len(checkString)] == checkString {
			counter++
		}
	}

	fmt.Printf("In string '%s' <%s>: %d times\n", inputString, checkString, counter)
	fmt.Println("==============================")

	return counter
}
