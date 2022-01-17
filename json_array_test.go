package GolangJSON

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName: "Galih",
		LastName:  "Setiadi",
		Age:       22,
		Married:   false,
		Hobbies:   []string{"Gaming", "Reading", "Mencintaimu"},
	}
	//proses encode
	bytes, _ := json.Marshal(customer)
	//bytes dikonversi menjadi sring
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Galih","LastName":"Setiadi","Age":22,"Married":false,"Hobbies":["Gaming","Reading","Mencintaimu"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		return
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}
