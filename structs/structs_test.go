package structs

import "testing"

func TestPerimeter(t1 *testing.T) {

	t1.Run("Find the perimeter of a rectangle, given a length and width", func(t2 *testing.T) {
		got := Perimeter(Rectangle{Length: 12.4, Width: 10.0})
		want := 12.4*2 + 10*2

		if got != want {
			t2.Errorf("got %.2f want %.2f", got, want)
		}
	})

}

func TestArea(t1 *testing.T) {
	t1.Run("Find the area of a rectangle, given a length and width", func(t2 *testing.T) {
		got := Area(Rectangle{Length: 10.0, Width: 10.0})
		want := 100.0

		if got != want {
			t2.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
