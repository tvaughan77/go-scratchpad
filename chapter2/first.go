package main

import "fmt"
import "os"

// Run like "go run first.go"
func main() {
    fmt.Println("hello world. my name is" , "Tom")
    os.Exit(0)
    fmt.Println("probably shouldn't see this")
}