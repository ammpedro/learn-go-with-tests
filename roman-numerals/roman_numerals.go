package converter

import "strings"

//RomanNumeral describes a symbol and arabic value
type RomanNumeral struct {
	Value  int
	Symbol string
}

//RomanNumerals describe a map of symbols and values
type RomanNumerals []RomanNumeral

//ValueOf returns value of a roman numeral symbol
func (r RomanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ToRoman accepts int and returns roman numeral string
func ToRoman(arabic int) string {
	var result strings.Builder
	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

// ToArabic accepts roman numeral string and returns arabic int
func ToArabic(roman string) int {
	total := 0

	for i := 0; i < len(roman); i++ {
		total += allRomanNumerals.ValueOf(roman[i])

		if i != 0 && roman[i] == 'V' {
			if roman[i-1] == 'I' {
				total -= 2
			}
		}

		if i != 0 && roman[i] == 'X' {
			if roman[i-1] == 'I' {
				total -= 2
			}
		}

		if i != 0 && roman[i] == 'L' {
			if roman[i-1] == 'X' {
				total -= 20
			}
		}

		if i != 0 && roman[i] == 'C' {
			if roman[i-1] == 'X' {
				total -= 20
			}
		}

		if i != 0 && roman[i] == 'D' {
			if roman[i-1] == 'C' {
				total -= 200
			}
		}

		if i != 0 && roman[i] == 'M' {
			if roman[i-1] == 'C' {
				total -= 200
			}
		}
	}

	return total
}
