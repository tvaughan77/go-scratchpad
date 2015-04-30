package main

import "fmt"

func main() {
    // Problem 2
    half1, isEven1 := halveEvenOdd(1)
    fmt.Println("halveEvenOdd(1) is =", half1, isEven1)
    half2, isEven2 := halveEvenOdd(2)
    fmt.Println("halveEvenOdd(2) is =", half2, isEven2)

    // Problem 3
    xs := []int{5, 2, 88, 6, 1, 9, 44, 101, 4, 6, 8}
    fmt.Println("The greatest number in (", xs, ") is =", greatest(xs...))

    // Problem 4
    oddGenFunc := makeOddGenerator()
    fmt.Println("next odd =", oddGenFunc())
    fmt.Println("next odd =", oddGenFunc())
    fmt.Println("next odd =", oddGenFunc())

    // Problem 5
    for i := 0; i < 10; i++ {
        fmt.Println("fib(", i, ") =", fib(i))
    }
}

func halveEvenOdd(x int) (int, bool) {
    half := x / 2
    if x % 2 == 0 {
        return half, true
    } else {
        return half, false
    }
}

func greatest(xs ...int) int {
    max := -999999
    for _, v := range xs {
        if v > max {
            max = v
        }
    }
    return max
}

func makeOddGenerator() func() uint {
    i := uint(1)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}

func fib(n int) int {
    if n == 0 {
        return 1
    } else if n == 1 {
        return 1
    } else {
        return fib(n - 2) + fib (n -1)
    }
}