package main

//What is the largest prime factor of the number 600851475143 

import (
    "fmt"
)

func is_prime(number int) bool {
    for n:=2; n<number; n++ {
        if number % n == 0 {
            return false
        }
    }    
    return true
}

func is_factor(number int, factor int) bool{
    return number % factor == 0
}

func main() {
    number := 600851475143
    factors := make(map[int]int)
    lowprime := 2
    maxprime := 2
        
    for n, bound:=lowprime, number; n<=bound; {
        if is_prime(n) {
            if is_factor(bound, n) {
                bound = bound/n
                factors[n]++
                n = 2
                continue
            }
        } 
        n++
    }

    for k, _ := range factors {
        if k > maxprime {
            maxprime = k
        }
    }

    fmt.Printf("Los factores primos son: %v\n", factors)
    fmt.Printf("El meyor factor primo es: %d\n", maxprime)
}
