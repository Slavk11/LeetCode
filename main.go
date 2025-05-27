package main

import (
	"LeetCode/abundant_number"
	"LeetCode/check_a_number"
	"LeetCode/diamond_pattern"
	"LeetCode/double_even_numbers"
	"LeetCode/even_and_divisible_numbers"
	"LeetCode/even_and_odd_numbers"
	"LeetCode/generate_pattern"
	"LeetCode/longest_string"
	"LeetCode/non_repeated_indices"
	"LeetCode/number_rotation"
	"LeetCode/rehearsal"
	"LeetCode/slice_to_map"
	"LeetCode/sum"
	"LeetCode/sum_and_divide"
	"LeetCode/sum_average_minimum_and_max"
)

func main() {
	wordsArray := []string{"abc", "coddy", "golang", "java"}
	numbersArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbersForCalculate := []float64{3.5, 2.34, 8.9, 1.6, 4.3}
	arrForSum := []int{2, 4, 6, 7, 8, 8, 9}
	evenAndOddNumbersArray := []int{20, 6, 78, 89, 9, 12, 33}
	sliceToMap := []string{"Coddy", "LearnGolang", "LearnPython", "Challenges"}
	nonRepeatedIndicesArr := []string{"a", "fc", "ab", "a", "ab", "b"}

	sum_and_divide.SumAndDivide(4, 10)
	generate_pattern.GeneratePattern(10)
	number_rotation.RotateNumbers(20, 4)
	even_and_divisible_numbers.DivisibleNumbers(10)
	diamond_pattern.PrintDiamondPattern(5)
	rehearsal.SumAndDivide(4, 10)
	abundant_number.IsAbundantNumber(12)
	check_a_number.CheckNumber(87)
	longest_string.LongestString(wordsArray)
	double_even_numbers.DoubleEvenNumbers(numbersArray)
	sum_average_minimum_and_max.Calculate(numbersForCalculate)
	sum.CalculateSum(arrForSum)
	even_and_odd_numbers.PerformOperations(evenAndOddNumbersArray)
	slice_to_map.SliceToMap(sliceToMap)
	non_repeated_indices.NonRepeatedIndices(nonRepeatedIndicesArr)
}
