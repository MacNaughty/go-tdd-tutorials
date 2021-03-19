package slices

func Sum(intSlice [5]int) (acc int) {
	for i := 0; i < len(intSlice)-1; i++ {
		acc += intSlice[i]
	}
	return
}
