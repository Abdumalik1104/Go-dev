// Define a struct Person with fields Name and Age:
package main

import "fmt"

// Define the Person struct
type Person struct {
	Name string
	Age  int
}

// Method Greet for the Person struct
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	// Create an instance of Person
	person := Person{Name: "Alice", Age: 30}

	// Call the Greet method
	person.Greet()
}

/*
Questions:

How do you define a struct in Go?
Use the type keyword followed by the name of the struct and its fields. For example:
type Person struct {
    Name string
    Age  int
}

How do methods differ from regular functions in Go?
Methods are functions with a receiver argument. They are associated with a specific type (struct), while regular functions are not associated with any type.

Can a method in Go be associated with types other than structs?
No, methods can only be associated with named types like structs or pointers to structs. Go does not support methods for primitive types or interfaces.
*/
