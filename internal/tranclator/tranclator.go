package tranclator

import "strconv"

func TranclateArab(s string) int {
	var number int
	var err error

	if number, err = strconv.Atoi(s); err != nil {
		panic(err)
	}
	return number
}

func IsRoman(s string) bool {
	for _, r := range s {
		if r >= 48 && r <= 57 || r == '-' {
			return false
		}
	}
	return true
}

func RomanToArabic(roman string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	prevValue := 0

	for _, r := range roman {
		value := romanMap[r]
		if value > prevValue {
			total += value - 2*prevValue
		} else {
			total += value
		}
		prevValue = value
	}

	return total
}

func ArabicToRoman(num int) string {
	romanMap := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	var roman string
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for _, value := range values {
		for num >= value {
			roman += romanMap[value]
			num -= value
		}
	}
	return roman
}
