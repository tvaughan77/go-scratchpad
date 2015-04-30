package main

import "fmt"

func main() {
    myArray := [6]int{1, 2, 3, 4, 5, 6}
    fmt.Println("The fourth element of myArray is =", myArray[3])

    myArray2 := make([]int, 3, 9)
    fmt.Println("The length of an array created with 'make([]int, 3, 9)' is =", len(myArray2))

    x := [6]string{"a", "b", "c", "d", "e", "f"}
    fmt.Println("raw array x is =", x)
    fmt.Println("slice x[2:5] is = ", x[2:5])

    z := []int{
        48,96,86,68,
        57,82,63,70,
        37,34,83,27,
        19,97, 9,17,
    }
    var smallest = findSmallest(z)
    fmt.Println("The smallest element of array z (", z, ") is", smallest)
}

func findSmallest(array []int) int {
    var smallest int = 9999999
/*  more verbose for-loop
    for i := 0; i < len(array); i++ {
        if array[i] < smallest {
            smallest = array[i]
        }
    }
*/
    // for-loop using a range is kind of nice
    for _, value := range array {
        if (value < smallest) {
            smallest = value
        }
    }
    return smallest
}