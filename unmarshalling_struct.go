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
	var human1 Human

	// data in JSON format which
	// is to be decoded

	Data := []byte(`{
		"Name":"Amol",
		"Age":20,
		"Address":"Panvel"
		}`)

	// decoding human1 struct
	// from json format
	err := json.Unmarshal(Data, &human1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Struct is:", human1)
	fmt.Printf("%s's Age is %d.\n", human1.Name, human1.Age)

	// unmarshaling a JSON array
	// to array type in Golang

	// defining an array instance
	// of struct type
	var human2 []Human

	// JSON array to be decoded
	// to an array

	Data2 := []byte(`[
	{"Name":"Amol","Age":20,"Address":"Panvel"},
	{"Name":"Giriraj","Age":20,"Address":"Usmanabad"},
	{"Name":"Sushant","Age":20,"Address":"Beed"}
	]`)

	// decoding JSON array to
	// human2 array
	err = json.Unmarshal(Data2, &human2)

	if err != nil {
		fmt.Println(err)

	}

	for i := range human2 {
		fmt.Println(human2[i])
	}
}
