package main

import "fmt"

func main() {
    fmt.Println("The result of calling namedReturnFunction() is =", namedReturnFunction())

    x, y := returns2ValuesFunction()
    fmt.Println("The values coming back from returns2ValuesFunction() are", x, y)

    fmt.Println("Summing values(3, 5, 7) with my varargs function equals", varArgsFunction(3, 5, 7))

    // Note the "..." after the slice argument to the varArgsFunction. Converts the slice to varargs
    xs := []int{1,2,3}
    fmt.Println("Sum of a slice =", varArgsFunction(xs...))

    executeClosureExample()

    // Here's how to call a function that returns a function and then use it
    // I don't know why/how this works because "i" is a local scoped variable in the function that
    // returns the closure...
    // According to the "gobook.pdf", the local "i variable which - unlike normal local variables -
    // persists between calls" totally violates the principle of least surprise for me
    nextEven := makeEvenGenerator()
    fmt.Println("next even =", nextEven())
    fmt.Println("next even =", nextEven())
    fmt.Println("next even =", nextEven())

    // Recursive factorial
    fmt.Println("Using recursion, the factorial of 7 is =", factorial(7))

    fmt.Println("not using defer, calling second(), then first()")
    second()
    first()

    fmt.Println("now using defer, calling the same order")
    defer second()  // useful for closing resources where you open them. works like a java finally
    first()
}

// Notice this syntax where you can name a returned value
func namedReturnFunction() (r int) {
    r = 1
    return
}

func returns2ValuesFunction() (int, int) {
    return 5, 6
}

func varArgsFunction(args ...int) int {
    total := 0
    for _, val := range args {
        total += val
    }
    return total
}

func executeClosureExample() {
    // Closure!
    add := func(x, y int) int {
        return x + y
    }
    fmt.Println("Adding 1 and 3 =", add(1, 3))
}

// Here's how to return a function from a function
func makeEvenGenerator() func() uint {
    i := uint(0)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}

// Recursive factorial example
func factorial(x int) int {
    if x == 0 {
        return 1
    }
    return x * factorial(x - 1)
}

func first() {
    fmt.Println("first")
}

func second() {
    fmt.Println("second")
}