// Exercise  2: Variables and Data Types

package main

import "fmt"

func main() {
	// Using var
	var age int = 25
	var height float64 = 183.5
	var name string = "John"
	var isStudent bool = true

	// Using short declaration (:=)
	weight := 70.5
	country := "USA"
	isEmployed := false

	// Print values and types
	fmt.Printf("age: %d, type: %T\n", age, age)
	fmt.Printf("height: %.1f, type: %T\n", height, height)
	fmt.Printf("name: %s, type: %T\n", name, name)
	fmt.Printf("isStudent: %t, type: %T\n", isStudent, isStudent)

	fmt.Printf("weight: %.1f, type: %T\n", weight, weight)
	fmt.Printf("country: %s, type: %T\n", country, country)
	fmt.Printf("isEmployed: %t, type: %T\n", isEmployed, isEmployed)
}

// Difference between var and :=:
// var allows explicit type declaration and initialization, or declaration without initialization.
// := is used for short variable declaration with automatic type inference and requires initialization.

// Print the type of a variable: Use fmt.Printf with %T format specifier.

// Change the type of a variable: No, Go is statically typed,
// and once a variable's type is declared, it cannot be changed.
// You need to declare a new variable for a different type.
