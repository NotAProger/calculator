package calc

import (
	"calculator/internal/tranclator"
	"fmt"
	"strconv"
	"strings"
)

const reqNumberQuantity int = 3

func Run(s string) string {
	s = s[:len(s)-1]
	nums := strings.Split(s, " ")

	validateInput(nums)

	isRome := tranclator.IsRoman(nums[0])
	numsInt := convertToInteger(nums, isRome)

	answer := calculate(nums, numsInt)

	return formatAnswer(answer, isRome)
}

func validateInput(nums []string) {
	if len(nums) != reqNumberQuantity {
		panic("the operation format does not meet the requirements—two operands and one operator")
	}
}

func convertToInteger(nums []string, isRome bool) []int {
	numsInt := []int{}
	_ = numsInt
	if isRome && tranclator.IsRoman(nums[2]) {
		numsInt = convertRomanNumbers(nums)
	} else if !isRome && !tranclator.IsRoman(nums[2]) {
		numsInt = convertArabicNumbers(nums)
	} else {
		panic("different numeral systems are being used simultaneously")
	}
	return numsInt
}

func convertRomanNumbers(nums []string) []int {
	numsInt := []int{}
	_ = numsInt
	for i := 0; i < len(nums); i += 2 {
		num := tranclator.RomanToArabic(nums[i])
		if num < 1 || num > 10 {
			panic(fmt.Errorf("roman number \"%s\" is out of acceptable range", nums[i]))
		}
		numsInt = append(numsInt, num)
	}
	return numsInt
}

func convertArabicNumbers(nums []string) []int {
	numsInt := []int{}
	for i := 0; i < len(nums); i += 2 {
		num := tranclator.TranclateArab(nums[i])
		if num < -10 || num > 10 {
			panic(fmt.Errorf("arab number \"%d\" is out of acceptable range", num))
		}
		numsInt = append(numsInt, num)
	}
	return numsInt
}

func calculate(s []string, i []int) int {
	switch s[1] {
	case "*":
		return i[0] * i[1]
	case "/":
		return i[0] / i[1]
	case "+":
		return i[0] + i[1]
	case "-":
		return i[0] - i[1]
	default:
		panic(fmt.Errorf("symbol(s) \"%s\" is not supported", s[1]))
	}
}

func formatAnswer(answer int, isRome bool) string {
	if isRome {
		if answer > 0 {
			return tranclator.ArabicToRoman(answer)
		}
		panic("romans didnt invent negative numbers and 0 ;(")
	}
	return strconv.Itoa(answer)
}
