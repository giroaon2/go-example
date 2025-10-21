package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"product_name"`
	Price float64 `json:"price"`
}

func main() {
	prod := Product{
		ID:    "A123",
		Name:  "Laptop",
		Price: 1200.50,
	}

	jsonOutput, err := json.Marshal(prod)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}
	fmt.Println("Encoded JSON:", string(jsonOutput))
}
