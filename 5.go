package main

import "fmt"

func main() {
    fmt.Println(findSmallestMultiple());
}

func findSmallestMultiple() (int) {
    number := 20
    divisors := [10]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

    var valid bool
    for run := true; run; run = !valid {
        valid = true

        for _, divisor := range divisors {
            if number % divisor != 0 {
                valid = false
                break
            }
        }

        if valid { return number }

        number += 2
    }

    return number
}
