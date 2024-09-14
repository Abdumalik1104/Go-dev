package main

import "fmt"

// Function to add two integers
func add(a int, b int) int {
	return a + b
}

// Function to swap two strings
func swap(s1, s2 string) (string, string) {
	return s2, s1
}

// Function to return quotient and remainder
func divide(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	// Test the add function
	sum := add(3, 5)
	fmt.Println("Sum:", sum)

	// Test the swap function
	s1, s2 := "hello", "world"
	swapped1, swapped2 := swap(s1, s2)
	fmt.Println("Swapped:", swapped1, swapped2)

	// Test the divide function
	quotient, remainder := divide(10, 3)
	fmt.Println("Quotient:", quotient, "Remainder:", remainder)

	// Ignore certain return values
	_, remainderOnly := divide(10, 3)
	fmt.Println("Remainder only:", remainderOnly)
}

// Questions:
// How do you define a function with multiple return values in Go?
// You can define multiple return values by listing the types in parentheses, e.g., func swap(a, b string) (string, string).

// What is the significance of named return values in Go?
// Named return values act as variables, which are automatically returned at the end of the function if a return statement is called without arguments. They improve clarity and reduce errors in complex functions.

// How can you ignore certain return values if you don't need them?
// You can use the underscore (_) to ignore specific return values. For example: _, remainder := divide(10, 3).
