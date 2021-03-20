package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t1 *testing.T) {

	t1.Run("Find the perimeter of a rectangle, given a length and width", func(t2 *testing.T) {
		//rectangle :=
		got := Rectangle{Length: 12.4, Width: 10.0}.Perimeter()
		want := 12.4*2 + 10*2

		if got != want {
			t2.Errorf("got %.2f want %.2f", got, want)
		}
	})

}

func TestArea(t1 *testing.T) {
	t1.Run("Find the area of a rectangle, given a length and width", func(t2 *testing.T) {
		rectangle := Rectangle{Length: 15.0, Width: 15.0}
		got := rectangle.Area()
		want := 225.0

		if got != want {
			t2.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t1.Run("Circles", func(t2 *testing.T) {
		circle := Circle{Radius: 10}
		got := circle.Area()
		want := math.Pi * 100

		if got != want {
			t2.Errorf("got %.2f want %.2f", got, want)
		}

	})
}
