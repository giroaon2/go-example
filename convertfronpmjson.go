package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city,omitempty"` // omitempty excludes if zero value
}

func main() {
	jsonData := []byte(`{"name":"Jane Doe","age":28,"city":"New York"}`)
	var p Person
	err := json.Unmarshal(jsonData, &p)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}
	fmt.Printf("Decoded Person: %+v\n", p)
}
