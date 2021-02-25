package main
import "fmt"
func main() {
    
    n := 11
    if n % 2 == 0{
        fmt.Println("Number is Even")
    } else {
        fmt.Println(n)
        fmt.Println("Number is Odd")
    }
    fmt.Println("--------------------")
    
    if m := 12 ; m % 2 == 0{
        fmt.Println(m)
        fmt.Println("Number is Even")
    } else {
        fmt.Println("Number is Odd")
    }
    
}

