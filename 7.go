package main

import "fmt"

func main() {
    fmt.Println(findPrime(10001));
}

func findPrime(count int) int {
    number := 9
    primes := []int{2, 3, 5, 7}

    var valid bool
    for run := true; run; run = len(primes) < count {
        number += 2

        if number % 2 == 0 { continue }
        if number % 3 == 0 { continue }
        if number % 5 == 0 { continue }
        if number % 7 == 0 { continue }

        valid = true;
        for _, prime := range primes {
            if number % prime == 0 {
                valid = false
                break
            }

        }

        if valid {
            primes = append(primes, number)
        }
    }

    return primes[len(primes) - 1]
}
