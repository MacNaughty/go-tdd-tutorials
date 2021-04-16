// package main
package temp

import (
	"fmt"
)

// will print +Inf
func main() {
	fmt.Println(10.0 / zero())
}

func zero() float64 {
	return 0.0
}
