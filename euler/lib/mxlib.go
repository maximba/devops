package mxlib

func IntLog2(n uint) uint {
	// Using uint instead of uint64 is about 25% faster
	// on x86 systems with the default Go compiler.
	var r, v uint
	if n < 1<<32 {
		v = uint(n)
	} else {
		r = 32
		v = uint(n >> 32)
	}
	if v >= 1<<16 {
		r += 16
		v >>= 16
	}
	if v >= 1<<8 {
		r += 8
		v >>= 8
	}
	if v >= 1<<4 {
		r += 4
		v >>= 4
	}
	if v >= 1<<2 {
		r += 2
		v >>= 2
	}
	r += v >> 1
	return uint(r)
}

func IntSqrt(x uint) uint {
	// The most significant bit of the square root of x is equal to log2(x)/2.
	// Using uint instead of uint64 guarantees native word size, which is
	// more than 60% faster on my x86 Netbook.
	r := uint(1 << (IntLog2(x) >> 1))
	s := r >> 1
	t := r + s
	// Try to set a bit in the intermediate result, see if the square of the
	// resulting number is smaller than the input. If not, set it as the new
	// intermediate result. Shift a bit to the right, repeat.
	for s > 0 {
		if uint(t)*uint(t) <= x {
			r = t
		}
		s >>= 1
		t = r + s
	}
	return uint(r)
}

func IntExp(base int, exponente int) int {
	value := base
	for i := 2; i <= exponente; i++ {
		value = value * base
	}
	return value
}

func Is_prime(number int) bool {
	for n := 2; n < number; n++ {
		if number%n == 0 {
			return false
		}
	}
	return true
}

func Is_factor(number int, factor int) bool {
	return number%factor == 0
}

func Get_primes(number int) []int {
	var primes []int

	for i, bound := 2, number; i <= bound; i++ {
		if Is_prime(i) {
			if Is_factor(bound, i) {
				bound = bound / i
				primes = append(primes, i)
			}
		}
	}
	return primes
}

func Factorize(number int) map[int]int {
	factors := make(map[int]int)

	for n, bound := 2, number; n <= bound; {
		if Is_prime(n) {
			if Is_factor(bound, n) {
				bound = bound / n
				factors[n]++
				n = 2
				continue
			}
		}
		n++
	}
	return factors
}
