package slices

import (
	"fmt"
)

// need to change packages to main to run this
func main() {
	x := [3]string{"Лайка", "Белка", "Стрелка"}

	y := x[:] // slice "y" points to the underlying array "x"

	z := make([]string, len(x))
	copy(z, x[:]) // slice "z" is a copy of the slice created from array "x"

	y[1] = "Belka" // the value at index 1 is now "Belka" for both "y" and "x"
	y = append(y, "Evan")

	fmt.Printf("%T %v\n", x, x)
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)
}
