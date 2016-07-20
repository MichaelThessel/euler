package main

import "fmt"

func main() {
    fmt.Println(fibsum(4000000));
}

func fibsum(limit int) (sum int) {
    sum = 2

    prev := 1
    curr := 2
    temp := 0

    for run := true; run; run = curr < limit {
        temp = prev
        prev = curr
        curr += temp

        if curr % 2 == 0 {
            sum += curr
        }
    }

    return
}
