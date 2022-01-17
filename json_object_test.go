package GolangJSON

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName string
	LastName  string
	Age       int
	Married   bool
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "Galih",
		LastName:  "Setiadi",
		Age:       30,
		Married:   false,
	}
	bytes, _ := json.Marshal(customer)
	//bytes dikonversi menjadi sring
	fmt.Println(string(bytes))
}
