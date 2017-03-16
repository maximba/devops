// What is the 10 001st prime number?
package main

import (
    "fmt"
    "./lib"
)

func main() {
    var last_prime int 
    for n, index := 2, 1; index<10002;  n++ {
        if mxlib.Is_prime(n) {
            index++
            last_prime = n
        }
    }
    
    fmt.Printf("The 10001st prime is: %d\n", last_prime)
}

