package converter

import "strings"

//RomanNumeral describes a symbol and arabic value
type RomanNumeral struct {
	Value  uint16
	Symbol string
}

//RomanNumerals describe a map of symbols and values
type RomanNumerals []RomanNumeral

//ValueOf returns value of a roman numeral symbol
func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
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
func ToRoman(arabic uint16) string {
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
func ToArabic(roman string) uint16 {
	var total uint16 = 0

	for i := 0; i < len(roman); i++ {
		total = total + allRomanNumerals.ValueOf(roman[i])

		switch {
		case i != 0 && (roman[i] == 'V' || roman[i] == 'X') && roman[i-1] == 'I':
			total -= 2
		case i != 0 && (roman[i] == 'L' || roman[i] == 'C') && roman[i-1] == 'X':
			total -= 20
		case i != 0 && (roman[i] == 'D' || roman[i] == 'M') && roman[i-1] == 'C':
			total -= 200
		default:
			total += 0
		}
	}

	return total
}
