package main

import "fmt"
import "math"

func main() {
    c := Circle{0,0,5}
    fmt.Println("Circle of radius", c.r, "has area", c.area(), "and perimeter", c.perimeter())

    r := Rectangle{0, 0, 10, 10}
    fmt.Println("Rectangle with dimensions {0,0,10,10} has area", r.area(), "and perimeter", r.perimeter())

    fmt.Println("So the total area of the circle and the rectangle is", totalArea(&c, &r))
    fmt.Println("And the total perimeter of the circle and the rectangle is", totalPerimeter(&c, &r))
}

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}


type Circle struct {
    x, y, r float64
}

type Rectangle struct {
    x1, y1, x2, y2 float64
}

// This "(c *Circle)" is called a "receiver" and lets us call c.area() 
func (c *Circle) area() float64 {
    return math.Pi * c.r * c.r
}

func (r *Rectangle) length() float64 {
    return distance(r.x1, r.y1, r.x1, r.y2)
}

func (r *Rectangle) width() float64 {
    return distance(r.x1, r.y1, r.x2, r.y1)
}

func (r *Rectangle) area() float64 {
    return r.length() * r.width()
}

func (c *Circle) perimeter() float64 {
    return 2 * math.Pi * c.r
}

func (r *Rectangle) perimeter() float64 {
    return r.length() * 2 + r.width() * 2
}

// Notice that this interface isn't directly referenced by Circle or Rectangle; it just kind
// of works
type Shape interface {
    area() float64
    perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, shape := range shapes {
        area += shape.area()
    }
    return area
}

func totalPerimeter(shapes ...Shape) float64 {
    var perimeter float64
    for _, shape := range shapes {
        perimeter += shape.perimeter()
    }
    return perimeter
}