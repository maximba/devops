// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

package main

import (
	"fmt"

	"./lib"
)

func main() {
	const target = 1000
	var a, b, c, py, bound int
	var found = false

	bound = int(mxlib.IntSqrt(target))

	for i := 1; !found && i < bound; i++ {
		for j := 1; !found && j < bound; j++ {
			for k := 1; !found && k < target; k++ {
				i2, j2, k2 := i*i, j*j, k*k
				py = i2 + j2 + k2
				if py == target {
					found = true
					a, b, c = i, j, k
				}
			}
		}
	}
	fmt.Printf("The terna is: %d^2 + %d^2 + %d^2 = 1000\n", a, b, c)
}
