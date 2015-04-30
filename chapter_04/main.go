package main

import "fmt"

// Can't use ":=" assignment operator here
var global string = "Some global string"

func main() {
    // verbose variable declaration (with type, non-preferred)
    var x string = "Hello World"
    fmt.Println(x)

    var y string
    y = "Whatever"
    fmt.Println(y)

    y = y + " for real"
    fmt.Println(y)

    y += " with a plus-equals operator"
    fmt.Println(y)

    // shorthand variable declaration, type inferred
    a := 1.0
    b := "a string"
    fmt.Println("a is ", a, "and b is", b)

    // Note the global outside the scope of main()
    fmt.Println("my global string is", global)

    // Constants 
    const z string = "This is a constant string"
    
    // defining multiple variables
    var (
        d = 5
        e = 10
        f = 15
    )
    fmt.Println("d is ", d, "e is ", e, "and f is", f)

}