package main

import "fmt"

func main(){
    i := 1
    for i < 10{
        if i == 5{
        break
        }
        
        fmt.Println(i)
        i++
    }
    

}
