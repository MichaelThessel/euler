package main

import "fmt"
import "math"

type triplet struct {
    a int
    b int
    c int
}

func main() {
    fmt.Println(pythagoreanTriplet())
}

func pythagoreanTriplet() int {
    var t triplet

    for t.c = 1000; t.c >= 0; t.c-- {
        for t.b = t.c - 1; t.b >= 0; t.b-- {
            for t.a = t.b - 1; t.a >= 0; t.a-- {
                if
                    t.a + t.b + t.c == 1000 &&
                    math.Pow(float64(t.a), 2) + math.Pow(float64(t.b), 2) == math.Pow(float64(t.c), 2) {
                        return t.a * t.b * t.c
                 }

            }
        }
    }

    return -1
}
