package main

import (
	"encoding/json"
	"fmt"
)

// Define the Product struct
type Product struct {
	Name     string  `json:"name"`     // Struct tag to map JSON key "name"
	Price    float64 `json:"price"`    // Struct tag to map JSON key "price"
	Quantity int     `json:"quantity"` // Struct tag to map JSON key "quantity"`
}

// Function to convert a Product to JSON
func toJSON(p Product) (string, error) {
	jsonData, err := json.Marshal(p)
	if err != nil {
		return "", err // Return error if encoding fails
	}
	return string(jsonData), nil
}

// Function to decode JSON into a Product
func fromJSON(jsonString string) (Product, error) {
	var p Product
	err := json.Unmarshal([]byte(jsonString), &p)
	if err != nil {
		return Product{}, err // Return error if decoding fails
	}
	return p, nil
}

func main() {
	// Create an instance of Product
	product := Product{Name: "Laptop", Price: 1200.50, Quantity: 5}

	// Convert Product to JSON
	jsonData, err := toJSON(product)
	if err != nil {
		fmt.Println("Error encoding to JSON:", err)
	} else {
		fmt.Println("JSON:", jsonData)
	}

	// Decode JSON back into Product
	jsonString := `{"name":"Phone","price":799.99,"quantity":10}`
	decodedProduct, err := fromJSON(jsonString)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	} else {
		fmt.Printf("Decoded Product: %+v\n", decodedProduct)
	}
}

/*
Questions:

How do you work with JSON in Go?
In Go, the `encoding/json` package provides functions
to encode (`Marshal`) and decode (`Unmarshal`) JSON data.
To encode a struct into JSON, use `json.Marshal`.
To decode JSON into a struct, use `json.Unmarshal`.

What role do struct tags play in JSON encoding/decoding?
Struct tags specify how struct fields map to JSON keys.
The tag is placed after the field definition, like this:
`Name string \`json:"name"\``. Tags allow customization of
JSON key names, case, or other options.

How do you handle errors that may occur during JSON encoding/decoding?
Both `json.Marshal` and `json.Unmarshal` return an error as
their second return value. You should always check for errors and
handle them appropriately, typically by logging the error or returning
it to the caller.
*/
