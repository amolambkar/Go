//Golang Program code that accepts user input to read and write the files.

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile(filename, text string) {

	file, err := os.Create(filename)

	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}

	defer file.Close()

	len, err := file.WriteString(text)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)

}

func ReadFile(filename string) {

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", filename)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)
}

func main() {

	fmt.Println("Enter filename: ")
	var filename string
	fmt.Scanln(&filename)

	fmt.Println("Enter text: ")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')

	CreateFile(filename, input)
	ReadFile(filename)
}
