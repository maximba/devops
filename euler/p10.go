// Find the sum of all the primes below two million.

package main

import (
	"fmt"

	"./lib"
)

func main() {
	const bound = 2000000
	sum := 0

	for i := 2; i < bound; i++ {
		if mxlib.Is_prime(i) {
			sum += i
		}
	}

	fmt.Printf("Sum of All primes below 2M is: %d\n", sum)
}
