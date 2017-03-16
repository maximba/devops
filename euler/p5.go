// What is the smallest positive number that is evenly divisible 
// by all of the numbers from 1 to 20?


package main

import (
    "fmt"
    "./lib"
)

type Factors map[int]int
type MapFactors map[int]Factors

func intExp(base int, exponente int) int {
    value := base;
    for i:=2; i<=exponente; i++ {
        value = value * base
    }
    return value
}

func mapFactors(numbers []int) MapFactors {
    factors := make(MapFactors)
    for _, number := range numbers {
        factors[number] = mxlib.Factorize(number)
    }
    return factors
}

func reduceFactors(factors Factors) int {
    value := 1
    for base, exp := range factors {
        value = value * intExp(base, exp)
    }
    return value
}

func get_lcm(numbers []int) int {
    mapfactors := mapFactors(numbers)
    lcm := make(Factors)
    for _, factors := range mapfactors {
        for factor, cur_exp := range factors {
            if prev_exp, ok := lcm[factor]; ok {
                if prev_exp > cur_exp {
                    continue
                }
            } 
            lcm[factor] = cur_exp
        }
    }
    return reduceFactors(lcm)
}

func main() {
    const maxdivisor = 20    
    var divisors []int
    
    for i:=2; i<=maxdivisor; i++ {
        divisors = append(divisors, i)
    }
    
    fmt.Printf("The minimum divisible number is: %d\n", get_lcm(divisors)) 
}
