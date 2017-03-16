// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

package main

import (
	"fmt"
)

func ternaPitagorica() (int, int, int) {
	const bound = 1000
	var a, b, c int
	var found = false

	for i := 1; !found && i < bound; i++ {
		for j := 1; !found && j < bound; j++ {
			for k := 1; !found && k < bound; k++ {
				i2, j2, k2 := i*i, j*j, k*k
				if (i2+j2 == k2) && (i+j+k == 1000) {
					found = true
					a, b, c = i, j, k
				}
			}
		}
	}
	return a, b, c
}

func main() {
	a, b, c := ternaPitagorica()
	fmt.Printf("The terna is: %d^2 + %d^2 = %d^2\n", a, b, c)
	fmt.Printf("Where: %d + %d + %d = %d\n", a, b, c, a+b+c)
	fmt.Printf("The product is a*b*c = %d\n", a*b*c)
}
