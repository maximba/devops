package main

// If we list all the natural numbers below 10 that are 
// multiples of 3 or 5, we get 3, 5, 6 and 9. 
// The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

import (
    "fmt"
)

func isMultiple(number int, divisor int) bool {
    if number % divisor == 0 {
        return true
    } else {
        return false
    }
}



func main() {
    sum := 0
    for n, bound := 3, 1000; n<bound; n++ {
        if isMultiple(n, 3) || isMultiple(n,5) {
            sum += n
        }
    }
    fmt.Printf("Sum of multiples of 3 or 5 below 1000 is: %d\n", sum)
}

