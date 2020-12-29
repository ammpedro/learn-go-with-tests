package arraysandslices

import (
	"fmt"
	"reflect"
	"testing"
)

func assertSumOfNumbers(t *testing.T, sum, expected int, input []int) {
	t.Helper()
	if sum != expected {
		t.Errorf("got %d want %d given %v", sum, expected, input)
	}
}

func TestSum(t *testing.T) {

	t.Run("should add array of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		assertSumOfNumbers(t, got, want, numbers)
	})

	// Commenting to check difference in code coverage
	// t.Run("should add array of any size", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3}
	// 	got := Sum(numbers)
	// 	want := 6

	// 	assertSumOfNumbers(t, got, want, numbers)
	// })
}

func assertEqualArrays(t *testing.T, sums, expected []int) {
	t.Helper()
	if !reflect.DeepEqual(sums, expected) {
		t.Errorf("got %v want %v ", sums, expected)
	}
}

func TestSumAll(t *testing.T) {

	t.Run("should return a blank array when input is an empty array", func(t *testing.T) {
		got := SumAll([]int{})
		want := []int{0}

		assertEqualArrays(t, got, want)
	})

	t.Run("should add array of numbers", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3, 4, 5})
		want := []int{15}

		assertEqualArrays(t, got, want)
	})

	t.Run("should add multiple arrays of numbers", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3, 4, 5}, []int{1, 1, 1})
		want := []int{15, 3}

		assertEqualArrays(t, got, want)
	})

}

func TestSumTails(t *testing.T) {

	t.Run("should return the sum of all items apart from the first element", func(t *testing.T) {
		got := SumTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assertEqualArrays(t, got, want)
	})

	t.Run("should safely sum empty slices", func(t *testing.T) {
		got := SumTails([]int{}, []int{1, 1, 1})
		want := []int{0, 2}

		assertEqualArrays(t, got, want)
	})
}

func ExampleSum() {
	sum := Sum([]int{1, 3, 5})
	fmt.Println(sum)
	// Output: 9
}

func ExampleSumAll_single_input() {
	sum := SumAll([]int{1, 2, 3})
	fmt.Println(sum)
	// Output: [6]
}

func ExampleSumAll_multiple_input() {
	sum := SumAll([]int{1, 2, 3}, []int{1, 3, 5}, []int{0, 2, 4})
	fmt.Println(sum)
	// Output: [6 9 6]
}

func ExampleSumTails() {
	sum := SumTails([]int{1, 2, 3}, []int{1, 3, 5}, []int{0, 2, 4})
	fmt.Println(sum)
	// Output: [5 8 6]
}
