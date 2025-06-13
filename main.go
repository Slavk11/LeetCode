package main

import (
	"LeetCode/abundantnumber"
	"LeetCode/adddigits"
	"LeetCode/characteroccurrencecounter"
	"LeetCode/charextractor"
	"LeetCode/checkanumber"
	"LeetCode/diamondpattern"
	"LeetCode/doubleevennumbers"
	"LeetCode/evenanddivisiblenumbers"
	"LeetCode/evendndoddnumbers"
	"LeetCode/generatepattern"
	"LeetCode/longeststring"
	"LeetCode/nonrepeatedindices"
	"LeetCode/numberrotation"
	"LeetCode/rehearsal"
	"LeetCode/removeduplicatesfromsortedarray"
	"LeetCode/slicetomap"
	"LeetCode/stringcompression"
	"LeetCode/stringreversalandduplication"
	"LeetCode/stringtokenizer"
	"LeetCode/substringextractor"
	"LeetCode/sum"
	"LeetCode/sumanddivide"
	"LeetCode/sumaverageminimumandmax"
	"LeetCode/wordcounter"
)

func main() {
	wordsArray := []string{"abc", "coddy", "golang", "java"}
	numbersArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbersForCalculate := []float64{3.5, 2.34, 8.9, 1.6, 4.3}
	arrForSum := []int{2, 4, 6, 7, 8, 8, 9}
	evenAndOddNumbersArray := []int{20, 6, 78, 89, 9, 12, 33}
	stringArray := []string{"Coddy", "LearnGolang", "LearnPython", "Challenges"}
	nonRepeatedIndicesArr := []string{"a", "fc", "ab", "a", "ab", "b"}
	removeDuplicates := []int{1, 1, 2}
	lastIndex := 7
	counter := 6
	checkNumber := 87
	firstNumber := 4
	lastNumber := 10

	adddigits.AddDigits(firstNumber)
	sumanddivide.SumAndDivide(firstNumber, lastNumber)
	generatepattern.GeneratePattern(lastNumber)
	numberrotation.RotateNumbers(checkNumber, lastIndex)
	evenanddivisiblenumbers.DivisibleNumbers(lastNumber)
	diamondpattern.PrintDiamondPattern(firstNumber)
	rehearsal.SumAndDivide(firstNumber, lastNumber)
	abundantnumber.IsAbundantNumber(checkNumber)
	checkanumber.CheckNumber(checkNumber)
	longeststring.LongestString(wordsArray)
	doubleevennumbers.DoubleEvenNumbers(numbersArray)
	sumaverageminimumandmax.Calculate(numbersForCalculate)
	sum.CalculateSum(arrForSum)
	evendndoddnumbers.PerformOperations(evenAndOddNumbersArray)
	slicetomap.SliceToMap(stringArray)
	nonrepeatedindices.NonRepeatedIndices(nonRepeatedIndicesArr)
	stringtokenizer.InsertSpace(counter, "cooddypythongolang")
	charextractor.ExtractChar("shvjng", "s", "j")
	substringextractor.ExtractSubstring("abshvjngsh", firstNumber, lastIndex)
	characteroccurrencecounter.CharacterOccurrenceCounter("PythonMagick", "a")
	stringreversalandduplication.ReverseAndDouble("abcdfgh")
	removeduplicatesfromsortedarray.RemoveDuplicates(removeDuplicates)
	wordcounter.CountWords("Learn Golang with Coddy")
	stringcompression.CompressString("aabccchbbccaaa")
}
