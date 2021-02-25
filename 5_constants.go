package main

import(
 "fmt"
 "reflect"
 "math"
)

func main(){
    const s string = "string constant"
    fmt.Println(s)
    fmt.Println(reflect.TypeOf(s))
    fmt.Println("----------------------------")
    
    const a = 50000000
    
    const b = 3e20/a
    
    fmt.Println(b)
    fmt.Println(reflect.TypeOf(b))
    fmt.Println("----------------------------")

    fmt.Println(int64(b))
        
    
    fmt.Println("----------------------------")
    
    fmt.Println(math.Sin(a))
    fmt.Println("----------------------------")
    
    
}
