package main

import (
	"LeetCode/abundant_number"
	"LeetCode/check_a_number"
	"LeetCode/diamond_pattern"
	"LeetCode/even_and_divisible_numbers"
	"LeetCode/generate_pattern"
	"LeetCode/number_rotation"
	"LeetCode/rehearsal"
	"LeetCode/sum_and_divide"
)

func main() {
	sum_and_divide.SumAndDivide(4, 10)
	generate_pattern.GeneratePattern(10)
	number_rotation.RotateNumbers(20, 4)
	even_and_divisible_numbers.DivisibleNumbers(10)
	diamond_pattern.PrintDiamondPattern(5)
	rehearsal.SumAndDivide(4, 10)
	abundant_number.IsAbundantNumber(12)
	check_a_number.CheckNumber(87)
}
