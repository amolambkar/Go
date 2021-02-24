package main

import "fmt"

func main() {
	var x float32
	for i:=0 ; i < 2 ; i++ {
	fmt.Println("Enter a float number:")
	fmt.Scan(&x)
	fmt.Println("The truncated number you just entered is ", int32(x))
	fmt.Println("-------------------------------------------------------- ")
	}
}
