package main

import "fmt"

func main() {
    x := 5 
    zeroByVal(x)
    fmt.Println("x is =", x)

    assignByPointer(&x, 0)
    fmt.Println("x is =", x)

    xPtr := new(int)
    assignByPointer(xPtr, 1)
    fmt.Println("xPtr is =", *xPtr)

    a := 3
    b := 5
    fmt.Println("swapping a=3 and b=5 yields :")
    swap(&a, &b)
    fmt.Println("a =", a, "and b =", b)
}

func zeroByVal(x int) {
    x = 0
}

func assignByPointer(xPtr *int, val int) {
    *xPtr = val
}

func swap(aPtr *int, bPtr *int) {
    tmp := *aPtr
    *aPtr = *bPtr
    *bPtr = tmp
} 