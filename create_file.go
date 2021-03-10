// Golang program to read and write the files
package main

import (
	"fmt"
	"log"
	"os"
)

func CreateFile() {
	//Fatalf method of
	// log prints the error message and stops
	// program execution

	file, err := os.Create("test.txt")

	if err != nil {
		log.Fatalf("Failed to create file : %s", err)
	}

	// Defer is used for purposes of cleanup like
	// closing a running file after the file has
	// been written and main //function has
	// completed execution
	defer file.Close()

	// len variable captures the length
	// of the string written to the file.

	len, err := file.WriteString("Welcome to the world of Amol Ambkar")

	if err != nil {
		log.Fatalf("Failed writting in file : %s", err)

	}

	// Name() method returns the name of the
	// file as presented to Create() method.

	fmt.Printf("\n Name of file created is %s", file.Name())
	fmt.Printf("\n Length of string in file is %d bytes", len)
}

func main() {
	CreateFile()
}
