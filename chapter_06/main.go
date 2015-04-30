package main

import "fmt"

func main() {
    var x [5]int
    x[4] = 100
    fmt.Println(x)

    y := []float64{ 12, 13, 14, 15, 16 }
    fmt.Println("The average of ", y, "is", average(y))
    fmt.Println("Using the range operator, the average was computed as", averageUsingRange(y))

    var slice = makeSlice(5, 10)
    fmt.Println("my slice =", slice)
    fmt.Println("my slice has a length of =", len(slice))
    var slice2 = append(slice, 8, 9)
    fmt.Println("my slice2 =", slice2)

    // Note that when you copy a larger array (slice2) into a smaller array (slice3)
    // it only copies up to the size of the smaller array
    var slice3 = makeSlice(2, 2)
    copy(slice3, slice2)   
    fmt.Println("my slice3 =", slice3)
}

func average(array []float64) float64 {
    var sum float64
    var arrayLen = len(array)
    for i := 0; i < arrayLen; i++ {
        sum += array[i]
    }
    return sum / float64(arrayLen)
}

// Just like the "average" method, above but with Go's "range" syntax
func averageUsingRange(array []float64) float64 {
    var sum float64
    for _, value := range array {   // don't use the var "i", so sub "_"
        sum += value
    }
    return sum / float64(len(array))
}

func makeSlice(low int, high int) []int {
    x := make([]int, low, high)
    return x
}