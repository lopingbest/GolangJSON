package GolangJSON

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":"Galih","LastName":"Setiadi","Age":30,"Married":false}`
	jsonBytes := []byte(jsonString)

	//tanpa pointer, data akan diubah di `Unmarshal`. Customer tidak akan mendapatkan perubahan datanya. By default JSON, pass by value
	customer := &Customer{}

	//melakukan konversi dari JSON ke tipe data di Go-Lang (Decode)
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		return
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
}
