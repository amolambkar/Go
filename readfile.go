package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadFile() {

	FileName := "test.txt"

	data, err := ioutil.ReadFile(FileName)

	if err != nil {
		log.Panicf("Failed reading file : %s", err)

	}

	fmt.Printf("\n File Name : %s ", FileName)
	fmt.Printf("\n Size : %d bytes", len(data))
	fmt.Printf("\n Data : %s", data)
}

func main() {
	ReadFile()
}
