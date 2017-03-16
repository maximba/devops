package main

//Find the largest palindrome made from the product of two 3-digit numbers.

import (
    "fmt"
    "strconv"
)

func flipstring(str string) string {
    fs := []byte(str)
    length := len(str)
    /*
    for i, value := range str {
        fs[length-1-i] = byte(value)
    }
    */
    for i:=0; i< length; i++ {
        fs[length-1-i] = str[i]
    }
    
    return string(fs)
}

func isPalyndrome(str string) bool {
    fstr := flipstring(str)
    if str == fstr {
        return true
    }
    return false
}

func main() {
    max3d := 999
    var p, f1, f2 int
    max := 0
    
    for i := max3d; i>0; i-- {
        for j := max3d; j>0; j-- {
            p = i*j
            if p>max {
                if isPalyndrome(strconv.Itoa(p)) {
                    if max < p {
                        max, f1, f2 = p, i, j
                    }
                }        
            }
        }
    }
    fmt.Printf("Max Palyndrome is: %d = %d * %d\n", max, f1, f2)
}
