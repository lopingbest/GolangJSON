package GolangJSON

import (
	"encoding/json"
	"fmt"
	"testing"
)

//interface kosong, berarti semua data bisa masuk
func logJson(data interface{}) {
	//proses encode
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	//dikonversi menjadi string
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJson("Galih")                      //string
	logJson(1)                            //number
	logJson(true)                         //boolean
	logJson([]string{"Galih", "Setiadi"}) //array
}
