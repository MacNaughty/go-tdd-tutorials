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

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, math.Pi * 100},
		{Triangle{12, 6}, 36.0},
	}

	for _, t2 := range areaTests {
		got := t2.shape.Area()
		if got != t2.want {
			t1.Errorf("%#v got %g want %g", t2.shape, got, t2.want)
		}
	}

}
