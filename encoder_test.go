package GolangJSON

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Galih",
		LastName:  "Setiadi",
	}

	encoder.Encode(customer)
}
