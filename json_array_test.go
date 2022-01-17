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

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Galih",
		Addressses: []Address{
			{
				Street:     "Jalanin aja dulu",
				Country:    "Indonesia",
				Postalcode: "50215",
			},
			{
				Street:     "Sering jalan jadian kaga",
				Country:    "wakanda",
				Postalcode: "50503",
			},
		},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Galih","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addressses":[{"Street":"Jalanin aja dulu","Country":"Indonesia","Postalcode":"50215"},{"Street":"Sering jalan jadian kaga","Country":"wakanda","Postalcode":"50503"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		return
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addressses)
}

func TestONlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalanin aja dulu","Country":"Indonesia","Postalcode":"50215"},{"Street":"Sering jalan jadian kaga","Country":"wakanda","Postalcode":"50503"}]`
	jsonBytes := []byte(jsonString)

	addreses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addreses)
	if err != nil {
		return
	}
	fmt.Println(addreses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {

	addressses := []Address{
		{
			Street:     "Jalanin aja dulu",
			Country:    "Indonesia",
			Postalcode: "50215",
		},
		{
			Street:     "Sering jalan jadian kaga",
			Country:    "wakanda",
			Postalcode: "50503",
		},
	}
	bytes, _ := json.Marshal(addressses)
	fmt.Println(string(bytes))
}
