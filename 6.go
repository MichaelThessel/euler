package main

import "fmt"
import "math"

func main() {
    fmt.Println(sumSquareDiff());
}

func sumSquareDiff() (int) {
    sumOfSquares := 0
    sum := 0

    for i := 1; i <= 100; i++ {
        sumOfSquares += int(math.Pow(float64(i), 2))
        sum += i
    }

    return int(math.Pow(float64(sum), 2)) - sumOfSquares
}
