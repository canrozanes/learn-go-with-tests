package main

// Sum adds elements of an integer array
func Sum(array []int) int {
	var sum int
	for _, number := range array {
		sum += number
	}
	return sum
}

// SumAll takes a varying number of slices, returning a new slice containing the totals for each slice passed in.
func SumAll(arrays ...[]int) []int {
	var sums []int

	for _, array := range arrays {
		sums = append(sums, Sum(array))
	}

	return sums
}

// SumAllTails calculates the totals of the "tails" of each slice. The tail of a collection is all the items apart from the first one (the "head")
func SumAllTails(slices ...[]int) []int {
	var sums []int

	for _, array := range slices {
		if len(array) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(array[1:]))
		}
	}

	return sums
}

func main() {

}
