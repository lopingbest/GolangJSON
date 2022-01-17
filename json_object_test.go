package GolangJSON

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	Postalcode string
}

type Customer struct {
	FirstName  string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addressses []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName: "Galih",
		LastName:  "Setiadi",
		Age:       30,
		Married:   false,
	}
	//proses encode
	bytes, _ := json.Marshal(customer)
	//bytes dikonversi menjadi sring
	fmt.Println(string(bytes))
}
