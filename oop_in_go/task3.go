package main

import (
	"fmt"
	"math"
)

// Define the Shape interface
type Shape interface {
	Area() float64
}

// Implement Shape interface for Circle
type Circle struct {
	Radius float64
}

// Method Area for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Implement Shape interface for Rectangle
type Rectangle struct {
	Width, Height float64
}

// Method Area for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Function to print the area of any Shape
func PrintArea(s Shape) {
	fmt.Printf("The area is: %.2f\n", s.Area())
}

func main() {
	// Create instances of Circle and Rectangle
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 3}

	// Call PrintArea with different shapes
	PrintArea(circle)    // Polymorphism: Circle as Shape
	PrintArea(rectangle) // Polymorphism: Rectangle as Shape
}

/*
Questions:

How do you define and implement an interface in Go?
An interface in Go is defined using the `type` keyword, followed by the interface name and a list of method signatures. For example:
     ```go
     type Shape interface {
         Area() float64
     }
     ```
To implement an interface, a type (struct) must define all the methods declared in the interface. Go automatically knows if a type implements an interface based on the method signatures, without explicit declarations.

What is the role of interfaces in achieving polymorphism in Go?
Interfaces enable polymorphism by allowing different types to be treated uniformly if they implement the same interface. For instance, both `Circle` and `Rectangle` implement the `Shape` interface, so we can pass either type to the `PrintArea` function, which expects a `Shape`.

How can you check if a type implements a certain interface?**
You can check if a type implements an interface by using a type assertion or the `ok` idiom:
     ```go
     if shape, ok := someVariable.(Shape); ok {
         // someVariable implements Shape
     }
     ```
Alternatively, you can simply assign a value to an interface variable. If the type doesn't implement the interface, it will cause a compile-time error.
*/
