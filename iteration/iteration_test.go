package iteration

import (
	"fmt"
	"testing"
)

func assertRepeatedString(t *testing.T, expected, repeated string) {
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeat(t *testing.T) {

	t.Run("shoud return blank if count is zero", func(t *testing.T) {
		repeated := Repeat("#", 0)
		expected := ""

		assertRepeatedString(t, expected, repeated)
	})

	t.Run("shoud return blank if count is negative", func(t *testing.T) {
		repeated := Repeat("#", -99)
		expected := ""

		assertRepeatedString(t, expected, repeated)
	})

	t.Run("shoud accept a character and integer then output a repeated string", func(t *testing.T) {
		repeated := Repeat("b", 5)
		expected := "bbbbb"

		assertRepeatedString(t, expected, repeated)
	})

}

func ExampleRepeat() {
	output := Repeat("x", 5)
	fmt.Println(output)
	// Output: xxxxx
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 100000)
	}
}
