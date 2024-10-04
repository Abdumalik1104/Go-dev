// Exercise 3: Control Structures

// 1. Using if Statement: Write a program that takes an integer input and prints whether the number is positive, negative, or zero.
// package main

// import "fmt"

// func main() {
// 	var num int
// 	fmt.Print("Enter an integer: ")
// 	fmt.Scan(&num)

// 	if num > 0 {
// 		fmt.Println("The number is positive.")
// 	} else if num < 0 {
// 		fmt.Println("The number is negative.")
// 	} else {
// 		fmt.Println("The number is zero.")
// 	}
// }

// 2. Using for Loop: Implement a for loop that calculates the sum of the first 10 natural numbers.
// package main

// import "fmt"

// func main() {
// 	sum := 0
// 	for i := 1; i <= 10; i++ {
// 		sum += i
// 	}
// 	fmt.Println("The sum of the first 10 natural numbers is:", sum)
// }

// 3. Using switch Statement: Write a switch statement that prints the day of the week based on an integer input (1 for Monday, 2 for Tuesday, etc.).
package main

import "fmt"

func main() {
	var day int
	fmt.Print("Enter a number (1-7) for the day of the week: ")
	fmt.Scan(&day)

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid input. Please enter a number between 1 and 7.")
	}
}

// Questions:
// if in Go vs. Python/Java: Go's if allows an optional short statement (e.g., variable declaration) before the condition. Go does not require parentheses around conditions, unlike Java, but like Python, it uses braces {} for the block.

// for loops in Go:
// Basic form: for i := 0; i < 10; i++ { }.
// While-like loop: for i < 10 { }.
// Infinite loop: for { }.

// switch in Go vs. C/Java:
// No break needed in Go (implicit break after each case).
// Go allows multiple expressions in a case, supports switch on types (type switch), and allows cases to evaluate expressions (e.g., ranges).
