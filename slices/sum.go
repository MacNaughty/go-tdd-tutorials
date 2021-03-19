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
