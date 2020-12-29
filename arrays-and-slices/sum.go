package arraysandslices

// Sum accepts an array of integers and returns its sum
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll accepts a set of arrays and returns its sums as an array
func SumAll(numberArrays ...[]int) (sums []int) {
	for _, numbers := range numberArrays {
		sums = append(sums, Sum(numbers))
	}
	return
}

// SumTails accepts a set of array and returns the sum of its tail elements
func SumTails(numberArrays ...[]int) (sums []int) {
	for _, numbers := range numberArrays {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}
