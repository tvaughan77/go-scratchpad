package main

import "fmt"

func main() {
    defer func() {
        str := recover()    // Recover stops the panic and returns the value passed to the call to panic
        fmt.Println(str)
    }()
    panic("PANIC")
}

// This won't work because the panic immediately exits the current frame
/*
func main() {
    panic("PANIC")
    str := recover()
    fmt.Println(str)
}
*/