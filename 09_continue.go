package main

import "fmt"

func main(){
    i := 1
    for i < 10{
        i++
       if i == 5{
            fmt.Println("amol")
            continue
        } 
        
        fmt.Println(i)
        
    }    
    
}
