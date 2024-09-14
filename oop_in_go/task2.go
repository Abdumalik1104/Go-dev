package main

import "fmt"

// Define the Employee struct
type Employee struct {
	Name string
	ID   int
}

// Method for Employee called Work
func (e Employee) Work() {
	fmt.Printf("Employee %s with ID %d is working.\n", e.Name, e.ID)
}

// Define the Manager struct that embeds Employee and adds Department
type Manager struct {
	Employee   // Embedding Employee struct
	Department string
}

func main() {
	// Create an instance of Manager, set fields, and call the Work method
	manager := Manager{
		Employee:   Employee{Name: "Alice", ID: 101},
		Department: "HR",
	}

	// Call the Work method
	manager.Work() // This calls the Work method from the embedded Employee struct
}

/*
Questions:

What is embedding in Go, and how does it relate to composition?
Embedding in Go is a way to include one struct inside another,
allowing the outer struct to reuse the fields and methods of
the embedded struct. It is Go's approach to composition,
where a struct can "inherit" fields and behavior from another struct.

How does Go handle method calls on embedded types?
Go automatically promotes the methods of an embedded type,
so they can be called directly on the outer struct.
For example, calling `manager.Work()` is the same as
calling `manager.Employee.Work()`.

Can an embedded type override a method from the outer struct?
An embedded type cannot override methods, but the outer struct
can define a method with the same name, which will take precedence.
If the outer struct defines a method with the same name,
that method will be called instead of the embedded type's method.
*/
