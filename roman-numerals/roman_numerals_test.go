package converter

import (
	"fmt"
	"testing"
	"testing/quick"
)

func TestConverter(t *testing.T) {
	cases := []struct {
		Arabic uint16
		Roman  string
	}{
		{Arabic: 1, Roman: "I"},
		{Arabic: 2, Roman: "II"},
		{Arabic: 3, Roman: "III"},
		{Arabic: 4, Roman: "IV"},
		{Arabic: 5, Roman: "V"},
		{Arabic: 6, Roman: "VI"},
		{Arabic: 7, Roman: "VII"},
		{Arabic: 8, Roman: "VIII"},
		{Arabic: 9, Roman: "IX"},
		{Arabic: 10, Roman: "X"},
		{Arabic: 15, Roman: "XV"},
		{Arabic: 18, Roman: "XVIII"},
		{Arabic: 20, Roman: "XX"},
		{Arabic: 24, Roman: "XXIV"},
		{Arabic: 29, Roman: "XXIX"},
		{Arabic: 35, Roman: "XXXV"},
		{Arabic: 39, Roman: "XXXIX"},
		{Arabic: 40, Roman: "XL"},
		{Arabic: 44, Roman: "XLIV"},
		{Arabic: 49, Roman: "XLIX"},
		{Arabic: 50, Roman: "L"},
		{Arabic: 53, Roman: "LIII"},
		{Arabic: 60, Roman: "LX"},
		{Arabic: 90, Roman: "XC"},
		{Arabic: 150, Roman: "CL"},
		{Arabic: 190, Roman: "CXC"},
		{Arabic: 448, Roman: "CDXLVIII"},
		{Arabic: 500, Roman: "D"},
		{Arabic: 900, Roman: "CM"},
		{Arabic: 1000, Roman: "M"},
		{Arabic: 1984, Roman: "MCMLXXXIV"},
		{Arabic: 2014, Roman: "MMXIV"},
		{Arabic: 3999, Roman: "MMMCMXCIX"},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d should be %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ToRoman(test.Arabic)
			want := test.Roman
			if got != want {
				t.Errorf("got %q but want %q", got, want)
			}
		})
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%q should be %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ToArabic(test.Roman)
			want := test.Arabic

			if got != want {
				t.Errorf("got %d but want %d", got, want)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ToRoman(arabic)
		fromRoman := ToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
