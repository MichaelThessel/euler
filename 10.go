package main

import (
    "fmt"
    "primesieve"
)

func main() {
    fmt.Println(primeSum(2000000));
}

func primeSum(count int) int {
    primes := primesieve.Primes(count)

    sum := 0
    for i := 0; i < len(primes); i++ {
        sum += primes[i]
    }

    return sum
}
