// Golang program to illustrate the
// concept of encoding using JSON

package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name    string
	Age     int
	Address string
}

func main() {
	// defining a struct instance
	human1 := Human{"Amol", 20, "Panvel"}

	// encoding human1 struct
	// into json format
	human_enc, err := json.Marshal(human1)

	if err != nil {
		fmt.Println(err)
	}

	// as human_enc is in a byte array
	// format, it needs to be
	// converted into a string
	fmt.Println(string(human_enc))

	// converting slices from
	// golang to JSON fomat

	// defining an array
	// of struct instance
	human2 := []Human{
		{Name: "Amol", Age: 20, Address: "Panvel"},
		{Name: "Giriraj", Age: 20, Address: "Usmanabad"},
		{Name: "Sushant", Age: 20, Address: "Beed"},
	}

	human2_enc, err := json.Marshal(human2)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	fmt.Println(string(human2_enc))

	fmt.Println(human_enc)
}
