package main

import "fmt"


func main() {
    // Initialize a map of strings to ints
    x := make(map[string]int)
    x["key"] = 10
    fmt.Println(x)
    fmt.Println("the value of 'key' in the map is", x["key"])

    var elements = makeElementsMapShorthand()
    fmt.Println("the element whose symbol is 'Li' is =", elements["Li"])

    // Here's some go-syntax to check if a key exists and do something with it if it does
    if symbol, ok := elements["Un"]; ok {
        fmt.Println("the element whose symbol is 'Un' is =", symbol)
    } else {
        fmt.Println("I don't know the element whose symbol is 'Un'")
    }
    
}

func makeElementsMap() map[string]string {
    elements := make(map[string]string)
    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["Li"] = "Lithium"
    elements["Be"] = "Beryllium"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["N"] = "Nitrogen"
    elements["O"] = "Oxygen"
    elements["F"] = "Fluorine"
    elements["Ne"] = "Neon"
    return elements
}

func makeElementsMapShorthand() map[string]string {
    elements := map[string]string {
        "H": "Hydrogen",
        "He": "Helium",
        "Li": "Lithium",
        "Be": "Beryllium",
        "B": "Boron",
        "C": "Carbon",
        "N": "Nitrogen",
        "O": "Oxygen",
        "F": "Fluorine",
        "Ne": "Neon",
    }
    return elements
}