package mxlib

func IntExp(base int, exponente int) int {
    value := base;
    for i:=2; i<=exponente; i++ {
        value = value * base
    }
    return value
}

func Is_prime(number int) bool {
    for n:=2; n<number; n++ {
        if number % n == 0 {
            return false
        }
    }
    return true
}

func Is_factor(number int, factor int) bool{
    return number % factor == 0
}

func Get_primes(number int) []int {
    var primes []int

    for i, bound := 2, number; i<=bound; i++ {
        if Is_prime(i) {
            if Is_factor(bound, i) {
                bound = bound/i
                primes = append(primes, i)
            }
        }
    }
    return primes
}

func Factorize(number int) map[int]int {
    factors := make(map[int]int)

    for n, bound := 2, number; n<=bound; {
        if Is_prime(n) {
            if Is_factor(bound, n) {
                bound = bound/n
                factors[n]++
                n = 2
                continue
            }
        }
        n++
    }
    return factors
}
