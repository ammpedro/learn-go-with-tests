package convert

import "testing"

func assertConversion(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 should be I", 1, "I"},
		{"2 should be II", 2, "II"},
		{"3 should be III", 3, "III"},
		{"4 should be IV", 4, "IV"},
		{"5 should be V", 5, "V"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			want := test.Want

			assertConversion(t, got, want)
		})
	}
}
