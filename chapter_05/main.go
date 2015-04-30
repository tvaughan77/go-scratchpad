package main

import "fmt"

func main() {
    forLoop(7)
    forLoop2(5)
    evenOddPrinter(10)

    switchStatement(4)
}

func forLoop(limit int) {
    i := 1
    for i <= limit {
        fmt.Println(i)
        i++
    }
}

func forLoop2(limit int) {
    for i := 1; i <= limit; i++ {
        fmt.Println(i)
    }
}

func evenOddPrinter(limit int) {
    for i := 1; i <= limit; i++ {
        if i % 2 == 0 {
            fmt.Println(i, "is even")
        } else {
            fmt.Println(i, "is odd")
        }
    }
}

func switchStatement(limit int) {
    for i := 1; i <= limit; i++ {
        fmt.Println(i, "in English is", intToWord(i))
    }
}

func intToWord(number int) string {
    switch number {
        case 0 : return "Zero"
        case 1 : return "One"
        case 2 : return "Two"
        case 3 : return "Three"
        default : return "Too big / I don't know"
    }
}