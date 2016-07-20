package main

import "fmt"

func main() {
    fmt.Println(largestPrimeFactor(600851475143));
}

func largestPrimeFactor(number int) (factor int) {
    i := 2;
    factor = number

    for run := true; run; run = i < factor {

        if factor % i == 0 {
            factor /= i
        }
        i++
    }

    return
}
