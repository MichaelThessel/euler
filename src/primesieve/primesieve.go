package primesieve

func Primes(count int) []int {
    return cull(filter(setup(count), 0))
}

func setup(count int) []int {
    primes := make([]int, count - 1, count - 1)

    for i := 2; i <= count; i++ {
        primes[i - 2] = i
    }

    return primes
}

func filter(primes []int, index int) []int {
    for i := index + 1; i < len(primes); i++ {
        if primes[i] != 0 && primes[i] % primes[index] == 0 {
            primes[i] = 0
        }
    }

    for i := index + 1; i < len(primes); i++ {
        if primes[i] != 0 {
            filter(primes, i)
            break
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
