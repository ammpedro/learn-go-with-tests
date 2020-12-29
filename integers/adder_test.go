package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {

	assertEqualSum := func(t *testing.T, expected, actual int) {
		t.Helper()
		if actual != expected {
			t.Errorf("expected \"%d\" but got \"%d\"", expected, actual)
		}
	}

	t.Run("should add signed 8-bit integers", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertEqualSum(t, expected, sum)
	})

	t.Run("should add two negative 8-bit integers", func(t *testing.T) {
		sum := Add(-2, -2)
		expected := -4

		assertEqualSum(t, expected, sum)
	})

	t.Run("should allow expressions", func(t *testing.T) {
		sum := Add((1 + 1), (1 - 1))
		expected := 2

		assertEqualSum(t, expected, sum)
	})

	t.Run("should accept and return signed 16-bit integers", func(t *testing.T) {
		sum := Add(-32768, 1)
		expected := -32767

		assertEqualSum(t, expected, sum)
	})

	t.Run("should return signed 32-bit integers", func(t *testing.T) {
		sum := Add(-2147483648, 1)
		expected := -2147483647

		assertEqualSum(t, expected, sum)
	})

	t.Run("should return signed 64-bit integers", func(t *testing.T) {
		sum := Add(-9223372036854775808, 1)
		expected := -9223372036854775807

		assertEqualSum(t, expected, sum)
	})

	t.Run("should not be able to handle 64-bit overflow", func(t *testing.T) {
		sum := Add(-9223372036854775808, -1)
		expected := 9223372036854775807 //this is incorrect

		assertEqualSum(t, expected, sum)
	})

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
