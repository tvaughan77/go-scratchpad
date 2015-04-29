package main

import "fmt"

func main() {
    // Numbers
    fmt.Println("1 + 1 =", 1.0 + 1.0)

    // Strings
    fmt.Println("length of 'hello world' is =", len("hello world"))
    fmt.Println("index[1] of the string is (ASCII code)=", "HelloWorld"[1])
    fmt.Println("Hello " + "World")

    // Booleans and Operators
    fmt.Println(true && true)
    fmt.Println(true && false)
    fmt.Println(true || true)
    fmt.Println(true || false)
    fmt.Println(!true)
}