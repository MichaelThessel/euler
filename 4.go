package main

import "fmt"
import "math"

type palindrome struct {
    product int
    a int
    b int
}

func main() {
    pd := findPalindrome()
    fmt.Println(pd.product, pd.a, pd.b)
}

func getNumberMagnitude(number int) int {
    if number < 0 { number *= - 1 }
    mag := 0
    for mag = 0; int(math.Pow(10, float64(mag))) < number; mag++ {}

    return mag - 1
}

func getNumberDigits(number int) []int {
    mag := getNumberMagnitude(number)
    digits := make([]int, mag + 1, mag + 1)

    for i := mag; i >= 0; i-- {
        div := math.Pow(10, float64(i))
        digits[i] = int(int(float64(number) / div * 100) / 100)
        number -= int(digits[i] * int(div))
    }

    return digits
}

func isPalindrome(number int) (bool) {
    digits := getNumberDigits(number)
    for i:= 0; i <= len(digits) / 2 - 1; i++ {
        if digits[i] != digits[len(digits) - i - 1] {
            return false
        }
    }

    return true
}

func findPalindrome() palindrome {
    pd := palindrome {
        product: 0,
        a: 0,
        b: 0,
    }

    prod := 0
    for i := 999; i >= 100; i-- {
        for j := 999; j >= 100; j-- {
            prod = i * j
            if prod > pd.product && isPalindrome(prod) {
                pd.product = prod
                pd.a = i
                pd.b = j
            }
        }
    }

    return pd
}
