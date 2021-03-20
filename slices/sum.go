package slices

func SumArray(intArray [5]int) (acc int) {
	for _, number := range intArray {
		acc += number
	}
	return
}

func SumSlice(intSlice []int) (acc int) {
	for _, number := range intSlice {
		acc += number
	}
	return
}

//func SumAll(numbersSlice ...[]int)  []int {
//	numbersSliceLength := len(numbersSlice)
//	sums := make([]int, numbersSliceLength)
//	for i, slice := range numbersSlice {
//		sums[i] = SumSlice(slice)
//	}
//	return sums
//}

func SumAll(numbersSlice ...[]int) []int {
	var sums []int
	for _, numbers := range numbersSlice {
		sums = append(sums, SumSlice(numbers))
	}
	return sums
}

func SumAllTails(numbersSlice ...[]int) []int {
	var sums []int
	for _, slice := range numbersSlice {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			tail := slice[1:]
			sums = append(sums, SumSlice(tail))
		}
	}
	return sums
}
