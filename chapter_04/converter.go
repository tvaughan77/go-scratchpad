package main

import "fmt"
import "os"

func main() {
    fmt.Println("Do you want to convert farenheit (enter 'f') or meters (enter 'm')")
    var answer string
    fmt.Scanf("%s", &answer)

    if (answer == "f") {
        // There's a bug somewhere here.  No matter the input, the value is 0...
        var value = promptForInput("Enter a number of Farenheit degrees to convert to Celcius")
        fmt.Println(value, "Farenheit is", convertFtoC(value), "Celcius")
    } else if (answer == "m") {
        var value = promptForInput("Enter a number of feet to convert to meters")
        fmt.Println(value, "meters is", convertMtoF(value), "feet")
    } else {
        fmt.Println("Unknown selection", answer)
        fmt.Println("Exiting")
        os.Exit(1)
    }
}

func promptForInput(prompt string) float64 {
    fmt.Println(prompt)
    var input float64
    fmt.Scanf("%f", &input)
    return input
}

func convertFtoC(farenheit float64) float64 {
    return (farenheit - 32.0) * (5/9)
}

func convertMtoF(meters float64) float64 {
    return meters / 0.3048
}