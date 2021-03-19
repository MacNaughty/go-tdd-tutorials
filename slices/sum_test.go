package slices

import "testing"

func TestSum(t1 *testing.T) {

	t1.Run("array of 5 integers", func(t2 *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := SumArray(numbers)
		want := 15
		if got != want {
			t2.Errorf("got %d want %d given, %v", got, want, numbers)
		}

	})

	t1.Run("slice of 6 integers", func(t2 *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}

		got := SumSlice(numbers)
		want := 21
		if got != want {
			t2.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}
