package main

import "fmt"

var romanToArabic = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if i > 0 && romanToArabic[rune(s[i])] > romanToArabic[rune(s[i-1])] {
			result += romanToArabic[rune(s[i])] - 2*romanToArabic[rune(s[i-1])]
		} else {
			result += romanToArabic[rune(s[i])]
		}
	}
	return result
}

func isValidRoman(s string) bool {
	count := 1
	prev := rune(s[0])
	for i := 1; i < len(s); i++ {
		if rune(s[i]) == prev {
			count++
			if count > 4 {
				return false
			}
		} else {
			count = 1
		}
		if romanToArabic[rune(s[i])] > romanToArabic[prev] {
			// IX - 9
			if (s[i] == 'V' || s[i] == 'X') && prev == 'I' {
				continue
			}
			// XC - 90
			if (s[i] == 'L' || s[i] == 'C') && prev == 'X' {
				continue
			}
			// CM - 900
			if (s[i] == 'D' || s[i] == 'M') && prev == 'C' {
				continue
			}
			return false
		}
		prev = rune(s[i])
	}
	return true
}

func isRomanChar(c rune) bool {
	_, exists := romanToArabic[c]
	return exists
}

func intToRoman(num int) string {
	result := ""
	romanSymbols := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	bases := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for _, base := range bases {
		for num >= base {
			result += romanSymbols[base]
			num -= base
		}
	}

	return result
}

func main() {
	var input string
	validInput := false

	for !validInput {
		fmt.Print("Введите римское число: ")
		fmt.Scanln(&input)

		for _, c := range input {
			if !isRomanChar(c) {
				fmt.Println("Ошибка: неверный символ римского числа.")
				validInput = false
				break
			} else {
				validInput = true
			}
		}

		if !isValidRoman(input) {
			fmt.Println("Ошибка: неверный формат римского числа.")
			validInput = false
		}
	}

	fmt.Println("Арабское представление:", romanToInt(input))
	number := romanToInt(input)
	fmt.Println("Римское представление:", intToRoman(number))
}
