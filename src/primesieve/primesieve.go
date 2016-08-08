package primesieve

import "math"

func Primes(count int) []int {
    return cull(filter(setup(count)))
}

func setup(count int) []int {
    primes := make([]int, count + 1, count + 1)

    for i := 2; i <= count; i++ {
        primes[i] = i
    }

    return primes
}

func filter(primes []int) []int {
    primeCount := len(primes) - 1
    for i := 2; i < int(math.Sqrt(float64(primeCount))); i++ {
        if primes[i] != 0 {
            it := 0
            i2 := int(math.Pow(float64(i), 2))
            for j:= i2; j <= primeCount; j = i2 + it * i  {
                primes[j] = 0;
                it++
            }
        }
    }

    return primes
}

func cull(primes []int) []int {
    j := 0;
    for i := 0; i < len(primes); i++ {
        if primes[i] != 0 {
            primes[j] = primes[i]
            j++
        }
    }

    return primes[0:j]
}
