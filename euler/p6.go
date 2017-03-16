// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

package main

import (
    "fmt"
)

func sumsq(numbers []int) int {
    sum := 0
    for _, number := range numbers {
        number *= number
        sum += number
    }
    return sum
}

func sqsum(numbers []int) int {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return sum*sum
}

func main() {
    var numbers []int
    
    for i:=1; i<101; i++ {
        numbers = append(numbers, i)
    }
    
    diff := sqsum(numbers) - sumsq(numbers) 
    fmt.Printf("square_of_sum - sum_of_square = %d\n", diff)
}
