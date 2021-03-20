package slices

import (
	"reflect"
	"testing"
)

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

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t1.Run("2D slice of integers, returns a single slice containing sum of each nested slice", func(t2 *testing.T) {

		got := SumAll([]int{1, 2, 3}, []int{2, 3, 4, 5})
		want := []int{6, 14}

		checkSums(t2, got, want)
	})

	t1.Run("Same as SumAll, but not including first element of each slice for each sum", func(t2 *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{2, 3, 4, 5})
		want := []int{5, 12}

		checkSums(t2, got, want)
	})

	t1.Run("safely sum empty slices", func(t2 *testing.T) {
		got := SumAllTails([]int{2}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t2, got, want)
	})

}
